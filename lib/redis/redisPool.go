package redis

import (
	"fmt"
	"code_tpl_go/rpc_server/log"
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"os"
	"time"
)

const dftTimeOut = 3000 // 毫秒

func NewPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3000,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				server,
				redis.DialConnectTimeout(dftTimeOut*time.Millisecond),
				redis.DialReadTimeout(dftTimeOut*time.Millisecond),
				redis.DialWriteTimeout(dftTimeOut*time.Millisecond),
			)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitScripts() {
}

func NewScripts(fileName string) *redis.Script {
	script, err := CreateScriptFromFile(fileName, 1)
	if err != nil {
		log.Error("load scripts from %v fail:%v", fileName, err)
		panic(err)
	}
	return script
}

func CreateScriptFromFile(filename string, keyCount int) (*redis.Script, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file err:%v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read file err:%v", err)
	}
	luaStr := string(b)
	return redis.NewScript(keyCount, luaStr), nil
}
