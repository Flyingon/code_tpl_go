package redislua

import (
	redigo "github.com/gomodule/redigo/redis"
)

var SPopToZSet *redigo.Script
var ZPopMax *redigo.Script
var ZPopByScore *redigo.Script
var ZPopByScoreToZSet *redigo.Script
var ZPopMaxToZSet *redigo.Script
var SeqSetAndIncr *redigo.Script
var SeqSetAndIncrFloat *redigo.Script
var SeqSetAndIncrV2 *redigo.Script
var SeqSetAndIncrFloatV2 *redigo.Script
var HGetAndHDel *redigo.Script
var HCheckAndSet *redigo.Script
var CheckAndZAdd *redigo.Script

func init() {
	SPopToZSet = redigo.NewScript(2, LSSPopToZSet)
	ZPopByScore = redigo.NewScript(1, LSZPopByScore)
	ZPopMax = redigo.NewScript(1, LsZPopMax)
	ZPopByScoreToZSet = redigo.NewScript(3, LSZPopByScoreToZSet)
	ZPopMaxToZSet = redigo.NewScript(3, LSZPopMaxToZSet)
	SeqSetAndIncr = redigo.NewScript(2, LuaScriptSeqSetAndIncr)
	SeqSetAndIncrFloat = redigo.NewScript(2, LuaScriptSeqSetAndIncrFloat)
	SeqSetAndIncrV2 = redigo.NewScript(3, LuaScriptSeqSetAndIncrV2)
	SeqSetAndIncrFloatV2 = redigo.NewScript(3, LuaScriptSeqSetAndIncrFloatV2)
	HGetAndHDel = redigo.NewScript(1, LuaScriptHGetAndHDel)
	HCheckAndSet = redigo.NewScript(2, LuaScriptHCheckAndSet)
	CheckAndZAdd = redigo.NewScript(2, LuaScriptCheckAndZadd)
}
