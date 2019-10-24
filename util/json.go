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
