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
