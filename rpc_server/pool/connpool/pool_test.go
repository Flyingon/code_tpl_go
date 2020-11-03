package connpool

import (
	"io"
	"testing"

	"code_tpl_go/rpc_server/codec"
	"github.com/stretchr/testify/assert"
)

func TestWithGetOptions(t *testing.T) {
	opts := &GetOptions{}

	fb := &emptyFramerBuilder{}
	WithFramerBuilder(fb)(opts)

	assert.Equal(t, opts.FramerBuilder, fb)
}

type emptyFramerBuilder struct{}

func (*emptyFramerBuilder) New(io.Reader) codec.Framer {
	return &emptyFramer{}
}

type emptyFramer struct{}

func (*emptyFramer) ReadFrame() ([]byte, error) {
	return nil, nil
}
