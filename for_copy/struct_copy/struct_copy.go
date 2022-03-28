package main

import "fmt"

type InnerStruct struct {
	D string
	E int
}

type TmpStruct struct {
	A string
	C int
	InnerStruct
}

func main() {
	s1 := &TmpStruct{
		A: "s1",
		C: 1,
	}
	s1.E = 3
	fmt.Printf("origin s1: %p, inner: %p, s1: %+v\n", &(*s1), &s1.InnerStruct, s1)
	s2 := &(*s1) // copy 不带innerStruct
	s2.A = "s2"
	s2.E = 5

	s3 := &TmpStruct{ // copy 带innerStruct
		A:           "s1",
		C:           1,
		InnerStruct: s1.InnerStruct,
	}
	s3.E = 6

	fmt.Printf("s1: %p, inner: %p, s1: %+v\n", &(*s1), &s1.InnerStruct, s1)
	fmt.Printf("s2: %p, inner: %p, s2: %+v\n", &(*s2), &s2.InnerStruct, s2)
	fmt.Printf("s3: %p, inner: %p, s3: %+v\n", &(*s3), &s3.InnerStruct, s3)
}
