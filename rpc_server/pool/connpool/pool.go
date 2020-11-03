// Package connpool 连接池
package connpool

import (
	"net"
	"time"

	"code_tpl_go/rpc_server/codec"
)

// GetOptions get conn configuration
type GetOptions struct {
	FramerBuilder codec.FramerBuilder
}

// GetOption Options helper
type GetOption func(*GetOptions)

// WithFramerBuilder 设置 FramerBuilder
func WithFramerBuilder(fb codec.FramerBuilder) GetOption {
	return func(opts *GetOptions) {
		opts.FramerBuilder = fb
	}
}

// Pool client connection pool
type Pool interface {
	Get(network string, address string, timeout time.Duration, opt ...GetOption) (net.Conn, error)
}
