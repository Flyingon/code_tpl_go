package main

import (
	"code_tpl_go/for_redis/redis"
	"code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

func (r *RedisFLow) HGetAndHDel(ctx context.Context, key, field string) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.HGetAndHDel.Do(conn, key, field)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("res: %+v\n", result)
	if result != nil {
		d, e := redigo.Int64(result, err)
		//fmt.Printf("res: %+v, e: %v\n", d, e)
		return d, e
	}
	return nil, nil
}

func main() {
	flow := RedisFLow{
		RedisPool: redis.NewPool("127.0.0.1:6379", ""),
	}
	for i := 0; i < 1000; i++ {
		go func() {
			for true {
				d, e := flow.HGetAndHDel(context.Background(), "a", "a")
				if e != nil {
					fmt.Printf("[ERROR] %v\n", e)
				}
				if d != nil {
					fmt.Printf("d: %v\n", d)
				}
			}
		}()
	}
	<-time.After(3 * 60 * time.Second)
}
