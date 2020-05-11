package cos

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/Xuanwo/storage/pkg/headers"
	"github.com/Xuanwo/storage/pkg/iowrap"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/info"
)

// Storage is the cos object storage service.
type Storage struct {
	bucket *cos.BucketService
	object *cos.ObjectService

	name     string
	location string
	workDir  string
}

// String implements Storager.String
func (s *Storage) String() string {
	return fmt.Sprintf(
		"Storager cos {Name: %s, WorkDir: %s}",
		s.name, s.workDir,
	)
}

// Metadata implements Storager.Metadata
func (s *Storage) Metadata(pairs ...*types.Pair) (m info.StorageMeta, err error) {
	m = info.NewStorageMeta()
	m.Name = s.name
	m.WorkDir = s.workDir
	return m, nil
}

// ListDir implements Storager.ListDir
func (s *Storage) ListDir(path string, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpListDir, err, path)
	}()

	opt, err := s.parsePairListDir(pairs...)
	if err != nil {
		return err
	}

	marker := ""
	delimiter := "/"
	limit := 200

	rp := s.getAbsPath(path)

	for {
		req := &cos.BucketGetOptions{
			Prefix:    rp,
			MaxKeys:   limit,
			Marker:    marker,
			Delimiter: delimiter,
		}

		resp, _, err := s.bucket.Get(opt.Context, req)
		if err != nil {
			return err
		}

		if opt.HasDirFunc {
			for _, v := range resp.CommonPrefixes {
				o := &types.Object{
					ID:         v,
					Name:       s.getRelPath(v),
					Type:       types.ObjectTypeDir,
					ObjectMeta: info.NewObjectMeta(),
				}

				opt.DirFunc(o)
			}
		}

		if opt.HasFileFunc {
			for _, v := range resp.Contents {
				o, err := s.formatFileObject(v)
				if err != nil {
					return err
				}

				opt.FileFunc(o)
			}
		}

		marker = resp.NextMarker
		if !resp.IsTruncated {
			break
		}
	}

	return
}

// ListPrefix implements Storager.ListPrefix
func (s *Storage) ListPrefix(prefix string, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpListPrefix, err, prefix)
	}()

	opt, err := s.parsePairListPrefix(pairs...)
	if err != nil {
		return err
	}

	marker := ""
	limit := 200

	rp := s.getAbsPath(prefix)

	for {
		req := &cos.BucketGetOptions{
			Prefix:  rp,
			MaxKeys: limit,
			Marker:  marker,
		}

		resp, _, err := s.bucket.Get(opt.Context, req)
		if err != nil {
			return err
		}

		for _, v := range resp.Contents {
			o, err := s.formatFileObject(v)
			if err != nil {
				return err
			}

			opt.ObjectFunc(o)
		}

		marker = resp.NextMarker
		if !resp.IsTruncated {
			break
		}
	}

	return
}

// Read implements Storager.Read
func (s *Storage) Read(path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	defer func() {
		err = s.formatError(services.OpRead, err, path)
	}()

	opt, err := s.parsePairRead(pairs...)
	if err != nil {
		return nil, err
	}

	rp := s.getAbsPath(path)

	resp, err := s.object.Get(opt.Context, rp, nil)
	if err != nil {
		return nil, err
	}

	r = resp.Body

	if opt.HasReadCallbackFunc {
		r = iowrap.CallbackReadCloser(r, opt.ReadCallbackFunc)
	}
	return
}

// Write implements Storager.Write
func (s *Storage) Write(path string, r io.Reader, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpWrite, err, path)
	}()

	opt, err := s.parsePairWrite(pairs...)
	if err != nil {
		return err
	}

	rp := s.getAbsPath(path)

	putOptions := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentLength: int(opt.Size),
		},
	}
	if opt.HasChecksum {
		putOptions.ContentMD5 = opt.Checksum
	}
	if opt.HasStorageClass {
		putOptions.XCosStorageClass = opt.StorageClass
	}
	if opt.HasReadCallbackFunc {
		r = iowrap.CallbackReader(r, opt.ReadCallbackFunc)
	}

	_, err = s.object.Put(opt.Context, rp, r, putOptions)
	if err != nil {
		return err
	}
	return
}

