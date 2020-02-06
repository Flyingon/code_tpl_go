package main

import "fmt"

var dataMap = make(map[string]interface{})

func init(){
	dataMap["nil"] = nil
}

func mapAssertNotExist() {
	d, ok := dataMap["nil"].(string)
	fmt.Printf("dataMap.nil: %s, ok: %v", d, ok)
}

func main() {
	mapAssertNotExist()
}