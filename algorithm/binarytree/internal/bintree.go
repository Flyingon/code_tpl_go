package internal

// TreeNode 树结构定义
type TreeNode struct {
	Data  interface{} // 数据域
	Left  *TreeNode   // 指向左子树的根结点
	Right *TreeNode   // 指向右子树的根结点
}

// NewTreeNode 构造树节点
func NewTreeNode(data interface{}) *TreeNode {
	return &TreeNode{Data: data}
}

// BuildBinTree 构造二叉树
func BuildBinTree(treeList []interface{}, i, n int) *TreeNode {
	if i >= n {
		return nil
	}
	//fmt.Printf("%d: %v\n", i, treeList[i])
	if treeList[i] == nil {
		return nil
	}
	root := NewTreeNode(treeList[i])
	root.Left = BuildBinTree(treeList, 2*i+1, n)
	root.Right = BuildBinTree(treeList, 2*i+2, n)
	return root
}
