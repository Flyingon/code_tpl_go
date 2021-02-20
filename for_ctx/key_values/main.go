package main

import (
	"context"
	"fmt"
)

type strKey string
const kkk strKey = "aaa"

func main() {
	ctx := context.Background()
	nCtx := context.WithValue(ctx, kkk, "bbb")
	ccc := nCtx.Value(kkk).(string)
	fmt.Println(ccc)
}
