package main

import (
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
	defer func() {
		fmt.Println("A")
		//d1Json, _:= json.Marshal(d1)
		go func() {
			fmt.Printf("d1: %v, c: %s\n", d1, c)
		}()
	}()
	d1.A = "11"
	d1.B = 22
	c = "c"
	fmt.Println("B")
}

func main() {
	testDefer()
	time.Sleep(5 * time.Second)
}
