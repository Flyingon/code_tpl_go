package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "Hello word!"

func run(ctx context.Context) {
	defer fmt.Println("ctx return")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx timeout", time.Now().String())
			return
		default:
			// 注意: 这里不能有阻塞操作，io或者sleep，否则还是等待相应之间，这样写只能控制每次循环前判断是否超时
			fmt.Printf("ctx running time: %s\n", time.Now().String())
			time.Sleep(1 * time.Second)
		}
	}
}

func run2(ctx context.Context) {
	defer fmt.Println("nCtx return")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("nCtx timeout", time.Now().String())
			return
		default:
			fmt.Printf("nCtx running time: %s\n", time.Now().String())
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	//<-time.After(3 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//value := context.WithValue(ctx,key,"This is my test")
	nCtx, nCancel := context.WithTimeout(ctx, 10*time.Second)
	defer nCancel()

	ctxDead, _ := ctx.Deadline()
	nCtxDead, _ := ctx.Deadline()
	fmt.Println(ctxDead, nCtxDead)

	go run(ctx)
	go run2(nCtx)
	time.Sleep(10 * time.Second)
}
