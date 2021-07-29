package main

import (
	"code_tpl_go/for_redis/redis"
	redislua "code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"reflect"
	"time"
)

// Flow2 ...
type Flow2 struct {
	RedisPool *redigo.Pool
}

func (r *Flow2) SeqSetAndIncr(ctx context.Context, seqKey, incrKey, seqField string, seqVal int64, incrField string, incrVal int64) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
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

func (r *Flow2) SeqSetAndIncrFloat(ctx context.Context, seqKey, incrKey, seqField string, seqVal int64, incrField string, incrVal float64) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.SeqSetAndIncrFloat.Do(conn, seqKey, incrKey, seqField, seqVal, incrField, incrVal)
	if err != nil {
		return nil, err
	}
	if result != nil {
		//fmt.Printf("res: %s, type: %v\n", result, reflect.TypeOf(result))
		resFloat, e := redigo.Float64(result, err)
		fmt.Printf("res: %0.2f, err: %v\n", resFloat, e)
	}

	return nil, nil
}

func main() {
	flow := Flow2{
		RedisPool: redis.NewPool("127.0.0.1:6379", ""),
	}
	seq := fmt.Sprintf("test-%d", time.Now().Unix())
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(flow.SeqSetAndIncrFloat(context.Background(), "seq_key", "incr_key", seq,
				100, "float_field", 2.43))
		}()
	}
	<-time.After(3 * 60 * time.Second)

}
