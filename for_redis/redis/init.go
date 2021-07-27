package redis

import (
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

const dftTimeOut = 3000 // 毫秒

func NewPool(server, password string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     3000,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp",
				server,
				redigo.DialConnectTimeout(dftTimeOut*time.Millisecond),
				redigo.DialReadTimeout(dftTimeOut*time.Millisecond),
				redigo.DialWriteTimeout(dftTimeOut*time.Millisecond),
			)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
