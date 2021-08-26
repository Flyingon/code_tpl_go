package main

import (
	"code_tpl_go/for_redis/redis"
	"context"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// Pool redis 连接池
type Pool struct {
	redisPool *redigo.Pool
}

// SingleCmd 一次命令
type SingleCmd struct {
	Cmd  string
	Args []interface{}
	Res  interface{}
}

// PipeLineOri 基本管道操作
func (p *Pool) PipeLineOri(ctx context.Context, cmdList []*SingleCmd) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	attrs := make(map[string]string, 2)
	attrs["commandName"] = "PipeLine"
	paramsStr, _ := jsoniter.MarshalToString(cmdList)
	attrs["params"] = string(paramsStr)
	var err error
	for _, single := range cmdList {
		err = conn.Send(single.Cmd, single.Args...)
		if err != nil {
			return err
		}
	}
	err = conn.Flush()
	if err != nil {
		return err
	}
	for index := range cmdList {
		res, e := conn.Receive()
		if e != nil {
			return e
		}
		cmdList[index].Res = res
	}
	return nil
}

func isCmdResListValid(cmdList []*SingleCmd) bool {
	if cmdList == nil || len(cmdList) == 0 {
		return false
	}
	for _, cmd := range cmdList {
		if cmd.Res != nil {
			return true
		}
	}
	return false
}

func (p Pool) PipeLine(ctx context.Context, params []map[string]interface{}) ([]interface{}, error) {
	attrs := make(map[string]string, 2)
	attrs["commandName"] = "PipeLine"
	paramsStr, _ := jsoniter.MarshalToString(params)
	attrs["params"] = string(paramsStr)
	conn := p.redisPool.Get()
	defer conn.Close()

	var err error
	for _, param := range params {
		fStr := fmt.Sprint(param["func"])

		if argsSLice, ok := param["args"].([]interface{}); ok {
			err = conn.Send(fStr, argsSLice...)
			if err != nil {
				return nil, errors.Wrapf(err, "fStr=%s, args=%v", fStr, params)
			}
		} else {
			return nil, errors.New("1003")
		}

	}

	err = conn.Flush()
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	result := make([]interface{}, 0)
	for _, paramValue := range params {
		fStr := strings.ToLower(fmt.Sprint(paramValue["func"]))
		args := paramValue["args"].([]interface{})

		if inArray(fStr, []interface{}{"zrange", "zrevrange", "zrangebyscore", "zrevrangebyscore"}) {
			rs, err := redigo.Strings(conn.Receive())
			if err != nil {
				return nil, errors.Wrapf(err, "fStr=%s, args=%v", fStr, args)
			}
			if len(args) > 3 && fmt.Sprint(args[3]) == "withscores" {
				result = append(result, transSliceToKv(rs))
			} else {
				result = append(result, rs)
			}
		} else if fStr == "hgetall" {
			rs, err := redigo.StringMap(conn.Receive())
			if err != nil {
				return nil, errors.Wrapf(err, "fStr=%s, args=%v", fStr, args)
			}
			result = append(result, rs)
		} else if fStr == "hmget" {
			rs, err := redigo.Strings(conn.Receive())
			if err != nil {
				return nil, errors.Wrapf(err, "fStr=%s, args=%v", fStr, args)
			}

			rsp := make(map[string]string, 0)
			if len(args) > 1 {
				keyList := args[1:]
				for index, value := range keyList {
					key, err := redigo.String(value, nil)
					if err != nil {
						return nil, errors.Wrapf(err, "fStr=%s, args=%v", fStr, args)
					}
					rsp[key] = rs[index]
				}
				result = append(result, rsp)
			} else {
				result = append(result, rsp)
			}
		} else if inArray(fStr, GetStringReturnCommand()) {
			rs, err := redigo.String(conn.Receive())
			if err != nil && err != redigo.ErrNil {
				return nil, errors.Wrapf(err, "fStr=%s, args=%v", fStr, args)
			}
			result = append(result, rs)
		} else if inArray(fStr, GetIntReturnCommand()) {
			rs, err := redigo.Int(conn.Receive())
			if err != nil && err != redigo.ErrNil {
				return nil, errors.Wrapf(err, "fStr=%s, args=%v", fStr, args)
			}
			result = append(result, rs)
		} else {
			return nil, errors.New("1002")
		}
	}

	return result, err
}

func GetStringReturnCommand() []interface{} {
	return []interface{}{"set", "setex", "psetex", "get", "getset", "mset", "hget", "rpop", "lpop", "spop", "hmset", "srandmember", "rename"}
}

func GetIntReturnCommand() []interface{} {
	return []interface{}{"del", "setnx", "setex", "incr", "incrby", "decr", "decyby", "hsetnx", "hdel", "hlen",
		"hincrby", "hexists", "lpush", "rpush", "llen", "sadd", "sismembers", "srem", "scard", "zadd", "zadd",
		"zscore", "zincrby", "zcard", "zcount", "zrank", "zrevrank", "zrem", "zremrangebyrank", "zremrangebyscore",
		"setbit", "getbit", "expire", "expireat", "ttl", "hset"}
}

func inArray(needle interface{}, arr []interface{}) bool {
	for _, value := range arr {
		if value == needle {
			return true
		}
	}
	return false
}

func transSliceToKv(strList []string) []map[string]string {
	rspArr := make([]map[string]string, 0)
	for i := 0; i < len(strList); i += 2 {
		key := strList[i]
		value := strList[i+1]
		rspArr = append(rspArr, map[string]string{
			"member": key,
			"value":  value,
		})
	}
	return rspArr
}

func isResListValid(resLis []interface{}) bool {
	if resLis == nil || len(resLis) == 0 {
		return false
	}
	for _, res := range resLis {
		if res != nil {
			return true
		}
	}
	return false
}

func testPipeLineOri() {
	p := Pool{
		redisPool: redis.NewPool("ssd1.userarchive.xytlsq.db:50001", "WOvoNlifCbCTZgpb"),
	}
	cmdList := []*SingleCmd{
		{
			Cmd:  "HGET",
			Args: []interface{}{"1111", "abc"},
			Res:  nil,
		},
		{
			Cmd:  "HMGET",
			Args: []interface{}{"seq_key", "test-1627528420", "test-1627528397", "cccc"},
			Res:  nil,
		},
	}
	err := p.PipeLineOri(context.Background(), cmdList)
	fmt.Println("err: ", err, isCmdResListValid(cmdList))
	for _, cmd := range cmdList {
		d, _ := redigo.Strings(cmd.Res, nil)
		fmt.Println("res: ", cmd.Res, d)
	}
}

func testPipeLine() {
	p := Pool{
		redisPool: redis.NewPool("ssd1.userarchive.xytlsq.db:50001", "WOvoNlifCbCTZgpb"),
	}
	params := []map[string]interface{}{
		{
			"func": "HGET",
			"args": []interface{}{"1111", "abc"},
		},
		{
			"func": "HGET",
			"args": []interface{}{"seq_key1", "test-1627528420"},
		},
	}
	resList, err := p.PipeLine(context.Background(), params)
	fmt.Println("err: ", err, isResListValid(resList))
	for _, res := range resList {
		fmt.Println("res: ", res, reflect.TypeOf(res))
	}
}

func main() {
	testPipeLineOri()
	testPipeLine()
}
