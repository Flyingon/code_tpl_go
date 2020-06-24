package transport_test

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"git.code.oa.com/trpc-go/trpc-go/pool/connpool"
	"git.code.oa.com/trpc-go/trpc-go/transport"
	"github.com/stretchr/testify/assert"
)

func TestNewClientTransport(t *testing.T) {
	st := transport.NewClientTransport()
	assert.NotNil(t, st)
}

func TestWithDialPool(t *testing.T) {
	opt := transport.WithDialPool(nil)
	opts := &transport.RoundTripOptions{}
	opt(opts)
	assert.Equal(t, nil, opts.Pool)
}

func TestWithClientChecker(t *testing.T) {
	opt := transport.WithClientChecker(nil)
	opts := &transport.RoundTripOptions{}
	opt(opts)
	assert.Equal(t, nil, opts.Checker)
}

func TestWithReqType(t *testing.T) {
	opt := transport.WithReqType(transport.SendOnly)
	opts := &transport.RoundTripOptions{}
	opt(opts)
	assert.Equal(t, transport.SendOnly, opts.ReqType)
}

type emptyPool struct {
}

func (p *emptyPool) Get(network string, address string, timeout time.Duration, opt ...connpool.GetOption) (net.Conn, error) {
	return nil, errors.New("empty")
}

var testReqByte = []byte{'a', 'b'}

func TestWithDialPoolError(t *testing.T) {
	ctx, f := context.WithTimeout(context.Background(), 3*time.Second)
	defer f()
	_, err := transport.RoundTrip(ctx, testReqByte,
		transport.WithDialPool(&emptyPool{}),
		transport.WithDialNetwork("tcp"))
	//fmt.Printf("err: %v", err)
	assert.NotNil(t, err)
}

func TestContextTimeout(t *testing.T) {
	ctx, f := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer f()
	time.Sleep(20 * time.Millisecond)
	_, err := transport.RoundTrip(ctx, testReqByte,
		transport.WithDialNetwork("tcp"),
		transport.WithDialAddress(":8888"))
	//fmt.Printf("err: %v", err)
	assert.NotNil(t, err)
}

func TestWithReqTypeSendOnly(t *testing.T) {
	ctx, f := context.WithTimeout(context.Background(), 3*time.Second)
	defer f()
	_, err := transport.RoundTrip(ctx, []byte{},
		transport.WithReqType(transport.SendOnly),
		transport.WithDialNetwork("tcp"))
	//fmt.Printf("err: %v", err)
	assert.NotNil(t, err)
}

func TestClientTransport_RoundTrip(t *testing.T) {

	go func() {
		err := transport.ListenAndServe(
			transport.WithListenNetwork("udp"),
			transport.WithListenAddress("localhost:9999"),
			transport.WithHandler(&echoHandler{}),
		)
		assert.Nil(t, err)
	}()
	time.Sleep(20 * time.Millisecond)

	var err error
	_, err = transport.RoundTrip(context.Background(), []byte("helloworld"))
	assert.NotNil(t, err)

	tc := transport.NewClientTransport(transport.WithClientUDPRecvSize(4096))
	_, err = tc.RoundTrip(context.Background(), []byte("helloworld"))
	assert.NotNil(t, err)

	// Test Address invalid
	_, err = tc.RoundTrip(context.Background(), []byte("helloworld"),
		transport.WithDialNetwork("udp"),
		transport.WithDialAddress("invalidaddress"),
		transport.WithReqType(transport.SendOnly))
	assert.NotNil(t, err)

	// Test For SendOnly
	rsp, err := tc.RoundTrip(context.Background(), []byte("helloworld"),
		transport.WithDialNetwork("udp"),
		transport.WithDialAddress("localhost:9999"),
		transport.WithReqType(transport.SendOnly),
		transport.WithConnectionMode(transport.NotConnected))
	assert.Nil(t, err)
	assert.Nil(t, rsp)

	// Test Context Done
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = tc.RoundTrip(ctx, []byte("helloworld"),
		transport.WithDialNetwork("udp"),
		transport.WithDialAddress("localhost:9999"))
	assert.NotNil(t, err)

	// Test RoundTrip
	rsp, err = tc.RoundTrip(context.Background(), []byte("helloworld"),
		transport.WithDialNetwork("udp"),
		transport.WithDialAddress("localhost:9999"),
		transport.WithConnectionMode(transport.NotConnected),
	)
	assert.Nil(t, err)
	assert.Equal(t, "helloworld", string(rsp))
}

func TestClientTransport_RoundTrip_PreConnected(t *testing.T) {

	go func() {
		err := transport.ListenAndServe(
			transport.WithListenNetwork("udp"),
			transport.WithListenAddress("localhost:9999"),
			transport.WithHandler(&echoHandler{}),
		)
		assert.Nil(t, err)
	}()
	time.Sleep(20 * time.Millisecond)

	var err error
	_, err = transport.RoundTrip(context.Background(), []byte("helloworld"))
	assert.NotNil(t, err)

	tc := transport.NewClientTransport(transport.WithClientUDPRecvSize(4096))

	// Test For Connected UDPConn
	rsp, err := tc.RoundTrip(context.Background(), []byte("helloworld"),
		transport.WithDialNetwork("udp"),
		transport.WithDialAddress("localhost:9999"),
		transport.WithDialPassword("passwd"),
		transport.WithReqType(transport.SendOnly),
		transport.WithConnectionMode(transport.Connected))
	assert.Nil(t, err)
	assert.Nil(t, rsp)

	// Test Context Done
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = tc.RoundTrip(ctx, []byte("helloworld"),
		transport.WithDialNetwork("udp"),
		transport.WithDialAddress("localhost:9999"),
		transport.WithConnectionMode(transport.Connected))
	assert.NotNil(t, err)

	// Test RoundTrip
	rsp, err = tc.RoundTrip(context.Background(), []byte("helloworld"),
		transport.WithDialNetwork("udp"),
		transport.WithDialAddress("localhost:9999"),
		transport.WithConnectionMode(transport.Connected))
	assert.Nil(t, err)
	assert.Equal(t, "helloworld", string(rsp))
}

func TestOptions(t *testing.T) {

	opts := &transport.RoundTripOptions{}

	o := transport.WithDialTLS("client.cert", "client.key", "ca.pem", "servername")
	o(opts)
	assert.Equal(t, "client.cert", opts.TLSCertFile)
	assert.Equal(t, "client.key", opts.TLSKeyFile)
	assert.Equal(t, "ca.pem", opts.CACertFile)
	assert.Equal(t, "servername", opts.TLSServerName)
}
