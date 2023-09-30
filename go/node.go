package leet

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertNode(n1, newN *ListNode) *ListNode {
	ret := newN.Next
	newN.Next = n1.Next
	n1.Next = newN
	return ret
}

func printNode(n *ListNode) {
	current := n
	for current != nil {
		log.Println(current.Val)
		current = current.Next
	}
}

func creatNodeList(first, last int) *ListNode {
	head := ListNode{first, nil}
	current := &head
	for i := first + 1; i < last; i++ {
		current.Next = &ListNode{i, nil}
		current = current.Next
	}
	return &head
}