// Stat implements Storager.Stat
func (s *Storage) Stat(path string, pairs ...*types.Pair) (o *types.Object, err error) {
	defer func() {
		err = s.formatError(services.OpStat, err, path)
	}()

	opt, err := s.parsePairStat(pairs...)
	if err != nil {
		return nil, err
	}

	rp := s.getAbsPath(path)

	output, err := s.object.Head(opt.Context, rp, nil)
	if err != nil {
		return nil, err
	}

	o = &types.Object{
		ID:         rp,
		Name:       path,
		Type:       types.ObjectTypeFile,
		Size:       output.ContentLength,
		ObjectMeta: info.NewObjectMeta(),
	}

	// COS uses RFC1123 format in HEAD
	//
	// > Last-Modified: Fri, 09 Aug 2019 10:20:56 GMT
	//
	// ref: https://cloud.tencent.com/document/product/436/7745
	if v := output.Header.Get(headers.LastModified); v != "" {
		lastModified, err := time.Parse(time.RFC1123, v)
		if err != nil {
			return nil, err
		}
		o.UpdatedAt = lastModified
	}

	if v := output.Header.Get(headers.ContentType); v != "" {
		o.SetContentType(v)
	}

	if v := output.Header.Get(headers.ETag); v != "" {
		o.SetETag(output.Header.Get(v))
	}

	if v := output.Header.Get(storageClassHeader); v != "" {
		setStorageClass(o.ObjectMeta, v)
	}

	return o, nil
}

// Delete implements Storager.Delete
func (s *Storage) Delete(path string, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpDelete, err, path)
	}()

	opt, err := s.parsePairDelete(pairs...)
	if err != nil {
		return err
	}

	rp := s.getAbsPath(path)

	_, err = s.object.Delete(opt.Context, rp)
	if err != nil {
		return err
	}
	return nil
}

// getAbsPath will calculate object storage's abs path
func (s *Storage) getAbsPath(path string) string {
	prefix := strings.TrimPrefix(s.workDir, "/")
	return prefix + path
}

// getRelPath will get object storage's rel path.
func (s *Storage) getRelPath(path string) string {
	prefix := strings.TrimPrefix(s.workDir, "/")
	return strings.TrimPrefix(path, prefix)
}

func (s *Storage) formatError(op string, err error, path ...string) error {
	if err == nil {
		return nil
	}

	return &services.StorageError{
		Op:       op,
		Err:      formatError(err),
		Storager: s,
		Path:     path,
	}
}

func (s *Storage) formatFileObject(v cos.Object) (o *types.Object, err error) {
	o = &types.Object{
		ID:         v.Key,
		Name:       s.getRelPath(v.Key),
		Type:       types.ObjectTypeFile,
		Size:       int64(v.Size),
		ObjectMeta: info.NewObjectMeta(),
	}

	// COS returns different value depends on object upload method or
	// encryption method, so we can't treat this value as content-md5
	//
	// ref: https://cloud.tencent.com/document/product/436/7729
	if v.ETag != "" {
		o.SetETag(v.ETag)
	}

	// COS uses ISO8601 format: "2019-05-27T11:26:14.000Z" in List
	//
	// ref: https://cloud.tencent.com/document/product/436/7729
	if v.LastModified != "" {
		t, err := time.Parse("2006-01-02T15:04:05.999Z", v.LastModified)
		if err != nil {
			return nil, err
		}
		o.UpdatedAt = t
	}

	if value := v.StorageClass; value != "" {
		setStorageClass(o.ObjectMeta, value)
	}

	return o, nil
}
