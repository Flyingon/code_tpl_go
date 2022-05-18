package main

import (
	"code_tpl_go/algorithm/binarytree/internal"
	"fmt"
	"github.com/Flyingon/go-common/util"
)

// LevelOrderComplete 层序遍历，按完全二叉树返回
/*
时间复杂度：每个点进队出队各一次，故渐进时间复杂度为 O(n)
空间复杂度：队列中元素的个数不超过 n 个，故渐进空间复杂度为 O(n)。
*/
func LevelOrderComplete(root *internal.TreeNode) [][]string {
	ret := make([][]string, 0) // 初始化返回数组
	if root == nil {
		return ret
	}
	queue := []*internal.TreeNode{root} // 循环开始，第一次只有一个root
	levelIndex := 0                     // 层次索引
	levelNum := 1                       // 每层元素个数
	for levelNum > 0 {                  // 按层进行遍历，只要这层有数据，就跑一遍
		ret = append(ret, []string{})                         // 记录当前层的结果
		nextQueue := make([]*internal.TreeNode, 2*len(queue)) // 记录下一层队列
		levelNum = 0

		for j := 0; j < len(queue); j++ { // 遍历当前层
			node := queue[j]

			if node == nil {
				ret[levelIndex] = append(ret[levelIndex], "null")
				continue
			}

			ret[levelIndex] = append(ret[levelIndex], fmt.Sprint(node.Data)) // 记录数据
			// 记录下一层
			if node.Left != nil {
				nextQueue[2*j] = node.Left
				levelNum += 1
			}
			if node.Right != nil {
				nextQueue[2*j+1] = node.Right
				levelNum += 1
			}
		}
		queue = nextQueue
		levelIndex += 1
	}
	return ret
}

// LevelOrder 层序遍历
/*
时间复杂度：每个点进队出队各一次，故渐进时间复杂度为 O(n)。
空间复杂度：队列中元素的个数不超过 n 个，故渐进空间复杂度为 O(n)。
*/
func LevelOrder(root *internal.TreeNode) [][]int {
	ret := make([][]int, 0) // 初始化返回数组
	if root == nil {
		return ret
	}
	queue := []*internal.TreeNode{root} // 循环开始，第一次只有一个root
	levelIndex := 0                     // 层次索引
	for len(queue) > 0 {                // 按层进行遍历，只要这层有数据，就跑一遍
		ret = append(ret, []int{})         // 记录当前层的结果
		var nextQueue []*internal.TreeNode // 记录下一层的节点
		for j := 0; j < len(queue); j++ {  // 遍历当前层
			node := queue[j]
			ret[levelIndex] = append(ret[levelIndex], util.InterfaceToInt(node.Data)) // 记录数据
			if node.Left != nil {                                                     // 记录下一层
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
		levelIndex += 1
	}
	return ret
}

func main() {
	nodeList := []interface{}{3, 9, 20, -1, -1, 15, 7, 10, 11, 12}
	newTree := internal.BuildBinTree(nodeList, 0, len(nodeList))

	levelRes := LevelOrder(newTree)
	fmt.Printf("层序遍历: %v\n", levelRes)
	for _, level := range levelRes {
		fmt.Println(level)
	}

	levelCompleteRes := LevelOrderComplete(newTree)
	fmt.Printf("层序遍历: %v\n", levelCompleteRes)
	for _, level := range levelCompleteRes {
		fmt.Println(level)
	}
}
