package transport_test

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"code_tpl_go/rpc_server/codec"
	"code_tpl_go/rpc_server/transport"
)

type mockChecker struct{}

func (c *mockChecker) Check([]byte) (int, error) {
	return 0, nil
}

type mockMsgReader struct{}

func (r *mockMsgReader) ReadMsg(io.Reader) ([]byte, error) {
	return nil, nil
}

type mockSvrTransport struct{}

func (t *mockSvrTransport) ListenAndServe(ctx context.Context, opts ...transport.ListenServeOption) error {
	return nil
}

type mockClientTransport struct{}

func (c *mockClientTransport) RoundTrip(ctx context.Context, req []byte,
	opts ...transport.RoundTripOption) ([]byte, error) {
	return nil, nil
}

type mockFramerBuilder struct{}

func (f *mockFramerBuilder) New(reader io.Reader) codec.Framer {
	return nil
}

func TestListenAndServe(t *testing.T) {
	var err error
	err = transport.ListenAndServe()
	assert.NotNil(t, err)
}

func TestRoundTrip(t *testing.T) {
	_, err := transport.RoundTrip(context.Background(), nil)
	assert.NotNil(t, err)
}

func TestGetChecker(t *testing.T) {
	transport.RegisterChecker("mock", &mockChecker{})
	c := transport.GetChecker("mock")
	assert.NotNil(t, c)
	assert.Equal(t, &mockChecker{}, c)
}

func TestGetFramerBuilder(t *testing.T) {
	transport.RegisterFramerBuilder("mock", &mockFramerBuilder{})
	f := transport.GetFramerBuilder("mock")
	assert.NotNil(t, f)
	assert.Equal(t, &mockFramerBuilder{}, f)
}

func TestRegisterFramerBuilder_BuilderNil(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	transport.RegisterFramerBuilder("mock", nil)
}

func TestRegisterFramerBuilder_NameNil(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	transport.RegisterFramerBuilder("", &mockFramerBuilder{})
}

func TestGetServerTransport(t *testing.T) {
	transport.RegisterServerTransport("mock", &mockSvrTransport{})
	ts := transport.GetServerTransport("mock")
	assert.NotNil(t, ts)
	assert.Equal(t, &mockSvrTransport{}, ts)
}

func TestGetClientTransport(t *testing.T) {
	transport.RegisterClientTransport("mock", &mockClientTransport{})
	tc := transport.GetClientTransport("mock")
	assert.NotNil(t, tc)
	assert.Equal(t, &mockClientTransport{}, tc)

	// test ClientTransport nil
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	transport.RegisterClientTransport("mock", nil)
}

func TestRegisterClientTransport_NameNil(t *testing.T) {
	// test name nil
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	transport.RegisterClientTransport("", &mockClientTransport{})
}

func TestRegisterChecker_EmptyName(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	transport.RegisterChecker("", &mockChecker{})
}

func TestRegisterNilChecker(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	transport.RegisterChecker("mock", nil)
}

func TestRegisterServerTransport_EmptyName(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	transport.RegisterServerTransport("", &mockSvrTransport{})
}

func TestRegisterNilSvrTransport(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	transport.RegisterServerTransport("mock", nil)
}

func TestRemoteAddrFromContext(t *testing.T) {
	addr := transport.RemoteAddrFromContext(context.Background())
	assert.Nil(t, addr)
}
