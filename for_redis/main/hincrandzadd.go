package main

import (
	"code_tpl_go/for_redis/redis"
	"code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
)

// RedisHIncrAndZAdd ...
type RedisHIncrAndZAdd struct {
	RedisPool *redigo.Pool
}

func (r *RedisHIncrAndZAdd) Do(ctx context.Context, params ...interface{}) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redigo.Int(redislua.HIncrAndZAdd.Do(conn, params...))
	if err != nil {
		return nil, err
	}
	return result, err
}

func main() {
	flow := RedisHIncrAndZAdd{
		RedisPool: redis.NewPool("127.0.0.1:6379", ""),
	}
	d, e := flow.Do(context.Background(), "incrkey", "flowkey", "billNo123", "432")
	fmt.Printf("[step1] %s %v\n", d, e)
	d, e = flow.Do(context.Background(), "incrkey", "flowkey", "billNo123", "432")
	fmt.Printf("[step2] %s %v", d, e)
}
