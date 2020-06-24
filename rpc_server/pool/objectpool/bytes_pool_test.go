package objectpool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.code.oa.com/trpc-go/trpc-go/pool/objectpool"
)

func TestBytesPool_Get(t *testing.T) {
	p := objectpool.NewBytesPool(100)
	assert.NotNil(t, p)

	buf := p.Get()
	assert.NotNil(t, buf)
	p.Put(buf)
}
