package main

import redigo "github.com/gomodule/redigo/redis"

// RedisFLow ...
type RedisFLow struct {
	RedisPool *redigo.Pool
}
