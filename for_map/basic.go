package main

import "fmt"

var dataMap = make(map[string]interface{})

func init(){
	dataMap["nil"] = nil
	chanInt := make(chan int)
	fmt.Printf("chan ptr: %p\n", chanInt)
	dataMap["a"] = chanInt
}

func mapAssertNotExist() {
	d, ok := dataMap["nil"].(string)
	fmt.Printf("dataMap.nil: %s, ok: %v", d, ok)
}

func main() {
	mapAssertNotExist()

	getChan, _:= dataMap["a"]
	fmt.Printf("get chan ptr: %p\n", getChan)
}