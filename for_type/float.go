package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := float64(100000)
	fmt.Println(strconv.FormatFloat(a/100, 'f', 2, 64))
}
