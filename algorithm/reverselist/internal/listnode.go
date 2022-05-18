package internal

import "fmt"

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// GenNodeListByIncr 生成链表，默认数据
func GenNodeListByIncr(length int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	pPos := head
	for i := 1; i < length; i++ {
		node := &ListNode{
			Val:  i,
			Next: nil,
		}
		pPos.Next = node
		pPos = node
	}
	return head
}

// Print 打印链表元素
func (head *ListNode) Print() {
	pPos := head
	for pPos != nil {
		fmt.Printf("%v ", pPos.Val)
		pPos = pPos.Next
	}
	fmt.Printf("\n")
}
