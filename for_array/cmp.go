package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxListLen = 5000
)

func IsListContainList(args ...interface{}) (interface{}, error) {
	listA, ok := args[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("args.0 type is not []interface{}")
	}
	listB, ok := args[1].([]interface{})
	if !ok {
		return nil, fmt.Errorf("args.1 type is not listB")
	}
	if len(listB) > maxListLen || len(listA) > maxListLen {
		return nil, fmt.Errorf("list len is exceed max[%d]", maxListLen)
	}
	if len(listB) > len(listA) {
		return false, nil
	}
	mapA := make(map[interface{}]bool)
	for _, k := range listA {
		mapA[k] = true
	}
	ret := true
	for _, e := range listB {
		if !mapA[e] {
			ret = false
		}
	}
	return ret, nil
}

// 随机生成指定位数的大写字母和数字的组合
func GenRandomList(l int) interface{} {
	numList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []interface{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		randNum := r.Intn(len(bytes))
		if randNum < 10 {
			result = append(result, numList[randNum])
		} else {
			result = append(result, string(bytes[randNum]))
		}
	}
	return result
}

func main() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		listA := GenRandomList(20)
		listB := GenRandomList(10)
		fmt.Println("listA: ", listA)
		fmt.Println("listB: ", listB)
		r, e := IsListContainList(listA, listB)
		fmt.Println(r, e)
	}
	subTtime := time.Now().Sub(startTime).Milliseconds()
	fmt.Printf("time cost: %d ms\n", subTtime)
}
