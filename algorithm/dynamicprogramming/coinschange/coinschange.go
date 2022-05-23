package main

import (
	"fmt"
	"math"
)

/*
零钱兑换问题:
leetcode-322: https://leetcode.cn/problems/coin-change/
*/

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dpv0 暴力破解
/*
时间复杂度: len(coins) ^ amount
leetcode: 超出时间限制
*/
func dpv0(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dpv0(coins, amount-coin)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		res = minInt(res, subProblem+1)
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

// dpv1 添加备忘录，记录递归时相同入参的返回值
/*
时间复杂度: len(coins) * amount
*/
var memo = map[int]int{}

func dpv1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// 查备忘录，防止重复计算
	if _, exist := memo[amount]; exist {
		return memo[amount]
	}

	res := math.MaxInt
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dpv1(coins, amount-coin)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		res = minInt(res, subProblem+1)
	}
	// 把计算结果存入备忘录
	memo[amount] = res
	if res == math.MaxInt {
		memo[amount] = -1
	}
	return memo[amount]
}

// coinChange
/*
执行用时：8 ms
内存消耗：6.3 MB
通过测试用例：
188 / 188
*/
func coinChange(coins []int, amount int) int {
	//return dpv0(coins, amount)
	//return dpv1(coins, amount)

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // 设置为不可能的最大值，方便后面求最小值
	}
	// base case
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = minInt(dp[i], 1+dp[i-coin])
		}
		fmt.Printf("金额[%d]: %v\n", i, dp)
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	res := coinChange(coins, amount)
	fmt.Println(res)
}
