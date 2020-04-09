package main

import (
	"math/rand"
	"time"
)

// 随机生成指定位数的大写字母和数字的组合
func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机生成指定位数的大写字母和数字的组合
func GetRandomList(l int) interface{} {
	numList := []int {0,1,2,3,4,5,6,7,8,9}
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []interface{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		randNum := r.Intn(len(bytes))
		if randNum < 10 {
			result = append(result,numList[randNum])
		} else {
			result = append(result,numList[randNum])
		}
	}
	return result
}
