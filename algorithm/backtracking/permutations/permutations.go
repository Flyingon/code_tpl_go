package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	var track []int
	used := make(map[int]bool)
	backtrack(nums, track, used)
	return res
}

func backtrack(nums []int, track []int, used map[int]bool) {
	if len(track) == len(nums) {
		//fmt.Printf("加入： %+v 前： %+v\n", track, res)
		temp := make([]int, len(track))
		for index, value := range track {
			temp[index] = value
		}
		res = append(res, temp)
		//fmt.Printf("加入： %+v 后： %+v\n", track, res)
		return
	}
	for index, num := range nums {
		if used[index] {
			continue
		}
		track = append(track, num)
		used[index] = true
		//fmt.Printf("递归前: %v\n", track)
		backtrack(nums, track, used)
		track = track[0 : len(track)-1]
		delete(used, index)
		//fmt.Printf("递归后: %v\n", track)
	}
}

func main() {
	nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Println(res)
}
