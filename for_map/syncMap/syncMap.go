package main

import (
	"fmt"
	"sync"
)


type MySyncMap struct {
	sync.Map
}

func (m MySyncMap) Print(k interface{}) {
	value, ok := m.Load(k)
	fmt.Println(value, ok)
}

func main() {
	var syncMap MySyncMap

	syncMap.Store("Key1", "Value1")
	syncMap.Store("Key2", "Value2")
	syncMap.Store("Key3", 2)
	syncMap.Store("Key1", "Value1")
	syncMap.Store(4, 4)

	//syncMap.Print("Key1")
	//syncMap.Print("Key1")
	//syncMap.Print("Key3")
	//syncMap.Print(4)
	//syncMap.Delete("Key1")
	//syncMap.Print("Key1")

	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}