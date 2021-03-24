package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

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

// InterfaceToString 将interface{} 转换成 string
func InterfaceToString(key interface{}) string {
	if key == nil { // nil返回空字符串，根据需求添加
		return ""
	}
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
	case json.Number:
		ret = key.(json.Number).String()
	default:
		retBytes, _ := json.Marshal(key)
		ret = string(retBytes)
	}
	return ret
}

// TranStringToType 转换string类型到指定类型
func TranStringToType(valStr, keyType string) (interface{}, error) {
	var valFmt interface{}
	var err error
	switch keyType {
	case "int":
		valFmt, err = strconv.ParseInt(valStr, 10, 0)
	case "int8":
		valFmt, err = strconv.ParseInt(valStr, 10, 8)
	case "int16":
		valFmt, err = strconv.ParseInt(valStr, 10, 16)
	case "int32":
		valFmt, err = strconv.ParseInt(valStr, 10, 32)
	case "int64":
		valFmt, err = strconv.ParseInt(valStr, 10, 64)
	case "uint":
		valFmt, err = strconv.ParseUint(valStr, 10, 0)
	case "uint8":
		valFmt, err = strconv.ParseUint(valStr, 10, 8)
	case "uint16":
		valFmt, err = strconv.ParseUint(valStr, 10, 16)
	case "uint32":
		valFmt, err = strconv.ParseUint(valStr, 10, 32)
	case "uint64":
		valFmt, err = strconv.ParseUint(valStr, 10, 64)
	case "float", "float32":
		valFmt, err = strconv.ParseFloat(valStr, 32)
	case "float64":
		valFmt, err = strconv.ParseFloat(valStr, 64)
	case "bool":
		valFmt, err = strconv.ParseBool(valStr)
	default:
		err = fmt.Errorf("invalid type: %v", keyType)
	}
	if err != nil {
		return nil, err
	}
	return valFmt, nil
}

// GetZeroValueByType 获得指定类型到0值
func GetZeroValueByType(keyType string) interface{} {
	var ret interface{}
	switch keyType {
	case "string":
		ret = ""
	case "int":
		ret = 0
	case "int8":
		ret = int8(0)
	case "int16":
		ret = int16(0)
	case "int32":
		ret = int32(0)
	case "int64":
		ret = int64(0)
	case "uint":
		ret = uint(0)
	case "uint8":
		ret = uint8(0)
	case "uint16":
		ret = uint16(0)
	case "uint32":
		ret = uint32(0)
	case "uint64":
		ret = uint64(0)
	case "float", "float32":
		ret = float32(0)
	case "float64":
		ret = float64(0)
	case "bool":
		ret = false
	default:
		ret = ""
	}
	return ret
}

// InterfaceToInt 将interface{} 转换成 int
// 强制转换，忽略精度丢失; 忽略忽略错误，错误返回0
func InterfaceToInt(key interface{}) int {
	if key == nil { // nil返回零值，根据需求添加
		return 0
	}
	var ret int
	switch key.(type) {
	case string:
		tmp, _ := strconv.ParseFloat(key.(string), 64)
		ret = int(tmp)
	case int32:
		ret = int(key.(int32))
	case int64:
		ret = int(key.(int64))
	case int: //todo 是否需要统一所有int int64
		ret = key.(int)
	case float32:
		ret = int(key.(float32))
	case float64:
		ret = int(key.(float64))
	default:
	}
	return ret
}

// IsEmptyValue 零值判断
func IsEmptyValue(key interface{}) bool {
	switch key.(type) {
	case string:
		return len(key.(string)) == 0
	case int:
		return key.(int) == 0
	case int8:
		return key.(int8) == 0
	case int16:
		return key.(int16) == 0
	case int32:
		return key.(int32) == 0
	case int64:
		return key.(int64) == 0
	case uint:
		return key.(uint) == 0
	case uint8:
		return key.(uint8) == 0
	case uint16:
		return key.(uint16) == 0
	case uint32:
		return key.(uint32) == 0
	case uint64:
		return key.(uint64) == 0
	case float32:
		return key.(float32) == 0
	case float64:
		return key.(float64) == 0
	case bool:
		return key.(bool) == false
	case []byte:
		return len(key.([]byte)) < 2
	case json.Number:
		return key.(json.Number).String() == "0"
	}
	return false
}

// InterfaceCopy 结构体拷贝-注意: 传入结构体对象, 传指针只能复制地址，没有意义
func InterfaceCopy(d interface{}) interface{} {
	n := reflect.New(reflect.TypeOf(d))
	n.Elem().Set(reflect.ValueOf(d))
	return n.Interface()
}
