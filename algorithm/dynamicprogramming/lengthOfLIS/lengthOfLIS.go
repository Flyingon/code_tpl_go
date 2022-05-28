package main

import "fmt"

/*
dp[i]=max(dp[j])+1,其中0≤j<i且num[j]<num[i]
*/

func lengthOfLIS(nums []int) int {
	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	d := make([]int, len(nums))
	d[0] = 0
	for i, v := range nums {
		for j := 0; j < i; j++ {
			if v > nums[j] {
				d[i] = maxInt(d[i], d[j]+1)
				if d[i] > res {
					res = d[i]
				}
			}
		}
		//fmt.Println(d)
	}
	return res + 1
}

func main() {
	//res := lengthOfLIS([]int{10,9,2,5,3,7,101,18})
	res := lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	//res := lengthOfLIS([]int{4,10,4,3,8,9})
	fmt.Println(res)
}
