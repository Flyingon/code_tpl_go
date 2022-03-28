package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"sync"
)

type MySyncMap struct {
	sync.Map
}

func (m MySyncMap) Print(k interface{}) {
	value, ok := m.Load(k)
	fmt.Printf("value ptr: %p\n", &value)
	fmt.Println("value: ", value, ok)
	valueStr, ok := value.(string)
	fmt.Printf("value str ptr: %p\n", &valueStr)
	fmt.Println("value str: ", valueStr, ok)
}

func (m MySyncMap) PrintJson() {
	mapJson, err := jsoniter.MarshalToString(&m)
	fmt.Printf("[JSON] %s, err: %v", mapJson, err)
}

func main() {
	var syncMap MySyncMap

	syncMap.Store("Key1", "Value1")
	syncMap.Store("Key2", "Value2")
	syncMap.Store("Key3", 2)
	syncMap.Store("Key1", "Value1")
	//syncMap.Store(4, 4)

	syncMap.Print("Key1")
	//syncMap.Print("Key1")
	//syncMap.Print("Key3")
	//syncMap.Print(4)
	//syncMap.Delete("Key1")
	//syncMap.Print("Key1")

	var ret []string
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		kStr := k.(string)
		ret = append(ret, kStr)
		return true
	})
	fmt.Printf("ret: %+v\n", ret)

	syncMap.PrintJson()
}
