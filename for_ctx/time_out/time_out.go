package main

import (
	"context"
	"fmt"
	"time"
)
var key string = "Hello word!"

func run(ctx context.Context){
	for {
		select {
		case <-ctx.Done():
			fmt.Println("run timeout", time.Now().String())
			return
		default:
			// 注意: 这里不能有阻塞操作，io或者sleep，否则还是等待相应之间，这样写只能控制每次循环前判断是否超时
			fmt.Printf("run running time: %s\n", time.Now().String())
			time.Sleep(1*time.Second)
		}
	}
}

func run2(ctx context.Context) () {
	defer fmt.Println("run2 return")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("run2 timeout", time.Now().String())
			return
		default:
			fmt.Printf("run2 running time: %s\n", time.Now().String())
			time.Sleep(1*time.Second)
		}
	}
}

func main() {
	//<-time.After(3 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(),3 * time.Second)
	defer cancel()
	//value := context.WithValue(ctx,key,"This is my test")
	nCtx, _ := context.WithTimeout(ctx, 5*time.Second)
	go run(ctx)
	go run2(nCtx)
	time.Sleep(10 * time.Second)
}