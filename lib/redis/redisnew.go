package redis

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"sync"
)

type SyncMap struct {
	sync.Map
}

var redisCltMap sync.Map

type Pool struct {
	redisPool *redis.Pool
}

func SetRedisClt(key string, clt *redis.Pool) {
	redisCltMap.Store(key, clt)
}

func GetRedisClt(key string) *Pool {
	cltPtr, exist := redisCltMap.Load(key)
	if !exist {
		return nil
	}
	clt, ok := cltPtr.(*redis.Pool)
	if !ok {
		return nil
	}
	return &Pool{clt}
}

// TRedisConf redis配置
/*
# redis配置
redispool:
  host: 9.140.160.121
  port: 41002
  password: testpass
*/
type TRedisConf struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
}

// Init redis clt初始化
func Init(redisConfMap map[string]TRedisConf) {
	confList := redisConfMap
	log.Infof("redis init, config list: %+v", confList)
	for k, v := range confList {
		if strings.HasPrefix(k, "redispool") {
			url := fmt.Sprintf("%s:%d", v.Host, v.Port)
			redisPool := NewPool(url, v.Password)
			SetRedisClt(k, redisPool)
			log.Infof("redis[%s] init, config: %+v", k, v)
		}
	}
	//redisPool := NewPool(server, password)
}

func getError(err error) error {
	log.Errorf("redis err : %v", err)
	log.Errorf("redis err debug : %v", string(debug.Stack()))
	return errors.New("1005")
}

//支持原生操作,注意使用完要conn.Close()
func (p Pool) Conn(ctx context.Context) (redis.Conn, error) {
	return p.redisPool.GetContext(ctx)
}

//string
func (p Pool) Set(key string, value string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) Del(keys []string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	param := []interface{}{}
	for _, v := range keys {
		param = append(param, v)
	}
	rs, err := redis.Int(conn.Do("DEL", param...))
	if err != nil {
		return 0, getError(err)
	}

	return rs, nil
}

/**
 * 第一个参数为是否设置成功，0失败1成功
 */
func (p Pool) SetNx(key string, value string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("SETNX", key, value))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) SetEx(key string, seconds int, value string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, seconds, value)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) PSetEx(key string, milliseconds int, value string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PSETEX", key, milliseconds, value)
	if err != nil {
		return getError(err)
	}

	return err
}

//返回OK表示设置成功
//返回""空字符串表示键名已存在
func (p Pool) SetPxNx(key string, value string, milliseconds int) (string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := conn.Do("SET", key, value, "PX", milliseconds, "NX")
	if err != nil {
		return "", getError(err)
	}
	if rs == nil {
		return "", nil
	}

	rsStr, err := redis.String(rs, err)

	return rsStr, err
}

func (p Pool) Get(key string) (value string, err error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := conn.Do("GET", key)
	if err != nil {
		value = ""
		err = getError(err)
		return
	}
	if rs == nil {
		value = ""
	} else {
		value, err = redis.String(rs, nil)
	}

	return
}

func (p Pool) GetSet(key string, value string) (oldValue string, err error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := conn.Do("GETSET", key, value)
	if err != nil {
		value = ""
		err = getError(err)
		return
	}
	if rs == nil {
		oldValue = ""
	} else {
		oldValue, err = redis.String(rs, nil)
	}

	return
}

func (p Pool) Incr(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) IncrBy(key string, increment int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("INCRBY", key, increment))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) Decr(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("DECR", key))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) DecrBy(key string, decrement int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("DECRBy", key, decrement))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) MSet(keyMapValues map[string]string) error {
	if len(keyMapValues) < 1 {
		return nil
	}

	conn := p.redisPool.Get()
	defer conn.Close()

	params := []interface{}{}
	for k, v := range keyMapValues {
		params = append(params, k, v)
	}

	_, err := conn.Do("MSET", params...)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) MGet(keys []string) ([]string, error) {
	if len(keys) < 1 {
		return []string{}, nil
	}

	conn := p.redisPool.Get()
	defer conn.Close()

	var keysInterface []interface{}
	for _, v := range keys {
		keysInterface = append(keysInterface, v)
	}
	rs, err := redis.Strings(conn.Do("MGET", keysInterface...))
	if err != nil {
		return nil, getError(err)
	}

	return rs, nil
}

//hash
func (p Pool) HSet(key string, field string, value string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", key, field, value)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) HSetNx(key string, field string, value string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("HSETNX", key, field, value))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) HGet(key string, field string) (string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := conn.Do("HGET", key, field)
	if err != nil {
		return "", getError(err)
	}

	if rs == nil {
		return "", err
	} else {
		return redis.String(rs, err)
	}
}

func (p Pool) HKeys(key string) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Strings(conn.Do("HKEYS", key))
	if err != nil {
		return rs, getError(err)
	}
	return rs, err
}

