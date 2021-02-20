package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strings"
)



func main2() {
	dataJson := `{"a": "123", "b": 11111111111111111111111}`
	commonMap := make(map[string]interface{})

	decoder := jsoniter.NewDecoder(strings.NewReader(dataJson))
	decoder.UseNumber()
	err := decoder.Decode(&commonMap)
	fmt.Println(err)
	fmt.Println(commonMap)
	fmt.Println(reflect.TypeOf(commonMap["a"]))
	fmt.Println(reflect.TypeOf(commonMap["b"]))
	data, _ := commonMap["c"].(json.Number)
	fmt.Println("data: ", data.String())
}

func main() {
	var params []interface{}
	dataJson := `["user:info:100422","total_credits",4653000]`
	jsoniter.Unmarshal([]byte(dataJson), &params)
	fmt.Println(params)
}
