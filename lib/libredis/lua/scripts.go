package redislua

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

var LuaScriptZPopByScore = `
local setname = KEYS[1]
local minscore = ARGV[1]
local maxscore = ARGV[2]
local order = ARGV[3]

local redisTable = nil
if order == "desc" then
    redisTable = redis.call('ZREVRANGEBYSCORE', setname, maxscore, minscore,'WITHSCORES')
else
    redisTable = redis.call('ZRANGEBYSCORE', setname, minscore, maxscore,'WITHSCORES')
end

if redisTable == nil
then
  return redisTable
end

for i=1,#redisTable,2 do
  local key = redisTable[i]
  local num = redis.call('ZREM', setname, key)
end
return redisTable
`

var LuaScriptZPopMax = `
local setname = KEYS[1]
local popsize = tonumber(ARGV[1])
if popsize  < 1
then
  return -1
end
local redisTable = redis.call('ZREVRANGE', setname, 0, popsize-1,'WITHSCORES')
if redisTable == nil
then
  return redisTable
end
for i=1,#redisTable,2 do
  local key = redisTable[i]
  local num = redis.call('ZREM', setname, key)
end
return redisTable
`
