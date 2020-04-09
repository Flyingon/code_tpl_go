package main

import (
	"fmt"
	"unsafe"
)

type mapData struct {
	DataStr string
	DataMap map[string]interface{}
}

var dataMap1 = map[int]*mapData{
	1: {"a", map[string]interface{}{"a": "a"}},
	2: {"b",map[string]interface{}{"b": "b"}},
	3: {"c",map[string]interface{}{"b": "b"}},
}

func getMapDataDirect(i int) *mapData {
	return dataMap1[i]
}

func getMapDataCopy(i int) *mapData {
	mapData := mapData{}
	return &mapData
}

func getMapDataCopyByUnsafe(i int) *mapData {
	var mapData = *(*mapData)(unsafe.Pointer(&*dataMap1[i]))
	return &mapData
}

func main() {
	fmt.Printf("dataMap1.1 adr: %p, dataMap1.1.DataMap adr: %p\n", dataMap1[1], dataMap1[1].DataMap)
	data1 := getMapDataDirect(1)
	fmt.Printf("data1 adr: %p, data1.DataMap adr: %p\n", data1, data1.DataMap)
	data2 := getMapDataCopy(1)
	fmt.Printf("data2 adr: %p, data2.DataMap adr: %p\n", data2, data2.DataMap)
	data3 := getMapDataCopyByUnsafe(1)
	data3.DataMap["3"] = "3"
	fmt.Printf("data3 adr: %p, data3.DataMap adr: %p\n, dataMap1.1.DataMap: %+v\n", data3, data3.DataMap, dataMap1[1].DataMap)
}
