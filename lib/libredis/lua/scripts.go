package redislua

var LSSPopToZSet = `
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

var LSZPopByScore = `
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

var LsZPopMax = `
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

var LSZPopByScoreToZSet = `
redis.replicate_commands()
local srcKey = KEYS[1]
local dstKey = KEYS[2]
local scoreKey = KEYS[3]
local minScore = ARGV[1]
local maxScore = ARGV[2]
local order = ARGV[3]
local newScore = tonumber(ARGV[4])

local zSetData = nil
if order == "desc" then
    zSetData = redis.call('ZREVRANGEBYSCORE', srcKey, maxScore, minScore, 'WITHSCORES')
else
    zSetData = redis.call('ZRANGEBYSCORE', srcKey, minScore, maxScore, 'WITHSCORES')
end

if zSetData == nil or #zSetData == 0
then
 return nil
end

for i=1,#zSetData,2 do
 local key = zSetData[i]
--解析score
  local taskType = ""
  for token in string.gmatch(key, "[^|]+") do
    taskType = token
    break
  end
  local addScore = redis.call('HGET', scoreKey, taskType)
--计算新score
  local score = newScore
  if addScore ~= nil and addScore ~= false
  then
    score = newScore + addScore
  end
  local num = redis.call('ZREM', srcKey, key)
  redis.call('ZADD', dstKey, score, key)
end
return zSetData
`

var LSZPopMaxToZSet = `
redis.replicate_commands()
local srcKey = KEYS[1]
local dstKey = KEYS[2]
local scoreKey = KEYS[3]
local popSize = tonumber(ARGV[1])
local newScore = tonumber(ARGV[2])
if popSize < 1
then
  return -1
end
local zSetData = redis.call('ZREVRANGE', srcKey, 0, popSize-1,'WITHSCORES')

if zSetData == nil or #zSetData == 0
then
 return nil
end

for i=1,#zSetData,2 do
  local key = zSetData[i]
--解析score
  local taskType = ""
  for token in string.gmatch(key, "[^|]+") do
    taskType = token
    break
  end
  local addScore = redis.call('HGET', scoreKey, taskType)
--计算新score
  local score = newScore
  if addScore ~= nil and addScore ~= false
  then
    score = newScore + addScore
  end
  redis.call('ZREM', srcKey, key)
  redis.call('ZADD', dstKey, score, key)
end
return zSetData
`

// LuaScriptSeqSetAndIncr 设置流水号并hincr数据
var LuaScriptSeqSetAndIncr = `
local seqKey = KEYS[1]
local incrKey = KEYS[2]
local seqField = ARGV[1]
local seqVal = ARGV[2]
local incrField = ARGV[3]
local incrVal = tonumber(ARGV[4])

local rspSeq = redis.call('HSET', seqKey, seqField, seqVal)

if rspSeq ~= 1
then
  return -1
end

local rspIncr = redis.call('HINCRBY', incrKey, incrField, incrVal)
return rspIncr
`

// LuaScriptSeqSetAndIncrFloat 设置流水号并hincr float数据
var LuaScriptSeqSetAndIncrFloat = `
local seqKey = KEYS[1]
local incrKey = KEYS[2]
local seqField = ARGV[1]
local seqVal = ARGV[2]
local incrField = ARGV[3]
local incrVal = tonumber(ARGV[4])

local rspSeq = redis.call('HSET', seqKey, seqField, seqVal)

if rspSeq ~= 1
then
  return "-1"
end

local rspIncr = redis.call('HINCRBYFLOAT', incrKey, incrField, incrVal)
return rspIncr
`

// LuaScriptHGetAndHDel hget获取数据并hdel删除
var LuaScriptHGetAndHDel = `
local rdKey = KEYS[1]
local field = ARGV[1]

local rspData = redis.call('HGET', rdKey, field)

redis.call('HDEL', rdKey, field)
return rspData
`

// LuaScriptHCheckAndSet hset field检查并设置值
var LuaScriptHCheckAndSet = `
local seqKey = KEYS[1]
local userKey = KEYS[2]
local seqField = ARGV[1]
local seqVal = ARGV[2]
local acctField = ARGV[3]
local cmpVal = tonumber(ARGV[4])
local newVal = tonumber(ARGV[5])

local balance = redis.call('HGET', userKey, acctField)
if balance ~= cmpVal
then 
	return -2
end

local rspSeq = redis.call('HSET', seqKey, seqField, seqVal)
if rspSeq ~= 1
then
  return -1
end

local rspSet = redis.call('HSET', userKey, acctField, newVal)
return rspSet
`
