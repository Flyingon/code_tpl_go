package main

import (
	"fmt"
	"time"
)

type Data struct {
	A string
	B int
}

func testDefer () {
	d1 := &Data{
		A: "1",
		B: 1,
	}
	defer func() {
		fmt.Println("A")
		//d1Json, _:= json.Marshal(d1)
		go func() {
			fmt.Printf("d1: %v\n", d1)
		}()
	}()
	d1.A = "11"
	d1.B = 22
	fmt.Println("B")
}

func main() {
	testDefer()
	time.Sleep(5 * time.Second)
}