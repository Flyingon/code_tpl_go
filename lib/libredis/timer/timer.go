package timer

import (
	redislua "code_tpl_go/lib/libredis/lua"
	"context"
	"errors"
	"fmt"
	"github.com/Flyingon/go-common/util"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/siddontang/go-log/log"
)

type Element struct {
	Key   string `json:"key,omitempty"`
	Score int64  `json:"score,omitempty"`
}

// TaskPush 定时任务写入
func TaskPush(ctx context.Context, redisPool *redigo.Pool, queueName, taskInfo string, executeTs int64) error {
	conn, err := redisPool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	key := fmt.Sprintf("%s", queueName)
	defer func() {
		log.Infof("GetLock[%s]: %v", key, err)
	}()
	_, err = redigo.Int(conn.Do("ZADD", key, int(executeTs), taskInfo))
	//redis返回失败
	if err != nil {
		log.Errorf(err.Error())
		util.ReportMonitor(fmt.Sprintf("定时任务队列(%s)push失败-异常", queueName), 1, 0)
		return err
	}
	return nil
}

// TaskPopMax
func TaskPopMax(ctx context.Context, redisPool *redigo.Pool, queueName string, num int32) ([]*Element, error) {
	conn, err := redisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.ZPopMax.Do(conn, queueName, num)
	//TODO:这里忽略了score顺序
	return result2elements(result, err, "ZPOPMAX")
}

// TaskPopLessThan 定时任务弹出
func TaskPopByScore(ctx context.Context, redisPool *redigo.Pool, queueName string, minScore, maxScore int64) ([]*Element, error) {
	conn, err := redisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.ZPopByScore.Do(conn, queueName, minScore, maxScore, "desc")
	//TODO:这里忽略了score顺序
	return result2elements(result, err, "ZPOPBYSCORE")
}

// TaskPopLessThan 定时任务弹出
func TaskPopLessThan(ctx context.Context, redisPool *redigo.Pool, queueName string, score int64) ([]*Element, error) {
	conn, err := redisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := redislua.ZPopByScore.Do(conn, queueName, "-inf", score, "desc")
	return result2elements(result, err, "ZPOPBYSCORE")
}

func result2elements(result interface{}, err error, opName string) ([]*Element, error) {
	dataMap, err := Float64Map(result, err)
	if err != nil {
		return nil, fmt.Errorf("%v err:%v", opName, err)
	}

	elements := make([]*Element, 0)
	for key, score := range dataMap {
		e := &Element{
			Key:   string(key),
			Score: int64(score),
		}
		elements = append(elements, e)
	}

	return elements, nil
}

func Float64Map(result interface{}, err error) (map[string]float64, error) {
	values, err := redigo.Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("Float64Map expects even number of values result")
	}
	m := make(map[string]float64, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("ScanMap key not a bulk string value")
		}
		value, err := redigo.Float64(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[string(key)] = value
	}
	return m, nil
}
