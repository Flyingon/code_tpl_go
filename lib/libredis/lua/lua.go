package redislua

import (
	redigo "github.com/gomodule/redigo/redis"
)

var SPopToZSet *redigo.Script
var ZPopMax *redigo.Script
var ZPopByScore *redigo.Script
var ZPopByScoreToZSet *redigo.Script
var ZPopMaxToZSet *redigo.Script

func init() {
	SPopToZSet = redigo.NewScript(2, LSSPopToZSet)
	ZPopByScore = redigo.NewScript(1, LSZPopByScore)
	ZPopMax = redigo.NewScript(1, LsZPopMax)
	ZPopByScoreToZSet = redigo.NewScript(3, LSZPopByScoreToZSet)
	ZPopMaxToZSet = redigo.NewScript(3, LSZPopMaxToZSet)
}
