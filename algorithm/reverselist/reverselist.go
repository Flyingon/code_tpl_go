package main

import (
	"code_tpl_go/algorithm/reverselist/internal"
)

func reverseList(head *internal.ListNode) *internal.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *internal.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseListV2(head *internal.ListNode) *internal.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head

	current := head.Next
	for current != nil {
		prev, current, current.Next = current, current.Next, prev
	}

	head.Next = nil

	return prev
}

func main() {
	head := internal.GenNodeListByIncr(10)
	head.Print()
	newHead := reverseListV2(head)
	newHead.Print()
}
