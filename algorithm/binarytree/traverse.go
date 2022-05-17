package main

import (
	"code_tpl_go/algorithm/binarytree/internal"
	"fmt"
)

// LevelOrder 层序遍历
/*
时间复杂度：每个点进队出队各一次，故渐进时间复杂度为 O(n)O(n)。
空间复杂度：队列中元素的个数不超过 nn 个，故渐进空间复杂度为 O(n)O(n)。
*/
func LevelOrder(root *internal.Node) [][]int {
	ret := make([][]int, 0) // 初始化返回数组
	if root == nil {
		return ret
	}
	queue := []*internal.Node{root} // 循环开始，第一次只有一个root
	levelIndex := 0                 // 层次索引
	for len(queue) > 0 {            // 按层进行遍历，只要这层有数据，就跑一遍
		ret = append(ret, []int{})        // 记录当前层的结果
		var nextQueue []*internal.Node    // 记录下一层的节点
		for j := 0; j < len(queue); j++ { // 遍历当前层
			node := queue[j]
			ret[levelIndex] = append(ret[levelIndex], node.Data) // 记录数据
			if node.Left != nil {                                // 记录下一层
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
	newTree := internal.GenOrderBinTreeRandom(1, 5, 100)
	levelRes := LevelOrder(newTree)
	fmt.Printf("层序遍历: %v\n", levelRes)
	for _, level := range levelRes {
		fmt.Println(level)
	}
}
