package main

import (
	"fmt"
	"strings"
)

// CheckKeyExist 检查keys中的key是否都存在于data中
func CheckKeyExist(data map[string]interface{}, keys []string) error {
	lackKeys := make([]string, 0, len(data))
	for _, key := range keys {
		if _, exist := data[key]; !exist {
			lackKeys = append(lackKeys, key)
		}
	}
	if len(lackKeys) > 0 {
		return fmt.Errorf("keys[%s] is not exist", strings.Join(lackKeys, ","))
	}
	return nil
}

func main() {
	data := "buyGoods"
	dataList := strings.Split(data, "/")
	fmt.Println(dataList[len(dataList) -1:])
}



