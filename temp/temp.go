package main

import (
	"fmt"
	"time"
)

func main() {
	ddd := make(map[uint64]int)
	ddd[123] = 123
	fmt.Println(ddd[333])

	ts := 24 * 7 * time.Hour
	fmt.Println(ts.String())
}