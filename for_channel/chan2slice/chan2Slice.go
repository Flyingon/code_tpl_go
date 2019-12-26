package main

import (
	"fmt"
	"strconv"
)

// 循环读取channel，取地址append到数组
type ChanStruct struct {
	Num  int
	Data string
	NumPtr *int
}

// 新数组地址全一样，不符合预期
func chan2SliceNotOK() {
	// 定义chan
	ptrChan := make(chan ChanStruct, 10)
	for i := 1; i < 10; i++ {
		num := i  // 注意，for循环取地址，都要新建变量，分配内存，否则i的地址都不会变，最后值只有最后一个
		cs := ChanStruct{
			Num:  i,
			Data: strconv.Itoa(i),
			NumPtr: &num,
		}
		ptrChan <- cs
	}
	close(ptrChan)
	var csList []*ChanStruct
	var NumPtrList []*int
	// 从chan取数据，放入csList中
	for c := range ptrChan {
		csList = append(csList, &c)
		NumPtrList = append(NumPtrList, c.NumPtr)
	}
	for _, cs := range csList {
		fmt.Printf("cs: %p, %d, %s\n", cs, cs.Num, cs.Data)
	}
	for _, numPtr := range NumPtrList {
		fmt.Printf("cs: %p, %d\n", numPtr, *numPtr)
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