func (p Pool) HExists(key string, field string) (bool, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Bool(conn.Do("HEXISTS", key, field))
	if err != nil {
		return false, getError(err)
	}
	return rs, err
}

func (p Pool) HDel(key string, fields []string) error {
	if len(fields) < 1 {
		return nil
	}

	conn := p.redisPool.Get()
	defer conn.Close()

	params := []interface{}{}
	params = append(params, key)
	for _, v := range fields {
		params = append(params, v)
	}

	_, err := conn.Do("HDEL", params...)
	if err != nil {
		return getError(err)
	}
	return err
}

func (p Pool) HLen(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("HLEN", key))
	if err != nil {
		return 0, getError(err)
	}

	return rs, err
}

func (p Pool) HIncrBy(key string, field string, increment int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("HINCRBY", key, field, increment))
	if err != nil {
		return 0, getError(err)
	}
	return rs, err
}

func (p Pool) HMSet(key string, fieldMapValues map[string]string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	params := []interface{}{}
	params = append(params, key)
	for k, v := range fieldMapValues {
		params = append(params, k, v)
	}

	_, err := conn.Do("HMSET", params...)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) HMGet(key string, fields []string) (map[string]string, error) {
	if len(fields) < 1 {
		return make(map[string]string, 0), nil
	}
	conn := p.redisPool.Get()
	defer conn.Close()

	params := []interface{}{}
	params = append(params, key)
	for _, v := range fields {
		params = append(params, v)
	}
	rs, err := redis.Strings(conn.Do("HMGET", params...))
	if err != nil {
		return nil, getError(err)
	}

	rspData := map[string]string{}
	for k, v := range fields {
		rspData[v] = rs[k]
	}

	return rspData, nil
}

func (p Pool) HGetAll(key string) (map[string]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, getError(err)
	}

	return rs, nil
}

//list //待实现
func (p Pool) LPush(key string, value string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("LPUSH", key, value))
	if err != nil {
		return rs, getError(err)
	}

	return rs, nil
}

func (p Pool) RPush(key string, value string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("RPUSH", key, value))
	if err != nil {
		return rs, getError(err)
	}

	return rs, nil
}

func (p Pool) LPop(key string) (string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := conn.Do("LPOP", key)
	if err != nil {
		return "", getError(err)
	}

	if rs == nil {
		return "", err
	} else {
		return redis.String(rs, err)
	}
}

func (p Pool) RPop(key string) (string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := conn.Do("RPOP", key)
	if err != nil {
		return "", getError(err)
	}

	if rs == nil {
		return "", err
	} else {
		return redis.String(rs, err)
	}
}

func (p Pool) LLen(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("LLEN", key))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) LRange(key string, start int, stop int) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Strings(conn.Do("LRANGE", key, start, stop))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

//set //待实现
func (p Pool) SAdd(key string, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("SADD", key, member))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) SMembers(key string) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		return rs, getError(err)
	}
	return rs, err
}

func (p Pool) SIsMember(key string, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("SISMEMBER", key, member))
	if err != nil {
		return rs, getError(err)
	}
	return rs, err
}

func (p Pool) SPop(key string) (string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := conn.Do("SPOP", key)
	if err != nil {
		return "", getError(err)
	}

	if rs == nil {
		return "", err
	} else {
		return redis.String(rs, err)
	}
}

func (p Pool) SRandMember() {}

func (p Pool) SRem(key string, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("SREM", key, member))
	if err != nil {
		return rs, getError(err)
	}
	return rs, err
}

func (p Pool) SCard(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("SCARD", key))
	if err != nil {
		return rs, getError(err)
	}
	return rs, err
}

//zset
func (p Pool) ZAdd(key string, score int, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("ZADD", key, score, member))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZAddMembers(key string, memberScoreInt map[string]int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	if len(memberScoreInt) < 1 {
		return 0, nil
	}

	kvArr := make([]interface{}, 0)
	kvArr = append(kvArr, key)
	for k, v := range memberScoreInt {
		kvArr = append(kvArr, v)
		kvArr = append(kvArr, k)
	}

	rs, err := redis.Int(conn.Do("ZADD", kvArr...))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZScore(key string, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	doRs, err := conn.Do("ZSCORE", key, member)
	if err != nil {
		return 0, getError(err)
	}
	if doRs != nil {
		rs, _ := redis.Int(doRs, err)
		return rs, nil
	} else {
		return 0, nil
	}
}

func (p Pool) ZIncrBy(key string, increment int, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("ZINCRBY", key, increment, member))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZCard(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("ZCARD", key))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZCount(key string, min interface{}, max interface{}) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("ZCOUNT", key, min, max))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZRange(key string, start int, stop int) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Strings(conn.Do("ZRANGE", key, start, stop))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZRangeWithScores(key string, start int, stop int) ([]map[string]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rspArr := []map[string]string{}
	rs, err := redis.Strings(conn.Do("ZRANGE", key, start, stop, "WITHSCORES"))
	if err != nil {
		return rspArr, getError(err)
	}
	for i := 0; i < len(rs); i += 2 {
		key := rs[i]
		value := rs[i+1]
		rspArr = append(rspArr, map[string]string{
			"member": key,
			"value":  value,
		})
	}

	return rspArr, err
}

