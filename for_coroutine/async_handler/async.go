package main

import (
	"fmt"
	"time"
)

type Data struct {
	A string
	B int
	C *string
	D *int
}

func async1(d *Data, f string) {
	time.Sleep(2 * time.Second)
	fmt.Println(d.A, d.B, *d.C, *d.D)
	fmt.Println(f)
}

func main() {

	go func() {
		c := "3"
		d := 4
		d1 := &Data{
			A: "1",
			B: 2,
			C: &c,
			D: &d,
		}
		//go async1(d1, c)
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println(d1.A, d1.B, *d1.C, *d1.D)
		}()

	}()
	<-time.After(10 * time.Second)
}
