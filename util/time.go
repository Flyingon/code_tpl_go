package util

import (
	"crypto/rand"
	"math/big"
	"time"
)

func GetRandomMilliSecDuration() time.Duration {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(50))
	t := time.Duration((*bigInt).Uint64())
	if t <= 0 {
		t = time.Duration(10)
	}

	return t * time.Millisecond
}
