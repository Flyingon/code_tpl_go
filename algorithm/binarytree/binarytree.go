package main

import (
	"math/rand"
)

// 二叉树结构定义
type node struct {
	data  int   // 数据域
	left  *node // 指向左子树的根结点
	right *node // 指向右子树的根结点
}

// NewNode 新建二叉树node节点
func NewNode(data int) *node {
	return &node{data: data}
}

// Insert 数据插入
func (nd *node) Insert(newNode *node) {
	// 1. 这个数据已经在该节点存在
	if newNode.data == nd.data {
		return
	}
	// 2. 新数据大于该节点，插入右子树
	if newNode.data > nd.data {
		if nd.right == nil {
			nd.right = newNode // 右孩子节点是空的，直接放进去
		} else {
			nd.right.Insert(newNode) // 右孩子节点已经有数据了，递归插入
		}
	} else { //3. 新数据大于该节点，插入左子树
		if nd.left == nil {
			nd.left = newNode // 左孩子节点是空的，直接放进去
		} else {
			nd.left.Insert(newNode) // 左孩子节点已经有数据了，递归插入
		}
	}
}

// Search 数据搜索
func (nd *node) Search(dt int) *node {
	if nd == nil {
		return nil
	}
	//1
	if dt == nd.data {
		return nd
	}

	//2 大于当前节点，递归右边
	if dt > nd.data {
		return nd.right.Search(dt)
	}
	//2 大于当前节点，递归左边
	if dt < nd.data {
		return nd.left.Search(dt)
	}
	return nil
}

func main() {
	root := NewNode(250)
	for i := 0; i < 14; i++ {
		n := rand.Intn(500)
		//fmt.Println("i=", i, "的随机数是", n)
		root.Insert(NewNode(n))
	}
	//spew.Dump(root)
}
