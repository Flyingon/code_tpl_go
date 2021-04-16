package util

import (
	"errors"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

)

func init() {
	rand.Seed(time.Now().Unix())
}

type II interface {
	II() *SS
}

type SI struct {
	SS     *SS
	Ch     chan int
	Others map[int]II
}

func (si *SI) II() *SS {
	return si.SS
}

type SS struct {
	MapSI       map[int]II
	Int         int
	Slice       []*SI
	Array       [2]int
	ZeroArray   [0]int
	ZeroPointer *SS
	ZeroSlice   []*SI
	ZeroMap     map[int]II

	unexportedField string
}

func NewSS() *SS {
	return &SS{
		MapSI: make(map[int]II),
		Slice: nil,
	}
}

func (ss *SS) NewSI() *SI {
	id := rand.Int()
	si := SI{
		SS:     ss,
		Ch:     make(chan int),
		Others: ss.MapSI,
	}
	ss.MapSI[id] = &si
	ss.Slice = append(ss.Slice, &si)
	ss.Array[id%2] = id
	return &si
}

func TestDeepCopy(t *testing.T) {
	v, err := DeepCopy(nil)
	require.Nil(t, err)
	require.Nil(t, v)

	ss := NewSS()
	for i := 0; i < 8; i++ {
		ss.NewSI()
	}

	v, err = DeepCopy(ss)
	require.Nil(t, err)
	sss, ok := v.(*SS)
	require.True(t, ok)
	require.True(t, reflect.DeepEqual(sss, v))

	_, err = DeepCopy(&NoCopy{})
	require.NotNil(t, err)
}

type NoCopy struct{}

func (v *NoCopy) Copy() (interface{}, error) {
	return nil, errors.New("do not copy me")
}
