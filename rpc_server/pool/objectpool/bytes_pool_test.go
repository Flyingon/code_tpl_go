package objectpool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code_tpl_go/rpc_server/pool/objectpool"
)

func TestBytesPool_Get(t *testing.T) {
	p := objectpool.NewBytesPool(100)
	assert.NotNil(t, p)

	buf := p.Get()
	assert.NotNil(t, buf)
	p.Put(buf)
}
