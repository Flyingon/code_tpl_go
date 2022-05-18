package main

import (
	"code_tpl_go/algorithm/binarytree/internal"
	"fmt"
)

// maxDepth 求树高度，深度优先遍历
/*
时间复杂度：O(n)
空间复杂度: O(height),递归所需的栈空间
leetcode-104: https://leetcode.cn/problems/maximum-depth-of-binary-tree/
*/
func maxDepth(root *internal.TreeNode) int {
	if root == nil {
		return 0
	}
	// 利用定义，计算左右子树的最大深度
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	// 整棵树的最大深度等于左右子树的最大深度取最大值，
	// 然后再加上根节点自己
	return internal.MaxInt(leftMax, rightMax) + 1
}

// maxDepthV2 遍历二叉树
func maxDepthV2(root *internal.TreeNode) int {
	if root == nil {
		return 0
	}
	mf := maxDepthFlow{}
	mf.traverse(root)
	return mf.Ret
}

type maxDepthFlow struct {
	Ret   int
	Depth int
}

func (m *maxDepthFlow) traverse(root *internal.TreeNode) {
	if root == nil {
		return
	}

	// 前序位置
	m.Depth++
	if root.Left == nil && root.Right == nil { // 到达叶子节点，更新最大深度
		m.Ret = internal.MaxInt(m.Ret, m.Depth)
	}
	m.traverse(root.Left)
	// 中顺位置
	m.traverse(root.Right)
	// 后序位置
	m.Depth--
}

func main() {
	nodeList := []interface{}{3, 9, 20, -1, -1, 15, 7}
	newTree := internal.BuildBinTree(nodeList, 0, len(nodeList))
	res := maxDepth(newTree)
	fmt.Println(res)
	resV2 := maxDepthV2(newTree)
	fmt.Println(resV2)
}
