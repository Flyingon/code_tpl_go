package connpool

import (
	"io"
	"testing"

	"git.code.oa.com/trpc-go/trpc-go/codec"
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
