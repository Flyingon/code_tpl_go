package main

import (
	"context"
	"encoding/json"
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

func callUic(text string) {
	udpOpts := UDPCallOpts{
		Addr:          "100.117.37.58:20001",
		ConnectedMode: false,
		NetWork:       "udp",
		SendOnly:      false,
	}
	reqMap := make(map[string]interface{})
	reqHead := map[string]interface{}{
		"account_":       "",
		"game_id_":       2264,
		"plat_id_":       0,
		"world_":         0,
		"busi_passwd_":   "ODKEg21cMRNGnIfi6y39",
		"cmd_":           2,
		"callback_data_": "",
		"err_code_":      0,
		"err_msg_":       "",
	}
	reqBody := map[string]map[string]interface{}{
		"text_pkg_": {
			"text_category_": 5,
			"text_":          text,
			"get_sens_words": 1,
		},
	}
	reqMap["head_"] = reqHead
	reqMap["body_"] = reqBody

	reqJson, _ := json.Marshal(reqMap)
	rsp, err := udpOpts.UDPCall(context.Background(), reqJson)
	fmt.Printf("uic call rsp: %s, err: %v\n", rsp, err)
}

func main() {
	callUic("习近平，跑步")
}
