package credential

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrInvalidConfig will return if config string invalid.
	ErrInvalidConfig = errors.New("invalid config")
	// ErrUnsupportedProtocol will return if protocol is unsupported.
	ErrUnsupportedProtocol = errors.New("unsupported protocol")
)

const (
	// ProtocolHmac will hold access key and secret key credential.
	//
	// HMAC means hash-based message authentication code, it may be inaccurate to represent credential
	// protocol ak/sk(access key + secret key with hmac), but it's simple and no confuse with other
	// protocol, so just keep this.
	//
	// value = [Access Key, Secret Key]
	ProtocolHmac = "hmac"
	// ProtocolAPIKey will hold api key credential.
	//
	// value = [API Key]
	ProtocolAPIKey = "apikey"
	// ProtocolFile will hold file credential.
	//
	// value = [File Path], service decide how to use this file
	ProtocolFile = "file"
	// ProtocolEnv will represent credential from env.
	//
	// value = [], service retrieves credential value from env.
	ProtocolEnv = "env"
)

// Provider will provide credential protocol and values.
type Provider struct {
	protocol string
	args     []string
}

// Protocol provides current credential's protocol.
func (p *Provider) Protocol() string {
	return p.protocol
}

// Value provides current credential's value in string array.
func (p *Provider) Value() []string {
	return p.args
}

// Parse will parse config string to create a credential Provider.
func Parse(cfg string) (*Provider, error) {
	errorMessage := "parse credential config [%s]: %w"

	s := strings.Split(cfg, ":")

	switch s[0] {
	case ProtocolHmac:
		return NewHmac(s[1:]...)
	case ProtocolAPIKey:
		return NewAPIKey(s[1:]...)
	case ProtocolFile:
		return NewFile(s[1:]...)
	case ProtocolEnv:
		return NewEnv()
	default:
		return nil, fmt.Errorf(errorMessage, cfg, ErrUnsupportedProtocol)
	}
}

// NewHmac create a hmac provider.
func NewHmac(value ...string) (*Provider, error) {
	errorMessage := "parse hmac credential [%s]: %w"

	if len(value) != 2 {
		return nil, fmt.Errorf(errorMessage, value, ErrInvalidConfig)
	}
	return &Provider{ProtocolHmac, []string{value[0], value[1]}}, nil
}

// MustNewHmac make sure Provider must be created if no panic happened.
func MustNewHmac(value ...string) *Provider {
	p, err := NewHmac(value...)
	if err != nil {
		panic(err)
	}
	return p
}

// NewAPIKey create a api key provider.
func NewAPIKey(value ...string) (*Provider, error) {
	errorMessage := "parse apikey credential [%s]: %w"

	if len(value) != 1 {
		return nil, fmt.Errorf(errorMessage, value, ErrInvalidConfig)
	}
	return &Provider{ProtocolAPIKey, []string{value[0]}}, nil
}

// MustNewAPIKey make sure Provider must be created if no panic happened.
func MustNewAPIKey(value ...string) *Provider {
	p, err := NewAPIKey(value...)
	if err != nil {
		panic(err)
	}
	return p
}

// NewFile create a file provider.
func NewFile(value ...string) (*Provider, error) {
	errorMessage := "parse file credential [%s]: %w"

	if len(value) != 1 {
		return nil, fmt.Errorf(errorMessage, value, ErrInvalidConfig)
	}
	return &Provider{ProtocolFile, []string{value[0]}}, nil
}

// MustNewFile make sure Provider must be created if no panic happened.
func MustNewFile(value ...string) *Provider {
	p, err := NewFile(value...)
	if err != nil {
		panic(err)
	}
	return p
}

// NewEnv create a env provider.
func NewEnv(_ ...string) (*Provider, error) {
	return &Provider{ProtocolEnv, nil}, nil
}

// MustNewEnv make sure Provider must be created if no panic happened.
func MustNewEnv(value ...string) *Provider {
	p, _ := NewEnv(value...)
	return p
}
