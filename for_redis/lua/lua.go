package main

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"math"
)

var LuaScriptKeyPopAndRecord = `
redis.replicate_commands()
local popQueue = KEYS[1]
local recordQueue = KEYS[2]
local ts=redis.call('TIME')[1]
 
local elems = redis.call("SPOP", popQueue, 1)
if #elems == 1 then
  redis.call('ZADD', recordQueue, ts, elems[1])
  return elems[1]
end
`

var ZPopAndRecord *redigo.Script

func init() {
	ZPopAndRecord = redigo.NewScript(2, LuaScriptKeyPopAndRecord)
}

func main() {
	fmt.Println(math.Trunc( 3 / 9 ))
}
