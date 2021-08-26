package main

import (
	"context"
	"fmt"
	"time"
)

type strKey string

const kkk strKey = "aaa"

// Deadline不能重置，只能新建ctx
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	deadTime, _ := ctx.Deadline()
	fmt.Printf("deadTime: %+v\n", deadTime)
	nCtx := context.WithValue(ctx, kkk, "bbb")
	deadTime, _ = nCtx.Deadline()
	fmt.Printf("deadTime: %+v\n", deadTime)
	nnCtx, cancel := context.WithTimeout(nCtx, 5*time.Second)
	cancel()
	deadTime, _ = nnCtx.Deadline()
	fmt.Printf("deadTime: %+v\n", deadTime)
	ccc := nnCtx.Value(kkk).(string)
	fmt.Println(ccc)
}
