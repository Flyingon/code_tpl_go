package redis

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"golang.org/x/net/proxy"
	"net"
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

// NewPoolWithDial args: {maxIdle} {MaxActive} {IdleTimeout}
func NewPoolWithDial(server, password string, dailFunc func(network, addr string) (net.Conn, error), args ...int) *redigo.Pool {
	maxIdle := 3000
	maxActive := 3000
	idleTimeout := 240
	if len(args) >= 3 {
		maxIdle = args[0]
		maxActive = args[1]
		idleTimeout = args[2]
	}
	return &redigo.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive, // 单机最大链接
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp",
				server,
				redigo.DialConnectTimeout(dftTimeOut*time.Millisecond),
				redigo.DialReadTimeout(dftTimeOut*time.Millisecond),
				redigo.DialWriteTimeout(dftTimeOut*time.Millisecond),
				redigo.DialNetDial(dailFunc),
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
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

// DialWithSocks5 通过socks建立tcp链接
func DialWithSocks5(network, addr string) (net.Conn, error) {
	dialerProxy, err := proxy.SOCKS5("tcp", "127.0.0.1:8080", nil, proxy.Direct)
	if err != nil {
		fmt.Printf("socks5 err: %v\n", err)
		return nil, err
	}
	return dialerProxy.Dial(network, addr)
}
