package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Addr2Int ...
func Addr2Int(ipAddr string) int {
	ret := make([][]int8, 4)
	elemList := strings.Split(ipAddr, ".")
	for index, elem := range elemList {
		elemInt, _ := strconv.ParseInt(elem, 10, 64)
		ret[index] = intToBit(int(elemInt))
		//fmt.Println(elemInt)
	}
	fmt.Println(ret)
	return 0
}

var bitArray = []int{
	256, 128, 64, 16, 8, 4, 2, 1,
}

func intToBit(num int) []int8 {
	ret := make([]int8, 8)
	elems := make(map[int]bool)
	for num > 0 {
		for _, bitNum := range bitArray {
			if num >= bitNum {
				elems[bitNum] = true
				num = num - bitNum
			}
		}
	}
	//fmt.Println(elems)
	for index, bitNum := range bitArray {
		if elems[bitNum] {
			ret[index] = 1
		}
	}
	return ret
}

func main() {
	//Addr2Int("1.1.1.2")
	Addr2Int("196.1.10.255")
	//fmt.Println(intToBit(1))
}
