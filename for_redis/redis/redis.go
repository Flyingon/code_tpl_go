package redis

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// redis错误信息，外层可通过这里判断具体错误
var (
	ErrorParamInvalid   = fmt.Errorf("param invalid")
	ErrorAddressingFail = fmt.Errorf("addressing fail")
	ErrorDataEmpty      = fmt.Errorf("data empty")
	ErrorDataInvalid    = fmt.Errorf("data invalid")
	ErrorMarshalFail    = fmt.Errorf("pb marshal fail")
	ErrorUnmarshalFail  = fmt.Errorf("pb unmarshal fail")
	ErrorGetConnFail    = fmt.Errorf("get conn fail")
	ErrorSetCasFail     = fmt.Errorf("set cas fail")
	ErrorGetConflict    = fmt.Errorf("get conflict")
)

var redisPool = make(map[string]*redigo.Pool, 0)
var redisPoolLock sync.RWMutex

type RedisConf struct {
	Address   string
	Timeout   time.Duration
	Password  string
	MaxIdle   int
	MaxActive int
}

// Redis 后端请求结构体
type Redis struct {
	address   string // ip://ip:port cmlb://appid
	timeout   time.Duration
	password  string
	maxIdle   int
	maxActive int

	command string
	key     string
	cost    time.Duration
	err     error
}

// New 新建一个redis后端请求结构体
// address="10.100.67.132:9736"
// timeout="800ms"
// password="!QAZ@WSX3e"
// maxIdle = 100
// maxActive = 0 // no limit
func New(conf RedisConf) *Redis {
	o := &Redis{
		address:   conf.Address,
		timeout:   conf.Timeout,
		password:  conf.Password,
		maxIdle:   conf.MaxIdle,
		maxActive: conf.MaxActive,
	}

	return o
}

// Do 执行redis命令
func (c *Redis) Do(ctx context.Context, dbNumber uint32, commandName string, args ...interface{}) (interface{}, error) {
	c.err = nil
	c.command = commandName
	if c.address == "" {
		c.err = errors.New("redis address empty")
		return nil, c.err
	}
	if dbNumber > 15 {
		c.err = errors.New("redis db index is out of range")
		return nil, c.err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if len(args) > 0 {
		if key, ok := args[0].(string); ok {
			c.key = key
		}
	}
	begin := time.Now()

	conn := c.GetConn(ctx)
	if conn == nil {
		return nil, c.err
	}
	defer conn.Close()
	if dbNumber != 0 {
		_, e := conn.Do("SELECT", dbNumber)
		if e != nil {
			c.err = e
			return nil, c.err
		}
	}
	reply, err := conn.Do(commandName, args...)
	if err != nil {
		c.err = err
	}

	c.cost = time.Now().Sub(begin)

	return reply, c.err
}

func (c *Redis) Strings(reply interface{}, err error) ([]string, error) {
	rlt, err := redigo.Strings(reply, err)
	if err != nil && err != redigo.ErrNil {
		return nil, err
	}
	return rlt, err
}

func (c *Redis) Ints(reply interface{}, err error) ([]int, error) {
	rlt, err := redigo.Ints(reply, err)
	if err != nil && err != redigo.ErrNil {
		return nil, err
	}
	return rlt, err
}

func (this *Redis) StringMap(reply interface{}, err error) (map[string]string, error) {
	rlt, err := redigo.StringMap(reply, err)
	if err != nil && err != redigo.ErrNil {
		return nil, err
	}
	return rlt, err
}

func (this *Redis) Bool(reply interface{}, err error) (bool, error) {
	rlt, err := redigo.Bool(reply, err)
	if err != nil && err != redigo.ErrNil {
		return false, err
	}
	return rlt, err
}

func (this *Redis) Bytes(reply interface{}, err error) ([]byte, error) {
	rlt, err := redigo.Bytes(reply, err)
	if err != nil && err != redigo.ErrNil {
		return nil, err
	}
	return rlt, err
}

func (this *Redis) String(reply interface{}, err error) (string, error) {
	rlt, err := redigo.String(reply, err)
	if err != nil && err != redigo.ErrNil {
		return "", nil
	}
	return rlt, err
}

func (this *Redis) Float64(reply interface{}, err error) (float64, error) {
	rlt, err := redigo.Float64(reply, err)
	if err != nil && err != redigo.ErrNil {
		return 0, err
	}
	return rlt, err
}

func (this *Redis) Uint64(reply interface{}, err error) (uint64, error) {
	rlt, err := redigo.Uint64(reply, err)
	if err != nil && err != redigo.ErrNil {
		return 0, err
	}
	return rlt, err
}

func (this *Redis) Int64(reply interface{}, err error) (int64, error) {
	rlt, err := redigo.Int64(reply, err)
	if err != nil && err != redigo.ErrNil {
		return 0, err
	}
	return rlt, err
}

func (this *Redis) Int(reply interface{}, err error) (int, error) {
	rlt, err := redigo.Int(reply, err)
	if err != nil && err != redigo.ErrNil {
		return 0, err
	}
	return rlt, err
}

// GetConn 获取redis链接，用于pipeline, conn.Send() ...
func (c *Redis) GetConn(ctx context.Context) redigo.Conn {
	key := fmt.Sprintf("%s:%s", c.address, c.password)

	var ok bool
	var pool *redigo.Pool
	redisPoolLock.RLock()
	pool, ok = redisPool[key]
	redisPoolLock.RUnlock()

	if ok {
		return pool.Get()
	}

	redisPoolLock.Lock()
	defer redisPoolLock.Unlock()

	pool, ok = redisPool[key]
	if ok {
		return pool.Get()
	}

	password := c.password
	timeout := c.timeout
	addr := c.address
	pool = &redigo.Pool{
		MaxIdle:     c.maxIdle,
		MaxActive:   c.maxActive,
		IdleTimeout: 3 * time.Minute,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp",
				addr,
				redigo.DialConnectTimeout(timeout),
				redigo.DialReadTimeout(timeout),
				redigo.DialWriteTimeout(timeout),
			)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	redisPool[key] = pool

	return pool.Get()
}

// String output string
func (c *Redis) DebugString() string {
	if c.err != nil {
		return fmt.Sprintf("redis[%s.%s], addr[%s], cost[%s], error[%+v]", c.command, c.key, c.address, c.cost, c.err)
	}
	return fmt.Sprintf("redis[%s.%s], addr[%s], cost[%s]", c.command, c.key, c.address, c.cost)
}
