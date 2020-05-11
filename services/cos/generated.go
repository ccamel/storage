// Code generated by go generate via internal/cmd/service; DO NOT EDIT.
package cos

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/endpoint"
	"github.com/Xuanwo/storage/pkg/httpclient"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/info"
	ps "github.com/Xuanwo/storage/types/pairs"
)

var _ credential.Provider
var _ endpoint.Provider
var _ segment.Segment
var _ storage.Storager
var _ services.ServiceError
var _ httpclient.Options

// Type is the type for cos
const Type = "cos"

// Service available pairs.
const (
	// StorageClass will
	PairStorageClass = "cos_storage_class"
)

// Service available infos.
const (
	InfoObjectMetaStorageClass = "cos-storage-class"
)

// WithStorageClass will apply storage_class value to Options
// This pair is used to
func WithStorageClass(v string) *types.Pair {
	return &types.Pair{
		Key:   PairStorageClass,
		Value: v,
	}
}

// GetStorageClass will get storage-class value from metadata.
func GetStorageClass(m info.ObjectMeta) (string, bool) {
	v, ok := m.Get(InfoObjectMetaStorageClass)
	if !ok {
		return "", false
	}
	return v.(string), true
}

// setstorage-class will set storage-class value into metadata.
func setStorageClass(m info.ObjectMeta, v string) info.ObjectMeta {
	return m.Set(InfoObjectMetaStorageClass, v)
}

