package main

import (
	"code_tpl_go/for_redis/redis"
	redislua "code_tpl_go/lib/libredis/lua"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"math"
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

func (r *Flow2) SeqSetAndIncrFloat(ctx context.Context, seqKey, incrKey, seqField, seqVal string, incrField string, incrVal float64) (interface{}, error) {
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

func (r *Flow2) SeqSetAndIncrV2(ctx context.Context, params ...interface{}) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.SeqSetAndIncrV2.Do(conn, params...)
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

func (r *Flow2) SeqSetAndIncrFloatV2(ctx context.Context, params ...interface{}) (interface{}, error) {
	conn, err := r.RedisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.SeqSetAndIncrFloatV2.Do(conn, params...)
	if err != nil {
		return nil, err
	}
	if result != nil {
		fmt.Printf("res: %s, type: %v\n", result, reflect.TypeOf(result))
		resFloatList, e := redigo.Float64s(result, err)
		fmt.Printf("res: %0.2f, err: %v\n", resFloatList, e)
	}

	return nil, nil
}

func main() {
	flow := Flow2{
		RedisPool: redis.NewPool("127.0.0.1:6379", ""),
	}
	val := float64(1)
	seq := fmt.Sprintf("test-%d", time.Now().Unix())
	seqVal := fmt.Sprintf("%0.2f", val)
	fmt.Println(flow.SeqSetAndIncrFloatV2(context.Background(), "seq_key", "incr_key", "version_key",
		seq, seqVal, "float_field", math.Floor(val*100)/100))

	seq = fmt.Sprintf("test2-%d", time.Now().Unix())
	fmt.Println(flow.SeqSetAndIncrV2(context.Background(), "seq_key", "incr_key", "version_key",
		seq, seqVal, "float_field", int(val)))
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		fmt.Println(flow.SeqSetAndIncrFloatV2(context.Background(), "seq_key", "incr_key", "version_key", seq,
	//			seqVal, "float_field", math.Floor(val*100)/100))
	//	}()
	//}
	//<-time.After(3 * 60 * time.Second)

}
