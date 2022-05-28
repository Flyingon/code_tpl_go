package main

/*
前序遍历：根结点 ---> 左子树 ---> 右子树

中序遍历：左子树---> 根结点 ---> 右子树

后序遍历：左子树 ---> 右子树 ---> 根结点
*/

import (
	"code_tpl_go/algorithm/binarytree/internal"
	"fmt"
)

func preorderTraversal(root *internal.TreeNode) []interface{} {
	var res []interface{}
	_preTraversal(root, &res)
	return res
}

func _preTraversal(root *internal.TreeNode, res *[]interface{}) {
	if root == nil {
		return
	}
	*res = append(*res, root.Data)
	_preTraversal(root.Left, res)
	_preTraversal(root.Right, res)
	return
}

func midOrderTraversal(root *internal.TreeNode) []interface{} {
	var res []interface{}
	_midTraversal(root, &res)
	return res
}

func _midTraversal(root *internal.TreeNode, res *[]interface{}) {
	if root == nil {
		return
	}
	//if root.Left != nil {
	//	_midTraversal(root.Left, res)
	//}
	//*res = append(*res, root.Data)
	//_midTraversal(root.Right, res)

	_midTraversal(root.Left, res)
	*res = append(*res, root.Data)
	_midTraversal(root.Right, res)
	return
}

func afterOrderTraversal(root *internal.TreeNode) []interface{} {
	var res []interface{}
	_afterTraversal(root, &res)
	return res
}

func _afterTraversal(root *internal.TreeNode, res *[]interface{}) {
	if root == nil {
		return
	}
	//if root.Left != nil {
	//	_afterTraversal(root.Left, res)
	//}
	//if root.Right != nil {
	//	_afterTraversal(root.Right, res)
	//}
	//*res = append(*res, root.Data)

	_afterTraversal(root.Left, res)
	_afterTraversal(root.Right, res)
	*res = append(*res, root.Data)
	return
}

func main() {
	nodeList := []interface{}{1, 2, 3, 4, 5, nil, 6, nil, nil, 7, 8}
	newTree := internal.BuildBinTree(nodeList, 0, len(nodeList))

	resPre := preorderTraversal(newTree)
	fmt.Println("前序遍历:", resPre)

	resMid := midOrderTraversal(newTree)
	fmt.Println("中序遍历:", resMid)

	resAfter := afterOrderTraversal(newTree)
	fmt.Println("后续遍历:", resAfter)
}
