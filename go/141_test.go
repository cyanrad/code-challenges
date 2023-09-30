package leet

import "testing"

func hasCycle(head *ListNode) bool {
	exists := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := exists[head]; !ok {
			exists[head] = struct{}{}
			head = head.Next
		} else {
			return true
		}
	}

	return false
}

func hasCycle2(head *ListNode) bool {
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	for head != nil {
		if tail.Next == head {
			return true
		}
		head = head.Next
	}

	return false
}

func TestHasCycle(t *testing.T) {
	hasCycle2(creatNodeList(0, 5))
}