func (p Pool) ZRevRange(key string, start int, stop int) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Strings(conn.Do("ZREVRANGE", key, start, stop))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZRevRangeWithScores(key string, start int, stop int) ([]map[string]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rspArr := []map[string]string{}
	rs, err := redis.Strings(conn.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
	if err != nil {
		return rspArr, getError(err)
	}
	for i := 0; i < len(rs); i += 2 {
		key := rs[i]
		value := rs[i+1]
		rspArr = append(rspArr, map[string]string{
			"member": key,
			"value":  value,
		})
	}

	return rspArr, err
}

func (p Pool) ZRangeByScore(key string, min interface{}, max interface{}, limitOffset int, count int) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", fmt.Sprint(limitOffset), fmt.Sprint(count)))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

/**
 * 返回数组，每个元素都是map 包含member和value
 */
func (p Pool) ZRangeByScoreWithScores(key string, min interface{}, max interface{}, limitOffset int, count int) ([]map[string]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rspArr := []map[string]string{}
	rs, err := redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", fmt.Sprint(limitOffset), fmt.Sprint(count)))
	if err != nil {
		return rspArr, getError(err)
	}

	for i := 0; i < len(rs); i += 2 {
		key := rs[i]
		value := rs[i+1]
		rspArr = append(rspArr, map[string]string{
			"member": key,
			"value":  value,
		})
	}

	return rspArr, err
}

func (p Pool) ZRevRangeByScore(key string, max interface{}, min interface{}, limitOffset int, count int) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, max, min, "LIMIT", fmt.Sprint(limitOffset), fmt.Sprint(count)))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

/**
 * 返回数组，每个元素都是map 包含member和value
 */
func (p Pool) ZRevRangeByScoreWithScores(key string, max interface{}, min interface{}, limitOffset int, count int) ([]map[string]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rspArr := []map[string]string{}
	rs, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, max, min, "WITHSCORES", "LIMIT", fmt.Sprint(limitOffset), fmt.Sprint(count)))
	if err != nil {
		return rspArr, getError(err)
	}

	for i := 0; i < len(rs); i += 2 {
		key := rs[i]
		value := rs[i+1]
		rspArr = append(rspArr, map[string]string{
			"member": key,
			"value":  value,
		})
	}

	return rspArr, err
}

func (p Pool) ZRank(key string, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	doRs, err := conn.Do("ZRANK", key, member)
	if err != nil {
		return 0, getError(err)
	} else {
		rs, _ := redis.Int(doRs, err)
		return rs, nil
	}
}

func (p Pool) ZRevRank(key string, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	doRs, err := conn.Do("ZREVRANK", key, member)
	if err != nil {
		return 0, getError(err)
	} else {
		rs, _ := redis.Int(doRs, err)
		return rs, nil
	}
}

func (p Pool) ZRem(key string, member string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("ZREM", key, member))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZRemMembers(key string, members []string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rediParam := []interface{}{}
	rediParam = append(rediParam, key)
	for _, v := range members {
		rediParam = append(rediParam, v)
	}

	rs, err := redis.Int(conn.Do("ZREM", rediParam...))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZRemRangeByRank(key string, start int, stop int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("ZREMRANGEBYRANK", key, start, stop))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) ZRemRangeByScore(key string, min int, max int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("ZREMRANGEBYSCORE", key, min, max))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

//ZUNIONSTORE destination numkeys key [key ...]
func (p Pool) ZUnionStore(destKey string, num int, keys []string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	rediParam := []interface{}{}
	rediParam = append(rediParam, destKey, num)
	for _, v := range keys {
		rediParam = append(rediParam, v)
	}
	_, err := conn.Do("ZUNIONSTORE", rediParam...)
	if err != nil {
		return getError(err)
	}
	return nil
}

//bit
func (p Pool) SetBit(key string, offset int, value int) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := redis.Int(conn.Do("SETBIT", key, offset, value))
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) GetBit(key string, offset int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("GETBIT", key, offset))
	if err != nil {
		return 0, getError(err)
	}

	return rs, err
}

func (p Pool) BitCount(key string, start, end int) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("BITCOUNT", key, start, end))
	if err != nil {
		return 0, getError(err)
	}

	return rs, err
}

