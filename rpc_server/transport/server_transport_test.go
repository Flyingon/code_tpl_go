package transport_test

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	_ "git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/transport"
)

func TestNewServerTransport(t *testing.T) {
	st := transport.NewServerTransport(transport.WithKeepAlivePeriod(time.Minute))
	assert.NotNil(t, st)
}

func TestTCPListenAndServe(t *testing.T) {
	go func() {
		st := transport.NewServerTransport(transport.WithKeepAlivePeriod(time.Minute))
		err := st.ListenAndServe(context.Background(),
			transport.WithListenNetwork("tcp4"),
			transport.WithListenAddress(":12000"),
			transport.WithHandler(&errorHandler{}),
			transport.WithServerFramerBuilder(&framerBuilder{}),
		)

		if err != nil {
			t.Fatalf("ListenAndServe fail:%v", err)
		}
	}()
	ctx, f := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer f()
	req := &helloRequest{
		Name: "trpc",
		Msg:  "HelloWorld",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("json marshal fail:%v", err)
	}
	lenData := make([]byte, 4)
	binary.BigEndian.PutUint32(lenData, uint32(len(data)))

	reqData := append(lenData, data...)

	time.Sleep(time.Millisecond * 10)
	_, err = transport.RoundTrip(ctx, reqData, transport.WithDialNetwork("tcp4"),
		transport.WithDialAddress(":12000"),
		transport.WithClientFramerBuilder(&framerBuilder{}))
	assert.NotNil(t, err)
}

func TestHandleError(t *testing.T) {
	go func() {
		err := transport.ListenAndServe(
			transport.WithListenNetwork("udp4"),
			transport.WithListenAddress(":8890"),
			transport.WithHandler(&errorHandler{}),
			transport.WithServerFramerBuilder(&framerBuilder{}),
		)

		if err != nil {
			t.Fatalf("test fail:%v", err)
		}
	}()
	ctx, f := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer f()
	req := &helloRequest{
		Name: "trpc",
		Msg:  "HelloWorld",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("test fail:%v", err)
	}
	lenData := make([]byte, 4)
	binary.BigEndian.PutUint32(lenData, uint32(len(data)))

	reqData := append(lenData, data...)

	_, err = transport.RoundTrip(ctx, reqData, transport.WithDialNetwork("udp4"),
		transport.WithDialAddress(":8890"),
		transport.WithClientFramerBuilder(&framerBuilder{}))
	assert.NotNil(t, err)
}

func TestNewServerTransport_NotSupport(t *testing.T) {
	st := transport.NewServerTransport()
	err := st.ListenAndServe(context.Background(), transport.WithListenNetwork("unix"))
	assert.NotNil(t, err)

	err = st.ListenAndServe(context.Background(), transport.WithListenNetwork("xxx"))
	assert.NotNil(t, err)
}

func TestServerTransport_ListenAndServeUDP(t *testing.T) {
	// NoReusePort
	st := transport.NewServerTransport(transport.WithReusePort(false),
		transport.WithKeepAlivePeriod(time.Minute))
	err := st.ListenAndServe(context.Background(), transport.WithListenNetwork("udp"))
	assert.Nil(t, err)

	st = transport.NewServerTransport(transport.WithReusePort(true))
	err = st.ListenAndServe(context.Background(), transport.WithListenNetwork("udp"))
	assert.Nil(t, err)

	st = transport.NewServerTransport(transport.WithReusePort(true))
	err = st.ListenAndServe(context.Background(), transport.WithListenNetwork("ip"))
	assert.NotNil(t, err)

	st = transport.NewServerTransport(transport.WithReusePort(true))
	err = st.ListenAndServe(context.Background(), transport.WithListenNetwork("unix"))
	assert.NotNil(t, err)
}

func TestServerTransport_ListenAndServe(t *testing.T) {
	// NoFramerBuilder
	st := transport.NewServerTransport(transport.WithReusePort(false))
	err := st.ListenAndServe(context.Background(), transport.WithListenNetwork("tcp"))
	assert.NotNil(t, err)

	fb := transport.GetFramerBuilder("trpc")
	// NoReusePort
	st = transport.NewServerTransport(transport.WithReusePort(false))
	err = st.ListenAndServe(context.Background(),
		transport.WithListenNetwork("tcp"),
		transport.WithServerFramerBuilder(fb))
	assert.Nil(t, err)

	// ReusePort
	st = transport.NewServerTransport(transport.WithReusePort(true))
	err = st.ListenAndServe(context.Background(),
		transport.WithListenNetwork("tcp"),
		transport.WithServerFramerBuilder(fb))
	assert.Nil(t, err)

	// ReusePort + Listen Error
	st = transport.NewServerTransport(transport.WithReusePort(true))
	err = st.ListenAndServe(context.Background(),
		transport.WithListenNetwork("tcperror"),
		transport.WithServerFramerBuilder(fb))
	assert.NotNil(t, err)

	// context cancel
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	st = transport.NewServerTransport(transport.WithReusePort(true))
	err = st.ListenAndServe(ctx, transport.WithListenNetwork("tcp"), transport.WithServerFramerBuilder(fb))
	assert.Nil(t, err)
}

func TestWithReusePort(t *testing.T) {

	opt := transport.WithReusePort(false)
	assert.NotNil(t, opt)
	opts := &transport.ServerTransportOptions{}
	opt(opts)
	assert.Equal(t, false, opts.ReusePort)
}

func TestWithRecvMsgChannelSize(t *testing.T) {
	opt := transport.WithRecvMsgChannelSize(1000)
	assert.NotNil(t, opt)
	opts := &transport.ServerTransportOptions{}
	opt(opts)
	assert.Equal(t, 1000, opts.RecvMsgChannelSize)
}

func TestWithSendMsgChannelSize(t *testing.T) {
	opt := transport.WithSendMsgChannelSize(1000)
	assert.NotNil(t, opt)
	opts := &transport.ServerTransportOptions{}
	opt(opts)
	assert.Equal(t, 1000, opts.SendMsgChannelSize)
}

func TestWithRecvUDPPacketBufferSize(t *testing.T) {
	opt := transport.WithRecvUDPPacketBufferSize(1000)
	assert.NotNil(t, opt)
	opts := &transport.ServerTransportOptions{}
	opt(opts)
	assert.Equal(t, 1000, opts.RecvUDPPacketBufferSize)
}

func TestWithIdleTimeout(t *testing.T) {
	opt := transport.WithIdleTimeout(time.Second)
	assert.NotNil(t, opt)
	opts := &transport.ServerTransportOptions{}
	opt(opts)
	assert.Equal(t, time.Second, opts.IdleTimeout)
}

func TestWithKeepAlivePeriod(t *testing.T) {
	opt := transport.WithKeepAlivePeriod(time.Minute)
	assert.NotNil(t, opt)
	opts := &transport.ServerTransportOptions{}
	opt(opts)
	assert.Equal(t, time.Minute, opts.KeepAlivePeriod)
}

func TestWithServeTLS(t *testing.T) {
	opt := transport.WithServeTLS("certfile", "keyfile", "")
	assert.NotNil(t, opt)
	opts := &transport.ListenServeOptions{}
	opt(opts)
	assert.Equal(t, "certfile", opts.TLSCertFile)
	assert.Equal(t, "keyfile", opts.TLSKeyFile)
}

func TestWithServerChecker(t *testing.T) {
	opt := transport.WithServerChecker(nil)
	assert.NotNil(t, opt)
	opts := &transport.ListenServeOptions{}
	opt(opts)
	assert.Equal(t, nil, opts.Checker)
}
