package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx, nCtx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done", time.Now().String())
			return
		case <-nCtx.Done():
			fmt.Println("nCtx done", time.Now().String())
			return
		default:
			fmt.Printf("nCtx running time: %s\n", time.Now().String())
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	nCtx, nCancel := context.WithCancel(ctx)
	defer nCancel()
	ctxDead, _ := ctx.Deadline()
	nCtxDead, _ := ctx.Deadline()
	fmt.Println("dead_time", ctxDead, nCtxDead)
	go run(ctx, nCtx)
	cancel()
	fmt.Println(ctx.Done())
	fmt.Println(nCtx.Done())
	time.Sleep(10 * time.Second)
}
