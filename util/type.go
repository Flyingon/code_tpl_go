package util

import "strconv"

// 将interface{} 转换成 []byte, 支持string, float64和[]byte
func GetBytesByInterface(key interface{}) []byte {
	var ret string
	switch key.(type) {
	case string:
		ret = key.(string)
	case int:
		temp := key.(int)
		ret = strconv.Itoa(temp)
	case float64:
		ret = strconv.FormatFloat(key.(float64), 'f', -1, 64)
	case []byte:
		ret = string(key.([]byte))
	}
	return []byte(ret)
}

// 将interface{} 转换成 string
func GetStringFromInterface(key interface{}) string {
	var ret string
	switch key.(type) {
	case string:
		ret = key.(string)
	case int:
		ret = strconv.FormatInt(int64(key.(int)), 10)
	case int8:
		ret = strconv.FormatInt(int64(key.(int8)), 10)
	case int16:
		ret = strconv.FormatInt(int64(key.(int16)), 10)
	case int32:
		ret = strconv.FormatInt(int64(key.(int32)), 10)
	case int64:
		ret = strconv.FormatInt(key.(int64), 10)
	case uint:
		ret = strconv.FormatUint(uint64(key.(uint)), 10)
	case uint8:
		ret = strconv.FormatUint(uint64(key.(uint8)), 10)
	case uint16:
		ret = strconv.FormatUint(uint64(key.(uint16)), 10)
	case uint32:
		ret = strconv.FormatUint(uint64(key.(uint32)), 10)
	case uint64:
		ret = strconv.FormatUint(key.(uint64), 10)
	case float32:
		ret = strconv.FormatFloat(float64(key.(float32)), 'f', -1, 64)
	case float64:
		ret = strconv.FormatFloat(key.(float64), 'f', -1, 64)
	case bool:
		ret = strconv.FormatBool(key.(bool))
	case []byte:
		ret = string(key.([]byte))
	}
	return ret
}
