package main

import (
	"code_tpl_go/for_redis/redis"
	"code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
)

// RedisCheckAndZAdd ...
type RedisCheckAndZAdd struct {
	RedisPool *redigo.Pool
}

func (r *RedisCheckAndZAdd) CheckAndZAdd(ctx context.Context, params ...interface{}) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.CheckAndZAdd.Do(conn, params...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func main() {
	flow := RedisCheckAndZAdd{
		RedisPool: redis.NewPool("127.0.0.1:6379", ""),
	}
	d, e := flow.CheckAndZAdd(context.Background(), "uidKey", "rankKey", "334", "member", "432")
	fmt.Printf("[step1] %s %v\n", d, e)
	d, e = flow.CheckAndZAdd(context.Background(), "uidKey", "rankKey", "334", "member", "432")
	fmt.Printf("[step2] %s %v", d, e)
}