var pairServiceCreateMap = map[string]struct{}{
	// Required pairs
	ps.Location: struct{}{},
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairServiceCreate struct {
	// Required pairs
	Location string
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Service) parsePairCreate(opts ...*types.Pair) (*pairServiceCreate, error) {
	result := &pairServiceCreate{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceCreateMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.Location]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Location)
	}
	if ok {
		result.Location = v.(string)
	}
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairServiceDeleteMap = map[string]struct{}{
	// Required pairs
	ps.Location: struct{}{},
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairServiceDelete struct {
	// Required pairs
	Location string
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Service) parsePairDelete(opts ...*types.Pair) (*pairServiceDelete, error) {
	result := &pairServiceDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceDeleteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.Location]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Location)
	}
	if ok {
		result.Location = v.(string)
	}
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairServiceGetMap = map[string]struct{}{
	// Required pairs
	ps.Location: struct{}{},
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairServiceGet struct {
	// Required pairs
	Location string
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Service) parsePairGet(opts ...*types.Pair) (*pairServiceGet, error) {
	result := &pairServiceGet{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceGetMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.Location]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Location)
	}
	if ok {
		result.Location = v.(string)
	}
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairServiceListMap = map[string]struct{}{
	// Required pairs
	ps.StoragerFunc: struct{}{},
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairServiceList struct {
	// Required pairs
	StoragerFunc storage.StoragerFunc
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Service) parsePairList(opts ...*types.Pair) (*pairServiceList, error) {
	result := &pairServiceList{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairServiceListMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.StoragerFunc]
	if !ok {
		return nil, services.NewPairRequiredError(ps.StoragerFunc)
	}
	if ok {
		result.StoragerFunc = v.(storage.StoragerFunc)
	}
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairServiceNewMap = map[string]struct{}{
	// Required pairs
	ps.Credential: struct{}{},
	// Optional pairs
	ps.HTTPClientOptions: struct{}{},
	// Generated pairs
	ps.Context: struct{}{},
}

type pairServiceNew struct {
	// Required pairs
	Credential *credential.Provider
	// Optional pairs
	HasHTTPClientOptions bool
	HTTPClientOptions    *httpclient.Options
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func parseServicePairNew(opts ...*types.Pair) (*pairServiceNew, error) {
	result := &pairServiceNew{}

	values := make(map[string]interface{})
	for _, v := range opts {
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.Credential]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Credential)
	}
	if ok {
		result.Credential = v.(*credential.Provider)
	}
	// Handle optional pairs
	v, ok = values[ps.HTTPClientOptions]
	if ok {
		result.HasHTTPClientOptions = true
		result.HTTPClientOptions = v.(*httpclient.Options)
	}
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

// CreateWithContext adds context support for Create.
func (s *Service) CreateWithContext(ctx context.Context, name string, pairs ...*types.Pair) (st storage.Storager, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.service.Create")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Create(name, pairs...)
}

// DeleteWithContext adds context support for Delete.
func (s *Service) DeleteWithContext(ctx context.Context, name string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.service.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(name, pairs...)
}

// GetWithContext adds context support for Get.
func (s *Service) GetWithContext(ctx context.Context, name string, pairs ...*types.Pair) (st storage.Storager, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.service.Get")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Get(name, pairs...)
}

// ListWithContext adds context support for List.
func (s *Service) ListWithContext(ctx context.Context, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.service.List")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.List(pairs...)
}

var pairStorageDeleteMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairStorageDelete struct {
	// Required pairs
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairDelete(opts ...*types.Pair) (*pairStorageDelete, error) {
	result := &pairStorageDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageDeleteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageListDirMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	ps.DirFunc:  struct{}{},
	ps.FileFunc: struct{}{},
	// Generated pairs
	ps.Context: struct{}{},
}

type pairStorageListDir struct {
	// Required pairs
	// Optional pairs
	HasDirFunc  bool
	DirFunc     types.ObjectFunc
	HasFileFunc bool
	FileFunc    types.ObjectFunc
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairListDir(opts ...*types.Pair) (*pairStorageListDir, error) {
	result := &pairStorageListDir{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageListDirMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	v, ok = values[ps.DirFunc]
	if ok {
		result.HasDirFunc = true
		result.DirFunc = v.(types.ObjectFunc)
	}
	v, ok = values[ps.FileFunc]
	if ok {
		result.HasFileFunc = true
		result.FileFunc = v.(types.ObjectFunc)
	}
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageListPrefixMap = map[string]struct{}{
	// Required pairs
	ps.ObjectFunc: struct{}{},
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairStorageListPrefix struct {
	// Required pairs
	ObjectFunc types.ObjectFunc
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairListPrefix(opts ...*types.Pair) (*pairStorageListPrefix, error) {
	result := &pairStorageListPrefix{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageListPrefixMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.ObjectFunc]
	if !ok {
		return nil, services.NewPairRequiredError(ps.ObjectFunc)
	}
	if ok {
		result.ObjectFunc = v.(types.ObjectFunc)
	}
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageMetadataMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairStorageMetadata struct {
	// Required pairs
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairMetadata(opts ...*types.Pair) (*pairStorageMetadata, error) {
	result := &pairStorageMetadata{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageMetadataMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageNewMap = map[string]struct{}{
	// Required pairs
	ps.Location: struct{}{},
	ps.Name:     struct{}{},
	// Optional pairs
	ps.WorkDir: struct{}{},
	// Generated pairs
	ps.Context: struct{}{},
}

type pairStorageNew struct {
	// Required pairs
	Location string
	Name     string
	// Optional pairs
	HasWorkDir bool
	WorkDir    string
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func parseStoragePairNew(opts ...*types.Pair) (*pairStorageNew, error) {
	result := &pairStorageNew{}

	values := make(map[string]interface{})
	for _, v := range opts {
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.Location]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Location)
	}
	if ok {
		result.Location = v.(string)
	}
	v, ok = values[ps.Name]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Name)
	}
	if ok {
		result.Name = v.(string)
	}
	// Handle optional pairs
	v, ok = values[ps.WorkDir]
	if ok {
		result.HasWorkDir = true
		result.WorkDir = v.(string)
	}
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageReadMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
	ps.ReadCallbackFunc: struct{}{},
	ps.Context:          struct{}{},
}

type pairStorageRead struct {
	// Required pairs
	// Optional pairs
	// Generated pairs
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
	HasContext          bool
	Context             context.Context
}

func (s *Storage) parsePairRead(opts ...*types.Pair) (*pairStorageRead, error) {
	result := &pairStorageRead{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageReadMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageStatMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
	ps.Context: struct{}{},
}

type pairStorageStat struct {
	// Required pairs
	// Optional pairs
	// Generated pairs
	HasContext bool
	Context    context.Context
}

func (s *Storage) parsePairStat(opts ...*types.Pair) (*pairStorageStat, error) {
	result := &pairStorageStat{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageStatMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

var pairStorageWriteMap = map[string]struct{}{
	// Required pairs
	ps.Size: struct{}{},
	// Optional pairs
	ps.Checksum:      struct{}{},
	PairStorageClass: struct{}{},
	// Generated pairs
	ps.ReadCallbackFunc: struct{}{},
	ps.Context:          struct{}{},
}

type pairStorageWrite struct {
	// Required pairs
	Size int64
	// Optional pairs
	HasChecksum     bool
	Checksum        string
	HasStorageClass bool
	StorageClass    string
	// Generated pairs
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
	HasContext          bool
	Context             context.Context
}

func (s *Storage) parsePairWrite(opts ...*types.Pair) (*pairStorageWrite, error) {
	result := &pairStorageWrite{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageWriteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	var v interface{}
	var ok bool

	// Handle required pairs
	v, ok = values[ps.Size]
	if !ok {
		return nil, services.NewPairRequiredError(ps.Size)
	}
	if ok {
		result.Size = v.(int64)
	}
	// Handle optional pairs
	v, ok = values[ps.Checksum]
	if ok {
		result.HasChecksum = true
		result.Checksum = v.(string)
	}
	v, ok = values[PairStorageClass]
	if ok {
		result.HasStorageClass = true
		result.StorageClass = v.(string)
	}
	// Handle generated pairs
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}
	v, ok = values[ps.Context]
	if ok {
		result.HasContext = true
		result.Context = v.(context.Context)
	}

	return result, nil
}

// DeleteWithContext adds context support for Delete.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.storage.Delete")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Delete(path, pairs...)
}

// ListDirWithContext adds context support for ListDir.
func (s *Storage) ListDirWithContext(ctx context.Context, path string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.storage.ListDir")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListDir(path, pairs...)
}

// ListPrefixWithContext adds context support for ListPrefix.
func (s *Storage) ListPrefixWithContext(ctx context.Context, prefix string, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.storage.ListPrefix")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.ListPrefix(prefix, pairs...)
}

// MetadataWithContext adds context support for Metadata.
func (s *Storage) MetadataWithContext(ctx context.Context, pairs ...*types.Pair) (m info.StorageMeta, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.storage.Metadata")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Metadata(pairs...)
}

// ReadWithContext adds context support for Read.
func (s *Storage) ReadWithContext(ctx context.Context, path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.storage.Read")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Read(path, pairs...)
}

// StatWithContext adds context support for Stat.
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...*types.Pair) (o *types.Object, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.storage.Stat")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Stat(path, pairs...)
}

// WriteWithContext adds context support for Write.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, pairs ...*types.Pair) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "github.com/Xuanwo/storage/services/cos.storage.Write")
	defer span.Finish()

	pairs = append(pairs, ps.WithContext(ctx))
	return s.Write(path, r, pairs...)
}
