package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLengthOfLCIS(nums []int) int {
	start := 0
	res := 0
	for i, v := range nums {
		if i > 0 && v <= nums[i-1] {
			start = i
		}
		res = maxInt(res, i-start+1)
	}
	return res
}

func main() {
	res := findLengthOfLCIS([]int{1, 3, 5, 4, 7})
	fmt.Println(res)
}
