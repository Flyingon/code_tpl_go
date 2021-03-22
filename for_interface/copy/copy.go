package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	A string
	B map[string]int
}



func main() {
	var d1 interface{}
	d1 = Data{
		B: map[string]int{"12": 1},
	}

	d2 := reflect.New(reflect.TypeOf(d1))
	d1 = d2
	fmt.Printf("d1: %+v, d2: %+v\n", d1, d2)
}
