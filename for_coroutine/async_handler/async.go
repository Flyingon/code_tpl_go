package main

import (
	"context"
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
	time.Sleep(1 * time.Second)
	d.A = "LLL"
	time.Sleep(1 * time.Second)
	fmt.Println(d.A, d.B, *d.C, *d.D)
	fmt.Println(f)
}

func asyncStructCheck() {
	go func() {
		c := "3"
		d := 4
		d1 := &Data{
			A: "1",
			B: 2,
			C: &c,
			D: &d,
		}
		go async1(d1, c)
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println(d1.A, d1.B, *d1.C, *d1.D)
		}()
		fmt.Println("first coroutine end, d1.A: ", d1.A)

	}()
	<-time.After(10 * time.Second)
}

func asyncCtxCheck() {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		go func(ctx context.Context) {

			fmt.Println("sub cor begin, ctx: ", ctx.Err())
			time.Sleep(2 * time.Second)
			fmt.Println("sub cor 2 second ago, ctx: ", ctx.Err())
		}(ctx)
		fmt.Println("main coroutine end")

	}()
	<-time.After(10 * time.Second)
}

func main() {
	//asyncStructCheck()
	asyncCtxCheck()
}
