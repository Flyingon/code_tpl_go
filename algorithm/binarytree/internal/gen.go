package internal

import (
	"math/rand"
	"time"
)

// Node 二叉树结构定义
type Node struct {
	Data  int   // 数据域
	Left  *Node // 指向左子树的根结点
	Right *Node // 指向右子树的根结点
}

// NewNode 新建二叉树node节点
func NewNode(data int) *Node {
	return &Node{Data: data}
}

// Insert 数据插入
func (nd *Node) Insert(newNode *Node) {
	// 1. 这个数据已经在该节点存在
	if newNode.Data == nd.Data {
		return
	}
	// 2. 新数据大于该节点，插入右子树
	if newNode.Data > nd.Data {
		if nd.Right == nil {
			nd.Right = newNode // 右孩子节点是空的，直接放进去
		} else {
			nd.Right.Insert(newNode) // 右孩子节点已经有数据了，递归插入
		}
	} else { //3. 新数据大于该节点，插入左子树
		if nd.Left == nil {
			nd.Left = newNode // 左孩子节点是空的，直接放进去
		} else {
			nd.Left.Insert(newNode) // 左孩子节点已经有数据了，递归插入
		}
	}
}

// GenOrderBinTreeRandom 随机生成有序二叉树
func GenOrderBinTreeRandom(rootData, times, randRange int) *Node {
	root := NewNode(rootData)
	for i := 0; i < times; i++ {
		n := getRandomInt(1, randRange)
		//fmt.Println("i=", i, "的随机数是", n)
		root.Insert(NewNode(n))
	}
	return root
}

func init() {
	rand.Seed(time.Now().Unix())
}
func getRandomInt(start, end int) int {
	if end < start {
		return 0
	}
	return start + rand.Intn(end-start)
}
