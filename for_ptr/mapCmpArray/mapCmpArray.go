package main

import (
	"fmt"
)

func updateArray(array []int, count int) {
	if count > 10 {
		return
	}
	//fmt.Println(count)
	array = append(array, count)
	count += 1
	updateArray(array, count)
}

func updateArrayPtr(arrayPtr *[]int, count int) {
	if count > 10 {
		return
	}
	//fmt.Println(count)
	*arrayPtr = append(*arrayPtr, count)
	count += 1
	updateArrayPtr(arrayPtr, count)
}

func main() {
	data := make([]int, 0)
	updateArray(data, 0)
	fmt.Println("data: ", data)

	updateArrayPtr(&data, 0)
	fmt.Println("data: ", data)

	dst := make([]int, 3)
	copy(dst, data[3:5]) // 不能给切片copy
	fmt.Println("data: ", data)
}
