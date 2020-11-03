package transport

import (
	"context"
	"errors"
	"io"
	"net"
	"time"

	"code_tpl_go/rpc_server/codec"
	"code_tpl_go/rpc_server/pool/objectpool"
	"code_tpl_go/rpc_server/log"
	"code_tpl_go/rpc_server/metrics"
)

var defaultRecvBufSize = 4096
var bytesPool = objectpool.NewBytesPool(defaultRecvBufSize)

func (s *serverTransport) serveTCP(ctx context.Context, ln *net.TCPListener,
	opts *ListenServeOptions) error {

	defer ln.Close()

	var tempDelay time.Duration
	for {

		select {
		case <-ctx.Done():
			return errors.New("recv server close event")
		default:
		}

		rwc, err := ln.AcceptTCP()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
		tempDelay = 0

		err = rwc.SetKeepAlive(true)
		if err != nil {
			log.Tracef("tcp conn set keepalive error:%v", err)
		}

		if s.opts.KeepAlivePeriod > 0 {
			err = rwc.SetKeepAlivePeriod(s.opts.KeepAlivePeriod)
			if err != nil {
				log.Tracef("tcp conn set keepalive period error:%v", err)
			}
		}

		tc := &tcpconn{
			conn:       s.newConn(ctx, opts),
			rwc:        rwc,
			fr:         opts.FramerBuilder.New(rwc),
			remoteAddr: rwc.RemoteAddr(),
			localAddr:  rwc.LocalAddr(),
		}

		go tc.serve()
	}
}

type tcpconn struct {
	*conn
	rwc *net.TCPConn
	fr  codec.Framer

	localAddr  net.Addr
	remoteAddr net.Addr
}

func (c *tcpconn) serve() {
	defer c.rwc.Close()

	for {

		// 检查上游是否关闭
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		if c.idleTimeout > 0 {
			now := time.Now()
			if now.Sub(c.lastVisited) > 5*time.Second { // SetReadDeadline性能损耗较严重，每5s才更新一次timeout
				c.lastVisited = now
				err := c.rwc.SetReadDeadline(now.Add(c.idleTimeout))
				if err != nil {
					log.Trace("transport: tcpconn SetReadDeadline fail ", err)
					return
				}
			}
		}

		req, err := c.fr.ReadFrame()
		if err != nil {
			if err == io.EOF {
				metrics.Counter("TcpServerTransportReadEOF").Incr() // 客户端主动断开连接
				return
			}
			if e, ok := err.(net.Error); ok && e.Timeout() { // 客户端超过空闲时间没有发包，服务端主动超时关闭
				metrics.Counter("TcpServerTransportIdleTimeout").Incr()
				return
			}
			metrics.Counter("TcpServerTransportReadFail").Incr()
			log.Trace("transport: tcpconn serve ReadFrame fail ", err)
			return
		}

		// 生成新的空的通用消息结构数据，并保存到ctx里面
		ctx, msg := codec.WithNewMessage(context.Background())

		// 记录LocalAddr和RemoteAddr到Context
		msg.WithLocalAddr(c.localAddr)
		msg.WithRemoteAddr(c.remoteAddr)

		rsp, err := c.handle(ctx, req)
		if err != nil {
			metrics.Counter("TcpServerTransportHandleFail").Incr()
			log.Trace("transport: tcpconn serve handle fail ", err)
			return
		}
		codec.PutBackMessage(msg)

		if _, err = c.rwc.Write(rsp); err != nil {
			metrics.Counter("TcpServerTransportWriteFail").Incr()
			log.Trace("transport: tcpconn serve Write fail ", err)
			return
		}
	}
}
