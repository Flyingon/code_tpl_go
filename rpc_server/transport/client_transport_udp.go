package transport

import (
	"context"
	"net"

	"code_tpl_go/rpc_server/errs"
)

// udpRoundTrip 发送udp请求
func (c *clientTransport) udpRoundTrip(ctx context.Context, reqData []byte,
	opts *RoundTripOptions) (rspData []byte, err error) {

	addr, err := net.ResolveUDPAddr(opts.Network, opts.Address)
	if err != nil {
		return nil, errs.NewFrameError(errs.RetClientNetErr, "udp client transport ResolveUDPAddr: "+err.Error())
	}

	var conn net.PacketConn

	if opts.ConnectionMode == Connected {
		conn, err = net.DialUDP(opts.Network, nil, addr)
	} else {
		conn, err = net.ListenPacket(opts.Network, ":")
	}

	if err != nil {
		return nil, errs.NewFrameError(errs.RetClientNetErr, "udp client transport Dial: "+err.Error())
	}
	defer conn.Close()

	d, ok := ctx.Deadline()
	if ok {
		conn.SetDeadline(d)
	}

	// 发包
	var num int
	if opts.ConnectionMode == Connected {
		udpconn := conn.(*net.UDPConn)
		num, err = udpconn.Write(reqData)
	} else {
		num, err = conn.WriteTo(reqData, addr)
	}
	if err != nil {
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, errs.NewFrameError(errs.RetClientTimeout, "udp client transport WriteTo: "+err.Error())
		}
		return nil, errs.NewFrameError(errs.RetClientNetErr, "udp client transport WriteTo: "+err.Error())
	}
	if num != len(reqData) {
		return nil, errs.NewFrameError(errs.RetClientNetErr, "udp client transport WriteTo: num mismatch")
	}

	// 只发不收
	if opts.ReqType == SendOnly {
		return nil, nil
	}

	select {
	case <-ctx.Done():
		return nil, errs.NewFrameError(errs.RetClientTimeout, "udp client transport select after Write: "+ctx.Err().Error())
	default:
	}

	// 收包
	recvData := make([]byte, 64*1024)

	num, _, err = conn.ReadFrom(recvData)
	if err != nil {
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, errs.NewFrameError(errs.RetClientTimeout, "udp client transport ReadFrom: "+err.Error())
		}
		return nil, errs.NewFrameError(errs.RetClientNetErr, "udp client transport ReadFrom: "+err.Error())
	}
	if num == 0 {
		return nil, errs.NewFrameError(errs.RetClientNetErr, "udp client transport ReadFrom: num empty")
	}

	return recvData[:num], nil
}
