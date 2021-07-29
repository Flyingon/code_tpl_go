package main

import (
	"fmt"
	"math"
)

var val = 3.14151265357989

// ShowWith2Decimal 取两位精度打印
func ShowWith2Decimal() {
	s := fmt.Sprintf("%.2f", val)
	fmt.Println(s)
}

// RoundDown 保留两位，舍弃后面的
func RoundDown() {
	fmt.Println(math.Floor(val*100) / 100) // 3.14 (round down)
}

// RoundNearest 保留两位，四舍五入
func RoundNearest() {
	fmt.Println(math.Round(val*100) / 100) // 3.14 (round to nearest)
}

// RoundCeil 保留两位，后面向上进位
func RoundCeil() {
	fmt.Println(math.Ceil(val*100) / 100) // 3.15 (round down)
}

func main() {
	ShowWith2Decimal()
	RoundDown()
	RoundNearest()
	RoundCeil()
}
