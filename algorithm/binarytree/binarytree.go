package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math/rand"
)

// 二叉树结构定义
type node struct {
	data  int   // 数据域
	left  *node // 指向左子树的根结点
	right *node // 指向右子树的根结点
}

func NewNode(data int) *node {
	return &node{data: data}
}

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

func (nd *node) Insert(newNode *node) {
	//1
	if newNode.data == nd.data {
		return
	}
	//2
	if newNode.data > nd.data {
		if nd.right == nil {
			nd.right = newNode
		} else {
			nd.right.Insert(newNode)
		}
	} else { //3 小于 继续比较插入到 左孩子
		if nd.left == nil {
			nd.left = newNode
		} else {
			nd.left.Insert(newNode)
		}
	}
}

func main() {
	root := NewNode(250)
	for i := 0; i < 14; i++ {
		n := rand.Intn(500)
		fmt.Println("i=", i, "的随机数是", n)
		root.Insert(NewNode(n))
	}
	spew.Dump(root)
}
