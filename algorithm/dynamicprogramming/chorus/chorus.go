package main

import "fmt"

// 合唱队形
/*
链接：https://www.nowcoder.com/questionTerminal/cf209ca9ac994015b8caf5bf2cae5c98
来源：牛客网

N位同学站成一排，音乐老师要请其中的(N-K)位同学出列，使得剩下的K位同学不交换位置就能排成合唱队形。 合唱队形是指这样的一种队形：设K位同学从左到右依次编号为1, 2, …, K，他们的身高分别为T1, T2, …, TK， 则他们的身高满足T1 < T2 < … < Ti , Ti > Ti+1 > … > TK (1 <= i <= K)。 你的任务是，已知所有N位同学的身高，计算最少需要几位同学出列，可以使得剩下的同学排成合唱队形。
*/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findChorus(nums []int) int {
	n := len(nums)
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	for i := 0; i < n; i++ {
		dp1[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				//if i == 6 {
				//	fmt.Println(dp1[i], dp1[j]+1)
				//}
				dp1[i] = maxInt(dp1[i], dp1[j]+1)
			}
		}
	}
	fmt.Println("dp1: ", dp1)
	for i := n - 1; i >= 0; i-- {
		dp2[i] = 1
		for j := n - 1; j > i; j-- {
			if nums[j] > nums[i] {
				dp2[i] = maxInt(dp2[i], dp2[j]+1)
			}
		}
	}
	fmt.Println("dp2: ", dp2)

	res := 0
	for i := 0; i < n; i++ {
		res = maxInt(res, dp1[i]+dp2[i]-1)
	}
	return n - res
}

func main() {
	res := findChorus([]int{186, 186, 150, 200, 160, 130, 197, 220})
	fmt.Println(res)
}
