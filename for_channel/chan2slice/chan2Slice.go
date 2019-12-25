package main

import (
	"fmt"
	"strconv"
)

// 循环读取channel，取地址append到数组
type ChanStruct struct {
	Num  int
	Data string
}

// 新数组地址全一样，不符合预期
func chan2SliceNotOK() {
	// 定义chan
	ptrChan := make(chan ChanStruct, 10)
	for i := 1; i < 10; i++ {
		cs := ChanStruct{
			Num:  i,
			Data: strconv.Itoa(i),
		}
		ptrChan <- cs
	}
	close(ptrChan)
	var csList []*ChanStruct
	// 从chan取数据，放入csList中
	for c := range ptrChan {
		csList = append(csList, &c)
	}
	for _, cs := range csList {
		fmt.Printf("cs: %p, %d, %s\n", cs, cs.Num, cs.Data)
	}
}

// 新数组地址不同，符合预期
func chan2SliceOK() {
	// 定义chan
	ptrChan := make(chan ChanStruct, 10)
	for i := 1; i < 10; i++ {
		cs := ChanStruct{
			Num:  i,
			Data: strconv.Itoa(i),
		}
		ptrChan <- cs
	}
	close(ptrChan)
	var csList []*ChanStruct
	// 从chan取数据，放入csList中
	for c := range ptrChan {
		temp := c
		csList = append(csList, &temp)
	}
	for _, cs := range csList {

		fmt.Printf("cs: %p, %d, %s\n", cs, cs.Num, cs.Data)
	}
}

func main() {
	chan2SliceNotOK()
	chan2SliceOK()
}
