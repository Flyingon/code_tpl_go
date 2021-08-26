package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RoundFloat64(t *testing.T) {
	d := RoundFloat64(0.1566, 2)
	assert.Equal(t, d, 0.16)
}

func Test_RoundDown(t *testing.T) {
	d1 := RoundDown(0.1566, 2)
	assert.Equal(t, d1, 0.15)
	d2 := RoundDown(1.001111, 2)
	assert.Equal(t, d2, 1.0)
	d3 := RoundDown(0.99999, 2)
	assert.Equal(t, d3, 0.99)
}
