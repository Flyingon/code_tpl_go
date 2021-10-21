package main

import (
	"errors"
	"fmt"
	"time"
)

type Data struct {
	A string
	B int
}

func testDefer() {
	d1 := &Data{
		A: "1",
		B: 1,
	}
	var c string
	var err error
	defer func() {
		fmt.Println("A")
		//d1Json, _:= json.Marshal(d1)
		go func() {
			fmt.Printf("d1: %v, c: %s, err: %v\n", d1, c, err)
		}()
	}()
	d1.A = "11"
	d1.B = 22
	c = "c"
	fmt.Println("B")
	err = errors.New("ERROR")
}

func main() {
	testDefer()
	time.Sleep(5 * time.Second)
}
