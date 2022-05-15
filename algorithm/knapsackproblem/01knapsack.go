package main

import (
	"fmt"
	"math"
)

// Knapsack01Solved 01背包问题基本解法
func Knapsack01Solved(weightList, valueList []int, bagWeight int) int {
	//var dp [len(weightList)][bagWeight]int
	dp := make([][]int, len(weightList))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
		if i == 0 {
			for j := range dp[i] {
				if weightList[i] <=j {
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
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-weightList[i]]+valueList[i])))
			}
		}
		fmt.Printf("获取前%d个物品的dp: %v\n", i+1, dp)
	}
	return dp[len(weightList)-1][bagWeight]
}

func main() {
	weightList := []int{1, 3, 4}
	valueList := []int{15, 20, 30}
	packWeight := 4
	gotVal := Knapsack01Solved(weightList, valueList, packWeight)
	fmt.Println("got value: ", gotVal)
}
