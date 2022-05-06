package main

import (
	"code_tpl_go/for_redis/redis"
	redislua "code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"reflect"
)

// Flow3 ...
type Flow3 struct {
	RedisPool *redigo.Pool
}

func (r *Flow3) SeqSetAndZIncr(ctx context.Context, params ...interface{}) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.SeqSetAndZIncr.Do(conn, params...)
	if err != nil {
		return nil, err
	}
	fmt.Printf("res: %+v\n", result)
	if result != nil {
		fmt.Printf("res: %+v, type: %v\n", result, reflect.TypeOf(result))
		resIntList, e := redigo.Int64s(result, err)
		fmt.Printf("res: %v, err: %v\n", resIntList, e)
	}
	return nil, nil
}

func (r *Flow3) HSetSeqAndVal(ctx context.Context, params ...interface{}) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.HSetSeqAndVal.Do(conn, params...)
	if err != nil {
		return nil, err
	}
	fmt.Printf("res: %+v\n", result)
	if result != nil {
		fmt.Printf("res: %+v, type: %v\n", result, reflect.TypeOf(result))
		resIntList, e := redigo.Int64s(result, err)
		fmt.Printf("res: %v, err: %v\n", resIntList, e)
	}
	return nil, nil
}

func main() {
	flow := Flow3{
		RedisPool: redis.NewPoolWithDial("127.0.0.1:6379", "XXXXXXXXX", redis.DialWithSocks5),
	}
	//fmt.Println(flow.HSetSeqAndVal(context.Background(),
	//	"qr:u:credit:369684877382002816","fp_2493218_10_1642092676_test4", 3, "total", 3, "week", 3))

	res, err := redigo.StringMap(flow.RedisPool.Get().Do("HGETALL", "key_hgetall"))
	fmt.Println(res, len(res), err)
}
