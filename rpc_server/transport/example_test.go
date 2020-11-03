package transport_test

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"code_tpl_go/rpc_server/codec"
	"code_tpl_go/rpc_server/transport"
)

func clientInvokeServer(network string) {
	go func() {
		err := transport.ListenAndServe(
			transport.WithListenNetwork(network),
			transport.WithListenAddress(":8888"),
			transport.WithHandler(&simpleHandler{}),
			transport.WithServerFramerBuilder(&framerBuilder{}),
		)

		if err != nil {
			log.Fatalln(err)
		}
	}()

	time.Sleep(time.Millisecond * 10)

	ctx, f := context.WithTimeout(context.Background(), 3*time.Second)
	defer f()
	req := &helloRequest{
		Name: "trpc",
		Msg:  "HelloWorld",
	}

	data, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
	}
	lenData := make([]byte, 4)
	binary.BigEndian.PutUint32(lenData, uint32(len(data)))

	reqData := append(lenData, data...)

	rspData, err := transport.RoundTrip(ctx, reqData, transport.WithDialNetwork(network),
		transport.WithDialAddress(":8888"),
		transport.WithClientFramerBuilder(&framerBuilder{}))
	if nil != err {
		log.Fatalf("RoundTip Error : %v", err)
	}

	length := binary.BigEndian.Uint32(rspData[:4])

	helloRsp := &helloResponse{}
	err = json.Unmarshal(rspData[4:4+length], helloRsp)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(helloRsp)

}

func UDP4TransportExample() {
	clientInvokeServer("udp4")
	// Output:
	// &{trpc HelloWorld 0}
}

func TCPTransportExample() {
	clientInvokeServer("tcp")
	// Output:
	// &{trpc HelloWorld 0}
}

type helloRequest struct {
	Name string
	Msg  string
}

type helloResponse struct {
	Name string
	Msg  string
	Code int
}

type framerBuilder struct{}

func (fb *framerBuilder) New(r io.Reader) codec.Framer {
	return &framer{r: r}
}

type framer struct {
	r io.Reader
}

func (f *framer) ReadFrame() ([]byte, error) {

	var lenData [4]byte

	_, err := io.ReadFull(f.r, lenData[:])
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lenData[:])

	msg := make([]byte, len(lenData)+int(length))
	copy(msg, lenData[:])

	_, err = io.ReadFull(f.r, msg[len(lenData):])
	if err != nil {
		return nil, err
	}

	return msg, nil
}

type simpleMsgReader struct{}

func (r *simpleMsgReader) ReadMsg(reader io.Reader) ([]byte, error) {
	var lenData [4]byte

	_, err := io.ReadFull(reader, lenData[:])
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lenData[:])

	msg := make([]byte, len(lenData)+int(length))
	copy(msg, lenData[:])

	_, err = io.ReadFull(reader, msg[len(lenData):])
	if err != nil {
		return nil, err
	}

	return msg, nil
}

type errorHandler struct{}

func (h *errorHandler) Handle(ctx context.Context, req []byte) ([]byte, error) {
	return nil, errors.New("handle error")
}

type simpleHandler struct{}

func (h *simpleHandler) Handle(ctx context.Context, reqdata []byte) ([]byte, error) {

	helloReq := &helloRequest{}
	helloRsp := &helloResponse{}

	if len(reqdata) < 4 {
		return nil, errors.New("reqData format error")
	}

	json.Unmarshal(reqdata[4:], helloReq)

	helloRsp.Name = helloReq.Name
	helloRsp.Msg = helloReq.Msg
	data, _ := json.Marshal(helloRsp)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint32(len(data)))
	binary.Write(buf, binary.BigEndian, data)
	return buf.Bytes(), nil
}

type echoHandler struct{}

func (h *echoHandler) Handle(ctx context.Context, req []byte) ([]byte, error) {
	rsp := make([]byte, len(req))
	copy(rsp, req)
	return rsp, nil
}
