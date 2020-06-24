package errs_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.code.oa.com/trpc-go/trpc-go/errs"
)

// go test -v -coverprofile=cover.out
// go tool cover -func=cover.out

func TestErrs(t *testing.T) {

	var err *errs.Error
	str := err.Error()
	assert.Contains(t, str, "success")

	e := errs.New(111, "inner fail")
	assert.NotNil(t, e)

	assert.Equal(t, 111, errs.Code(e))
	assert.Equal(t, "inner fail", errs.Msg(e))

	err, ok := e.(*errs.Error)
	assert.Equal(t, true, ok)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrorTypeBusiness, err.Type)

	str = err.Error()
	assert.Contains(t, str, "business")

	e = errs.NewFrameError(111, "inner fail")
	assert.NotNil(t, e)

	assert.Equal(t, 111, errs.Code(e))
	assert.Equal(t, "inner fail", errs.Msg(e))

	err, ok = e.(*errs.Error)
	assert.Equal(t, true, ok)
	assert.NotNil(t, err)
	assert.Equal(t, errs.ErrorTypeFramework, err.Type)

	str = err.Error()
	assert.Contains(t, str, "framework")

	assert.Equal(t, 0, errs.Code(nil))
	assert.Equal(t, "success", errs.Msg(nil))

	assert.Equal(t, 0, errs.Code((*errs.Error)(nil)))
	assert.Equal(t, "success", errs.Msg((*errs.Error)(nil)))

	e = errors.New("unknown error")
	assert.Equal(t, errs.RetUnknown, errs.Code(e))
	assert.Equal(t, "unknown error", errs.Msg(e))

}
