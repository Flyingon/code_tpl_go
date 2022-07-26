/*
GOMAXPROCS=1 go test -v -bench=. -count=3 -run=none -benchmem | tee bench.txt
# run `go get -u golang.org/x/perf/cmd/benchstat` if benchstat failed
benchstat bench.txt
*/

package main

import (
	"testing"
)

const it = uint(1 << 26)

func BenchmarkSetWithBoolValueWrite(b *testing.B) {
	set := make(map[uint]bool)

	for i := uint(0); i < it; i++ {
		set[i] = true
	}
}

func BenchmarkSetWithStructValueWrite(b *testing.B) {
	set := make(map[uint]struct{})

	for i := uint(0); i < it; i++ {
		set[i] = struct{}{}
	}
}

func BenchmarkSetWithInterfaceValueWrite(b *testing.B) {
	set := make(map[uint]interface{})

	for i := uint(0); i < it; i++ {
		set[i] = struct{}{}
	}
}
