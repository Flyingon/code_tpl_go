package main

import "fmt"

var dataMap = make(map[string]interface{})

func init(){
	dataMap["nil"] = nil
}

func WillNotPanic() {
	var d string
	d, _ = dataMap["nil"].(string)
	fmt.Printf("dataMap.nil: %s", d)
}

func WillPanic() {
	var d string
	d = dataMap["nil"].(string)
	fmt.Printf("dataMap.nil: %s", d)
}

func WillPanic2() {
	d := dataMap["nil"].(string)
	fmt.Printf("dataMap.nil: %s", d)
}

func main() {
	WillNotPanic()
	WillPanic2()
	WillPanic()
}
