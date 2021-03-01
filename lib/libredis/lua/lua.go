package redislua

import (
	redigo "github.com/gomodule/redigo/redis"
)



var ZPopAndRecord *redigo.Script
var ZPOPMAX *redigo.Script
var ZPopByScore *redigo.Script

func init() {
	ZPopAndRecord = redigo.NewScript(2, LuaScriptKeyPopAndRecord)
	ZPopByScore = redigo.NewScript(1, LuaScriptZPopByScore)
	ZPOPMAX = redigo.NewScript(1, LuaScriptZPopMax)
}

