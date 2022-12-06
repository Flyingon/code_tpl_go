package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ip2addr(ipStr string) {
	ipList := strings.Split(ipStr, ".")
	num := int64(0)
	for _, elem := range ipList {
		elemInt, _ := strconv.ParseInt(elem, 10, 64)
		fmt.Println(elemInt << 8)
		fmt.Println(elemInt % 255)
		num = num<<8 + elemInt%255
	}
	fmt.Println(num)
}

func main() {
	//ip2addr("192.168.0.1")
	fmt.Println(math.MaxInt32 / 86400)
	fmt.Println()
}
