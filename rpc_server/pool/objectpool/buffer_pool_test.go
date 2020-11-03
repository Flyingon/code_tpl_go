package objectpool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code_tpl_go/rpc_server/pool/objectpool"
)

func TestBufferPool_Get(t *testing.T) {
	p := objectpool.NewBufferPool()

	buf := p.Get()
	assert.NotNil(t, buf)
	buf.Reset()
	p.Put(buf)
}
