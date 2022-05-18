package main

import (
	"code_tpl_go/algorithm/reverselist/internal"
	"fmt"
)

/* 递归遍历单链表，倒序打印链表元素 */
func traverse(head *internal.ListNode) {
	if head == nil {
		return
	}

	traverse(head.Next)
	// 后序位置
	fmt.Printf("%v ", head.Val)
}
func main() {
	head := internal.GenNodeListByIncr(15)
	head.Print()
	traverse(head)
}
