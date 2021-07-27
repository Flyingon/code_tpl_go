package main

import (
	"code_tpl_go/for_redis/redis"
	"code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"reflect"
)

type Task struct {
	RedisPool *redigo.Pool
}

func (t *Task) ZPopByScoreToZSet(ctx context.Context, srcKey, dstKey string) (interface{}, error) {
	conn, err := t.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
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
	defer conn.Close()
	result, err := redislua.ZPopMaxToZSet.Do(conn, srcKey, dstKey, "score_map", 100, 20)
	if err != nil {
		return nil, err
	}
	fmt.Printf("res: %+v", result)
	if result != nil {
		fmt.Printf("res: %s, type: %v", result, reflect.TypeOf(result))
	}
	return nil, nil
}

func (t *Task) SeqSetAndIncr(ctx context.Context, seqKey, incrKey, seqField string, seqVal int64, incrField string, incrVal int64) (interface{}, error) {
	conn, err := t.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.SeqSetAndIncr.Do(conn, seqKey, incrKey, seqField, seqVal, incrField, incrVal)
	if err != nil {
		return nil, err
	}
	fmt.Printf("res: %+v\n", result)
	if result != nil {
		fmt.Printf("res: %+v, type: %v\n", result, reflect.TypeOf(result))
	}
	return nil, nil
}

func main() {
	task := Task{
		RedisPool: redis.NewPool("127.0.0.1:6379", ""),
	}
	seq := "1236"
	fmt.Println(task.SeqSetAndIncr(context.Background(), "seq_key", "incr_key", seq,
		100, "incr_field", 2))
	/* 任务调度测试
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
	*/

}
