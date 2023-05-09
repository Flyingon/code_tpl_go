package main

import "fmt"

type Struct1 struct {
	T *Struct2
}

type Struct2 struct {
	A string
	B int
}

func main() {
	st := &Struct1{T: &Struct2{
		A: "a",
		B: 1,
	}}
	stt := st.T
	st2 := &Struct1{T: &Struct2{
		A: "b",
		B: 2,
	}}
	st = st2
	fmt.Println(stt, st2.T)
}
