package objectpool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.code.oa.com/trpc-go/trpc-go/pool/objectpool"
)

func TestBufferPool_Get(t *testing.T) {
	p := objectpool.NewBufferPool()

	buf := p.Get()
	assert.NotNil(t, buf)
	buf.Reset()
	p.Put(buf)
}
