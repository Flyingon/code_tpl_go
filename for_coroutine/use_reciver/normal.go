package main

import (
	"fmt"
	"strconv"
	"time"
)

type Domain struct {
	A string
	B int
}

type In struct {
	C string
	D int
}

var domain = &Domain{
	A: "1",
	B: 2,
}

func (d *Domain) handler(in, in2 *In) {
	go func() {
		fmt.Println(in.C, in.D, d.A, d.B, in2.C, in2.D)
	}()
}

func b() {
	defer func() {
		fmt.Println("b closed")
	}()
	in2 := &In{
		C: "IN2",
		D: 999,
	}
	for i := 1; i < 1000; i++ {
		in := &In{
			C: strconv.FormatInt(int64(i), 10),
			D: i,
		}
		domain.handler(in, in2)
	}
	return
}

func main() {
	go b()
	<-time.After(10 * time.Second)
}
