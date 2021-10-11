package ListNode

import (
	"testing"

	. "github.com/LinWanCen/LeetCodeGo/assert"
)

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	// nil 1 2 3 4 5 nil
	//   a b t ->  a b
	var a *ListNode = nil
	for head != nil {
		t := head.Next
		head.Next = a
		a, head = head, t
	}
	return a
}

func TestReverseList(t *testing.T) {
	AssertEqualArr(t,
		ToArr(reverseList(NewListNode([]int{1, 2, 3, 4, 5}))),
		[]int{5, 4, 3, 2, 1},
	)

	AssertEqualArr(t,
		ToArr(reverseList(NewListNode([]int{1, 2}))),
		[]int{2, 1},
	)

	AssertEqualArr(t,
		ToArr(reverseList(NewListNode([]int{5}))),
		[]int{5},
	)

	AssertEqualArr(t,
		ToArr(reverseList(NewListNode([]int{}))),
		[]int{},
	)
}
