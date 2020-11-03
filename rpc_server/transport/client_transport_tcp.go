package transport

import (
	"context"
	"net"
	"time"

	"code_tpl_go/rpc_server/codec"
	"code_tpl_go/rpc_server/errs"
	"code_tpl_go/rpc_server/pool/connpool"
)

// tcpRoundTrip 发送tcp请求 支持 1.send 2. sendAndRcv 3. keepalive 4. multiplex
func (c *clientTransport) tcpRoundTrip(ctx context.Context, reqData []byte,
	opts *RoundTripOptions) (rspData []byte, err error) {

	var timeout time.Duration
	d, ok := ctx.Deadline()
	if ok {
		timeout = d.Sub(time.Now())
	}

	if opts.Pool == nil {
		return nil, errs.NewFrameError(errs.RetClientConnectFail,
			"tcp client transport: connection pool empty")
	}

	if opts.FramerBuilder == nil {
		return nil, errs.NewFrameError(errs.RetClientConnectFail,
			"tcp client transport: framer builder empty")
	}

	// 从连接池中获取连接
	conn, err := opts.Pool.Get(opts.Network, opts.Address, timeout, connpool.WithFramerBuilder(opts.FramerBuilder))
	if err != nil {
		return nil, errs.NewFrameError(errs.RetClientConnectFail,
			"tcp client transport connection pool: "+err.Error())
	}
	defer conn.Close()

	if ok {
		conn.SetDeadline(d)
	}

	if ctx.Err() == context.Canceled {
		return nil, errs.NewFrameError(errs.RetClientCanceled,
			"tcp client transport canceled before Write: "+ctx.Err().Error())
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errs.NewFrameError(errs.RetClientTimeout,
			"tcp client transport timeout before Write: "+ctx.Err().Error())
	}

	// 循环发包
	sentNum := 0
	num := 0
	for sentNum < len(reqData) {

		num, err = conn.Write(reqData[sentNum:])
		if err != nil {
			if e, ok := err.(net.Error); ok && e.Timeout() {
				return nil, errs.NewFrameError(errs.RetClientTimeout,
					"tcp client transport Write: "+err.Error())
			}
			return nil, errs.NewFrameError(errs.RetClientNetErr,
				"tcp client transport Write: "+err.Error())
		}

		sentNum += num

		if ctx.Err() == context.Canceled {
			return nil, errs.NewFrameError(errs.RetClientCanceled,
				"tcp client transport canceled after Write: "+ctx.Err().Error())
		}
		if ctx.Err() == context.DeadlineExceeded {
			return nil, errs.NewFrameError(errs.RetClientTimeout,
				"tcp client transport timeout after Write: "+ctx.Err().Error())
		}
	}

	// 只发不收
	if opts.ReqType == SendOnly {
		return nil, nil
	}

	fr, ok := conn.(codec.Framer)
	if !ok {
		return nil, errs.NewFrameError(errs.RetClientConnectFail,
			"tcp client transport: framer not implemented")
	}
	rspData, err = fr.ReadFrame()
	if err != nil {
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, errs.NewFrameError(errs.RetClientTimeout,
				"tcp client transport ReadFrame: "+err.Error())
		}
		return nil, errs.NewFrameError(errs.RetClientNetErr,
			"tcp client transport ReadFrame: "+err.Error())
	}

	return rspData, nil
}
