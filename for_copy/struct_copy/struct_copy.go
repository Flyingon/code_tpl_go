package main

import "fmt"

type TmpStruct struct {
	A string
	B string
	C int
}

func main() {
	s1 := TmpStruct{
		A: "s1",
		B: "s1",
		C: 1,
	}
	s2 := s1
	s2.A = "s2"

	fmt.Printf("s1: %p, s1: %+v\n", &s1, s1)
	fmt.Printf("s2: %p, s2: %+v\n", &s2, s2)
}