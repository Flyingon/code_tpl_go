package main

import (
	"fmt"
)

type Elem struct {
	A string
	B int
	C bool
}

// 循环中元素赋值
func rangeSet() {
	elemList := []*Elem{
		{
			"a",
			1,
			false,
		},
		{
			"b",
			2,
			false,
		},
		{
			"c",
			3,
			false,
		},
	}
	for i, e := range elemList {
		if i == 1 {
			e.C = true
		}
	}
	for _, e := range elemList {
		fmt.Printf("%+v\n", e)
	}
}

func main() {
	rangeSet()
}
