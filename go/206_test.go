package leet

import (
	"testing"
)

func reverseList(head *ListNode) *ListNode {
	stack := []int{}
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	returnedHead := ListNode{
		Val:  stack[len(stack)-1],
		Next: nil,
	}
	prevHead := &returnedHead
	for i := len(stack) - 2; i > -1; i-- {
		head = &ListNode{
			Val:  stack[i],
			Next: nil,
		}
		prevHead.Next = head
		prevHead = head
	}

	return &returnedHead
}

func TestReverseList(t *testing.T) {
	current := creatNodeList(1, 5)
	current = reverseList(current)
	printNode(current)
}
