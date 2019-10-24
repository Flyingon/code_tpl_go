package main

import (
	"../util"
	"fmt"
	"reflect"
)

func main() {
	contentStr := `{"a": 1, "b": "2", "c": {"d": 3, "e": "4"}}`
	contentMap := make(map[string]interface{})
	err := util.JSONUnMarshal(util.StringToBytesFast(contentStr), &contentMap)
	fmt.Println(err, contentMap)
	for k, v := range contentMap {
		fmt.Println(k, reflect.TypeOf(v))
	}
}
