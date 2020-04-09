package main

import (
	"fmt"
	"sync"
)

type Elem struct {
	Data string
}

// ClosureInForV1 不符合预期的使用
func ClosureInForV1(dataList []*Elem) {
	wg := sync.WaitGroup{}
	wg.Add(len(dataList))
	for i, v := range dataList {
		go func() {
			defer wg.Done()
			fmt.Printf("cycle %d time, index addr: %p, data addr: %p, value: %s\n", i, &i, &(v.Data), v.Data)
		}()
	}
	wg.Wait()
}

// ClosureInForV2 符合预期的使用
func ClosureInForV2(dataList []*Elem) {
	wg := sync.WaitGroup{}
	wg.Add(len(dataList))
	for i, v := range dataList {
		go func(index int, elem *Elem) {
			defer wg.Done()
			fmt.Printf("cycle %d time, index addr: %p, data addr: %p, value: %s\n", index, &index, &(elem.Data), elem.Data)
		}(i, v)
	}
	wg.Wait()
}

// ClosureInForV3 不符合预期的使用
func ClosureInForV3(dataList []*Elem) {
	wg := sync.WaitGroup{}
	wg.Add(len(dataList))
	for i, v := range dataList {
		go DataPrintV3(&wg, &i, v)
	}
	wg.Wait()
}

// ClosureInForV4 符合预期的使用
func ClosureInForV4(dataList []*Elem) {
	wg := sync.WaitGroup{}
	wg.Add(len(dataList))
	for i, v := range dataList {
		go DataPrintV4(&wg, i, v)
	}
	wg.Wait()
}

func DataPrintV3(wg *sync.WaitGroup, index *int, elem *Elem) {
	defer wg.Done()
	fmt.Printf("cycle %d time, index addr: %p, data addr: %p, value: %s\n", *index, index, &(elem.Data), elem.Data)
}

func DataPrintV4(wg *sync.WaitGroup, index int, elem *Elem) {
	defer wg.Done()
	fmt.Printf("cycle %d time, index addr: %p, data addr: %p, value: %s\n", index, &index, &(elem.Data), elem.Data)
}


func main() {
	var tArray []*Elem
	tArray = append(tArray, &Elem{"value1"})
	tArray = append(tArray, &Elem{"value2"})
	tArray = append(tArray, &Elem{"value3"})
	tArray = append(tArray, &Elem{"value4"})
	tArray = append(tArray, &Elem{"value5"})
	fmt.Printf("------v1-------\n")
	ClosureInForV1(tArray)
	fmt.Printf("------v2-------\n")
	ClosureInForV2(tArray)
	fmt.Printf("------v3-------\n")
	ClosureInForV3(tArray)
	fmt.Printf("------v4-------\n")
	ClosureInForV4(tArray)
}