//自动过期
func (p Pool) Expire(key string, seconds int) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, seconds)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) ExpireAt(key string, timestamp int) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIREAT", key, timestamp)
	if err != nil {
		return getError(err)
	}

	return err
}

/**
 *返回剩余生存时间
 *当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以秒为单位，返回 key 的剩余生存时间。
 */
func (p Pool) Ttl(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("TTL", key))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) Persist(key string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PERSIST", key)
	if err != nil {
		return getError(err)
	}

	return err
}

/**
 *以毫秒为单位
 */
func (p Pool) PExpire(key string, milliseconds int) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PEXPIRE", key, milliseconds)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) PExpireAt(key string, millisecTimestamp int) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIREAT", key, millisecTimestamp)
	if err != nil {
		return getError(err)
	}

	return err
}

func (p Pool) PTtl(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Int(conn.Do("PTTL", key))
	if err != nil {
		return rs, getError(err)
	}

	return rs, err
}

func (p Pool) PipeLine(params []map[string]interface{}) ([]interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	var err error
	for _, param := range params {
		fStr := fmt.Sprint(param["func"])

		if argsSLice, ok := param["args"].([]interface{}); ok {
			err = conn.Send(fStr, argsSLice...)
			if err != nil {
				return nil, getError(err)
			}
		} else {
			return nil, errors.New("1003")
		}

	}

	err = conn.Flush()
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0)
	for _, paramValue := range params {
		fStr := strings.ToLower(fmt.Sprint(paramValue["func"]))
		args := paramValue["args"].([]interface{})

		if inArray(fStr, []interface{}{"zrange", "zrevrange", "zrangebyscore", "zrevrangebyscore"}) {
			rs, err := redis.Strings(conn.Receive())
			if err != nil {
				return nil, getError(err)
			}
			if len(args) > 3 && fmt.Sprint(args[3]) == "withscores" {
				result = append(result, transSliceToKv(rs))
			} else {
				result = append(result, rs)
			}
		} else if fStr == "hgetall" {
			rs, err := redis.StringMap(conn.Receive())
			if err != nil {
				return nil, getError(err)
			}
			result = append(result, rs)
		} else if fStr == "hmget" {
			rs, err := redis.Strings(conn.Receive())
			if err != nil {
				return nil, getError(err)
			}

			rsp := make(map[string]string, 0)
			if len(args) > 1 {
				keyList := args[1:]
				for index, value := range keyList {
					key, err := redis.String(value, nil)
					if err != nil {
						return nil, getError(err)
					}
					rsp[key] = rs[index]
				}
				result = append(result, rsp)
			} else {
				result = append(result, rsp)
			}
		} else if inArray(fStr, GetStringReturnCommand()) {
			rs, err := redis.String(conn.Receive())
			if err != nil && err != redis.ErrNil {
				return nil, getError(err)
			}
			result = append(result, rs)
		} else if inArray(fStr, GetIntReturnCommand()) {
			rs, err := redis.Int(conn.Receive())
			if err != nil && err != redis.ErrNil {
				return nil, getError(err)
			}
			result = append(result, rs)
		} else {
			return nil, errors.New("1002")
		}
	}

	return result, err
}

func GetStringReturnCommand() []interface{} {
	return []interface{}{"set", "setex", "psetex", "get", "getset", "mset", "hset", "hget", "rpop", "lpop", "spop", "hmset", "srandmember"}
}

func GetIntReturnCommand() []interface{} {
	return []interface{}{"del", "setnx", "setex", "incr", "incrby", "decr", "decyby", "hsetnx", "hdel", "hlen",
		"hincrby", "hexists", "lpush", "rpush", "llen", "sadd", "sismembers", "srem", "scard", "zadd", "zadd",
		"zscore", "zincrby", "zcard", "zcount", "zrank", "zrevrank", "zrem", "zremrangebyrank", "zremrangebyscore",
		"setbit", "getbit", "expire", "expireat", "ttl"}
}

func inArray(needle interface{}, arr []interface{}) bool {
	for _, value := range arr {
		if value == needle {
			return true
		}
	}
	return false
}

func transSliceToKv(strList []string) []map[string]string {
	rspArr := make([]map[string]string, 0)
	for i := 0; i < len(strList); i += 2 {
		key := strList[i]
		value := strList[i+1]
		rspArr = append(rspArr, map[string]string{
			"member": key,
			"value":  value,
		})
	}
	return rspArr
}

// RENAME
func (p Pool) RENAME(key1 string, key2 string) error {
	conn := p.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("RENAME", key1, key2)
	if err != nil {
		return getError(err)
	}

	return err
}

// EXISTS
func (p Pool) EXISTS(key string) (bool, error) {
	conn := p.redisPool.Get()
	defer conn.Close()

	rs, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, getError(err)
	}
	return rs, err
}
