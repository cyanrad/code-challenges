package leet

import (
	"log"
	"testing"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := list1
	for list2 != nil && list1 != nil {
		log.Println(list2.Val, list1.Val)
		if list2.Val >= list1.Val {
			list2 = insertNode(list1, list2)
			list1 = list1.Next
		} else {
			list1 = list1.Next
		}
	}

	return head
}

func TestMergeTwoLists(t *testing.T) {
	l1 := creatNodeList(1, 5)
	l2 := creatNodeList(3, 7)
	l2 = mergeTwoLists(l1, l2)
	printNode(l2)
}
