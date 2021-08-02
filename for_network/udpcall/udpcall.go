package main

import (
	"context"
	"fmt"
	"net"
)

type UDPCallOpts struct {
	Addr          string
	ConnectedMode bool
	NetWork       string
	SendOnly      bool
}

func (u *UDPCallOpts) UDPCall(ctx context.Context, reqData []byte) ([]byte, error) {
	var err error
	addr, err := net.ResolveUDPAddr(u.NetWork, u.Addr)
	if err != nil {
		return nil, err
	}

	fmt.Println("!!!!!!", addr, u.ConnectedMode)
	fmt.Printf("?????%s\n", reqData)

	var conn net.PacketConn

	if !u.ConnectedMode {
		conn, err = net.DialUDP(u.NetWork, nil, addr)
	} else {
		conn, err = net.ListenPacket(u.NetWork, ":")
	}

	if err != nil {
		return nil, fmt.Errorf("udp client transport Dial: " + err.Error())
	}
	defer conn.Close()

	d, ok := ctx.Deadline()
	if ok {
		conn.SetDeadline(d)
	}

	// 发包
	var num int
	if !u.ConnectedMode {
		udpconn := conn.(*net.UDPConn)
		num, err = udpconn.Write(reqData)
	} else {
		num, err = conn.WriteTo(reqData, addr)
	}
	if err != nil {
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, fmt.Errorf("udp client transport WriteTo: " + err.Error())
		}
		return nil, fmt.Errorf("udp client transport WriteTo: " + err.Error())
	}
	if num != len(reqData) {
		return nil, fmt.Errorf("udp client transport WriteTo: num mismatch")
	}

	// 只发不收
	if u.SendOnly {
		return nil, nil
	}

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("udp client transport select after Write: " + ctx.Err().Error())
	default:
	}

	// 收包
	recvData := make([]byte, 64*1024)

	num, _, err = conn.ReadFrom(recvData)
	if err != nil {
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, fmt.Errorf("udp client transport ReadFrom: " + err.Error())
		}
		return nil, fmt.Errorf("udp client transport ReadFrom: " + err.Error())
	}
	if num == 0 {
		return nil, fmt.Errorf("udp client transport ReadFrom: num empty")
	}

	return recvData[:num], nil
}

func main() {

}
