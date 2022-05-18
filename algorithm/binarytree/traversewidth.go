package main

import (
	"code_tpl_go/algorithm/binarytree/internal"
	"fmt"
)

/*
时间复杂度：O(n)
空间复杂度：O(height)，其中 HeightHeight 为二叉树的高度。由于递归函数在递归过程中需要为每一层递归函数分配栈空间，所以这里需要额外的空间且该空间取决于递归的深度，而递归的深度显然为二叉树的高度，并且每次递归调用的函数里又只用了常数个变量，所以所需空间复杂度为 O(Height)O(Height)
*/
func diameterOfBinaryTree(root *internal.TreeNode) int {
	ans := 0
	depth(root, &ans)
	return ans - 1
}

func depth(root *internal.TreeNode, ans *int) int {
	if root == nil {
		return 0 // 访问到空节点了，返回0
	}
	leftDepth := depth(root.Left, ans)                   // 左儿子为根的子树的深度
	rightDepth := depth(root.Right, ans)                 // 右儿子为根的子树的深度
	*ans = internal.MaxInt(*ans, leftDepth+rightDepth+1) // 计算d_node即L+R+1 并更新ans
	fmt.Printf("root[%v]: %d\t%d\t%d\n", root.Data, *ans, leftDepth, rightDepth)
	return internal.MaxInt(leftDepth, rightDepth) + 1 // 返回该节点为根的子树的深度
}

func main() {
	nodeList := []interface{}{3, 9, 20, -1, -1, 15, 7}
	newTree := internal.BuildBinTree(nodeList, 0, len(nodeList))
	res := diameterOfBinaryTree(newTree)
	fmt.Println(res)
}
