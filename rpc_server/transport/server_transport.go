package transport

import (
	"context"
	"errors"
	"fmt"
	"net"
	"runtime"
	"time"

	reuseport "code_tpl_go/rpc_server/go_reuseport"
)

// DefaultServerTransport ServerTransport默认实现
var DefaultServerTransport = NewServerTransport(WithReusePort(true))

// NewServerTransport new出来server transport实现
func NewServerTransport(opt ...ServerTransportOption) ServerTransport {

	// option 默认值
	opts := &ServerTransportOptions{
		RecvUDPPacketBufferSize: 65536,
		SendMsgChannelSize:      100,
		RecvMsgChannelSize:      100,
		IdleTimeout:             time.Minute,
	}

	for _, o := range opt {
		o(opts)
	}

	return &serverTransport{opts: opts}
}

// serverTransport server transport具体实现 包括tcp udp serving
type serverTransport struct {
	opts *ServerTransportOptions
}

// ListenAndServe 启动监听，如果监听失败则返回错误
func (s *serverTransport) ListenAndServe(ctx context.Context, opts ...ListenServeOption) error {

	lsopts := &ListenServeOptions{}
	for _, opt := range opts {
		opt(lsopts)
	}

	switch lsopts.Network {
	case "tcp", "tcp4", "tcp6":
		return s.listenAndServeStream(ctx, lsopts)
	case "udp", "udp4", "udp6":
		return s.listenAndServePacket(ctx, lsopts)
	default:
		return fmt.Errorf("server transport: not support network type %s", lsopts.Network)
	}
}

//---------------------------------stream server-----------------------------------------//

// listenAndServeStream 启动监听，如果监听失败则返回错误
func (s *serverTransport) listenAndServeStream(ctx context.Context,
	opts *ListenServeOptions) error {

	if opts.FramerBuilder == nil {
		return errors.New("tcp transport FramerBuilder empty")
	}

	// 端口重用，内核分发IO ReadReady事件到多核多线程，加速IO效率
	if s.opts.ReusePort {
		listener, err := reuseport.Listen(opts.Network, opts.Address)
		if err != nil {
			return fmt.Errorf("tcp reuseport error:%v", err)
		}
		go s.serveStream(ctx, listener, opts)
	} else {
		listener, err := net.Listen(opts.Network, opts.Address)
		if err != nil {
			return err
		}
		go s.serveStream(ctx, listener, opts)
	}

	return nil
}

func (s *serverTransport) serveStream(ctx context.Context, ln net.Listener,
	opts *ListenServeOptions) error {

	switch ln := ln.(type) {
	case *net.TCPListener:
		return s.serveTCP(ctx, ln, opts)
	default:
		return errors.New("transport not support Listener impl")
	}
}

//---------------------------------packet server-----------------------------------------//

// listenAndServePacket 启动监听，如果监听失败则返回错误
func (s *serverTransport) listenAndServePacket(ctx context.Context,
	opts *ListenServeOptions) error {

	// 端口重用，内核分发IO ReadReady事件到多核多线程，加速IO效率
	if s.opts.ReusePort {
		reuseport.ListenerBacklogMaxSize = 4096
		cores := runtime.NumCPU()
		for i := 0; i < cores; i++ {
			udpconn, err := reuseport.ListenPacket(opts.Network, opts.Address)
			if err != nil {
				return fmt.Errorf("udp reuseport error:%v", err)
			}
			go s.servePacket(ctx, udpconn, opts)
		}
	} else {
		udpconn, err := net.ListenPacket(opts.Network, opts.Address)
		if err != nil {
			return err
		}
		go s.servePacket(ctx, udpconn, opts)
	}

	return nil
}

func (s *serverTransport) servePacket(ctx context.Context, rwc net.PacketConn,
	opts *ListenServeOptions) error {

	switch rwc := rwc.(type) {
	case *net.UDPConn:
		return s.serveUDP(ctx, rwc, opts)
	default:
		return errors.New("transport not support PacketConn impl")
	}
}

//------------------------tcp/udp connection通用结构 统一处理----------------------------//

func (s *serverTransport) newConn(ctx context.Context, opts *ListenServeOptions) *conn {

	return &conn{
		ctx:         ctx,
		handler:     opts.Handler,
		checker:     opts.Checker,
		idleTimeout: s.opts.IdleTimeout,
	}
}

type conn struct {
	ctx         context.Context
	cancelCtx   context.CancelFunc
	idleTimeout time.Duration
	lastVisited time.Time

	handler Handler
	checker Checker
}

func (c *conn) handle(ctx context.Context, req []byte) ([]byte, error) {
	return c.handler.Handle(ctx, req)
}
