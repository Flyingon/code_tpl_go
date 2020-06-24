package transport

import (
	"git.code.oa.com/trpc-go/trpc-go/codec"
)

// ListenServeOptions server每次启动参数
type ListenServeOptions struct {
	Address       string
	Network       string
	Checker       Checker
	Handler       Handler
	FramerBuilder codec.FramerBuilder

	CACertFile  string // ca证书
	TLSCertFile string // server证书
	TLSKeyFile  string // server秘钥
}

// ListenServeOption function type for config listenServeOptions
type ListenServeOption func(*ListenServeOptions)

// WithServerFramerBuilder 设置FramerBuilder
func WithServerFramerBuilder(fb codec.FramerBuilder) ListenServeOption {
	return func(opts *ListenServeOptions) {
		opts.FramerBuilder = fb
	}
}

// WithListenAddress 设置ListenAddress
func WithListenAddress(address string) ListenServeOption {
	return func(opts *ListenServeOptions) {
		opts.Address = address
	}
}

// WithListenNetwork 设置ListenNetwork
func WithListenNetwork(network string) ListenServeOption {
	return func(opts *ListenServeOptions) {
		opts.Network = network
	}
}

// WithServerChecker 设置ServerChecker
func WithServerChecker(checker Checker) ListenServeOption {
	return func(opts *ListenServeOptions) {
		opts.Checker = checker
	}
}

// WithHandler 设置业务处理抽象接口Handler
func WithHandler(handler Handler) ListenServeOption {
	return func(opts *ListenServeOptions) {
		opts.Handler = handler
	}
}

// WithServeTLS 设置服务支持TLS
func WithServeTLS(certFile, keyFile, caFile string) ListenServeOption {
	return func(opts *ListenServeOptions) {
		opts.TLSCertFile = certFile
		opts.TLSKeyFile = keyFile
		opts.CACertFile = caFile
	}
}
