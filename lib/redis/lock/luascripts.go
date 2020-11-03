package lock

// LuaScriptGetCmpDel 查询比对并删除
var LuaScriptGetCmpDel = `
local keyName = KEYS[1]
local keyVal = ARGV[1]

local rspGet = redis.call('GET', keyName)

if rspGet == false
then
 return -1
end

if rspGet == keyVal
then
 local rspDel = redis.call('DEL', keyName)
 return rspDel
end
return -2
`