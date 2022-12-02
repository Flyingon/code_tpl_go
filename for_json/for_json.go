package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strings"
)

type ABC struct {
	A string `json:"a"`
	B int    `json:"b"`
}

type DEF struct {
}

type JJJ struct {
	ABC `json:"abc"`
	DEF json.RawMessage `json:"def"`
}

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
	dataJson := `{
	   "abc": {
			"a": "1",
			"b": 3
		},
		"def": "xxadkajxk"
	}`
	jjj := JJJ{}
	err := json.Unmarshal([]byte(dataJson), &jjj)
	fmt.Println(err)
	fmt.Printf("%s", jjj.DEF)
}
