package main

import (
	"code_tpl_go/algorithm/dynamicprogramming/knapsackproblem/internal"
	"fmt"
)

// Knapsack01Solved 01背包问题基本解法
// 时间复杂度: O(len(weightList) * bagWeight)
// 空间复杂度: O(len(weightList) * bagWeight)
func Knapsack01Solved(weightList, valueList []int, bagWeight int) int {
	//var dp [len(weightList)][bagWeight]int
	dp := make([][]int, len(weightList))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
		if i == 0 {
			for j := range dp[i] {
				if weightList[i] <= j {
					dp[i][j] = valueList[i]
				}
			}
		}
	}
	fmt.Println("初始化的dp: ", dp)
	for i := 1; i < len(weightList); i++ { // 遍历物品
		for j := 0; j <= bagWeight; j++ { // 遍历背包容量
			if weightList[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = internal.MaxInt(dp[i-1][j], dp[i-1][j-weightList[i]]+valueList[i])
			}
		}
		fmt.Printf("获取前%d个物品的dp: %v\n", i+1, dp)
	}
	return dp[len(weightList)-1][bagWeight]
}

// Knapsack01SolvedV2 01背包问题基本解法-优化内存
// 时间复杂度: O(len(weightList) * bagWeight)
// 空间复杂度: O(bagWeight)
func Knapsack01SolvedV2(weightList, valueList []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	fmt.Println("初始化的dp: ", dp)
	for i := 0; i < len(weightList); i++ { // 遍历物品
		for j := bagWeight; j >= weightList[i]; j-- { // 遍历背包容量
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
	gotVal := Knapsack01Solved(weightList, valueList, packWeight)
	fmt.Println("got value: ", gotVal)
	fmt.Println("------------------ 解法二 ------------------ ")
	gotValV2 := Knapsack01SolvedV2(weightList, valueList, packWeight)
	fmt.Println("v2 got value: ", gotValV2)
}
