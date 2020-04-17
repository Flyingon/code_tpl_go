package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Your application logic here.
	fmt.Println("real GOMAXPROCS", runtime.GOMAXPROCS(-1))
}