package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseListV2(head *ListNode) *ListNode {
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

func printList(head *ListNode) {
	var intList []int
	pPos := head
	for pPos.Next != nil {
		intList = append(intList, pPos.Val)
		pPos = pPos.Next
	}
	fmt.Println(intList)
}

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	pPos := head
	for i := 1; i < 10; i++ {
		new := &ListNode{
			Val:  i,
			Next: nil,
		}
		pPos.Next = new
		pPos = new
	}
	printList(head)
	newHead := reverseList(head)
	printList(newHead)
}
