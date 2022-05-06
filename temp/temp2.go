package main

/*
	Inter to Roman
*/

import "fmt"

var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

var symbolMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",

	4:   "IV",
	9:   "IX",
	40:  "XL",
	90:  "XC",
	400: "CD",
	900: "CM",
}

// 请根据题目要求确定返回值类型和参数列表(输入)
func solution(inNum int) string {
	tempNum := inNum
	var elemList []map[string]int
	for _, num := range nums {
		d := tempNum / num
		if d > 0 {
			tempNum = tempNum - d*num
			elemList = append(elemList, map[string]int{
				symbolMap[num]: d,
			})
		}
	}
	ret := ""
	for _, elem := range elemList {
		for k, v := range elem {
			fmt.Println(k, v)
			for i := 0; i < v; i++ {
				ret += k
			}
		}
	}

	return ret
}

func main() {
	for _, num := range []int{
		3, 58, 1994, 3999, 2000, 2001,
	} {
		ret := solution(num)
		fmt.Printf("%d: %s\n", num, ret)
	}

}
