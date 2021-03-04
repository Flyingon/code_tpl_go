package main

import (
	"code_tpl_go/for_redis/redis"
	"code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"reflect"
	"time"
)

type Task struct {
	RedisPool *redigo.Pool
}

func (t *Task) ZPopByScoreToZSet(ctx context.Context, srcKey, dstKey string) (interface{}, error) {
	conn, err := t.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	result, err := redislua.ZPopByScoreToZSet.Do(conn, srcKey, dstKey, "-INF", "+INF", "desc", 10)
	if err != nil {
		return nil, err
	}
	if result != nil {
		fmt.Printf("res: %+v, type: %v", result, reflect.TypeOf(result))
	}
	return nil, nil
}

func (t *Task) ZPopMaxToZSet(ctx context.Context, srcKey, dstKey string) (interface{}, error) {
	conn, err := t.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	result, err := redislua.ZPopMaxToZSet.Do(conn, srcKey, dstKey, 100, 20)
	if err != nil {
		return nil, err
	}
	if result != nil {
		fmt.Printf("res: %+v, type: %v", result, reflect.TypeOf(result))
	}
	return nil, nil
}

func main() {
	task := Task{
		RedisPool: redis.NewPool("10.10.10.10:6380", "XXX"),
	}
	fmt.Println("begin loop")
	for true {
		var err error
		//_, err = task.ZPopByScoreToZSet(context.Background(), "test:task", "test:task2")
		//if err != nil {
		//	fmt.Println(err)
		//}
		_, err = task.ZPopMaxToZSet(context.Background(), "test:task2", "test:task")
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
