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

// 通用http返回
type CommonRsp map[string]interface{}

func (rsp *CommonRsp) GenResponse(data map[string]interface{}) error {
	fmt.Printf("rsp: %p, data: %p\n", *rsp, data)
	data["a"] = "b"
	return nil
}


func main() {
	//mapAssertNotExist()
	//
	//getChan, _:= dataMap["a"]
	//fmt.Printf("get chan ptr: %p\n", getChan)
	//
	//b := dataMap["c"]
	//fmt.Printf("b: %+v\n", b)
	rsp := &CommonRsp{}
	fmt.Printf("rsp init: %p\n", *rsp)
	rsp.GenResponse(*rsp)
	fmt.Printf("rsp after: %v\n", rsp)
}