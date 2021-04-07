package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// JSONMarshal json序列化，escape设置为false
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// JSONUnMarshal 设置UseNumber,防止uint64位精度缺失
func JSONUnMarshal(jsonStream []byte, ret interface{}) error {
	decoder := json.NewDecoder(strings.NewReader(string(jsonStream)))
	decoder.UseNumber()
	if err := decoder.Decode(&ret); err != nil {
		fmt.Println("error:", err)
		return err
	}
	return nil
}

// ValueToStr JSONUnMarshal到map[string]interface{}后interface{} to string
func ValueToStr(v interface{}) (ret string) {
	switch v.(type) {
	case string:
		ret = v.(string)
	case json.Number:
		ret = v.(json.Number).String()
	default:
		val, _ := JSONMarshal(v)
		ret = string(val)
	}
	return
}

// IsEmptyValue
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
