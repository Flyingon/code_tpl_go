package main

import (
	"code_tpl_go/algorithm/dynamicprogramming/knapsackproblem/internal"
)

/*
完全背包问题，背包内物品数量不限制
*/

import "fmt"

// KnapsackCompleteSolved 完全背包问题基本解法
// 时间复杂度: O(len(weightList) * bagWeight)
// 空间复杂度: O(bagWeight)
func KnapsackCompleteSolved(weightList, valueList []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	fmt.Println("初始化的dp: ", dp)
	for i := 0; i < len(weightList); i++ { // 遍历物品
		for j := weightList[i]; j <= bagWeight; j++ { // 遍历背包容量
			dp[j] = internal.MaxInt(dp[j], dp[j-weightList[i]]+valueList[i])
		}
		fmt.Printf("获取前%d个物品的dp: %v\n", i+1, dp)
	}
	return dp[bagWeight]
}

func main() {
	weightList := []int{1, 3, 4}
	valueList := []int{15, 20, 30}
	packWeight := 4
	fmt.Println("------------------ 解法一 ------------------ ")
	gotVal := KnapsackCompleteSolved(weightList, valueList, packWeight)
	fmt.Println("got value: ", gotVal)
}
