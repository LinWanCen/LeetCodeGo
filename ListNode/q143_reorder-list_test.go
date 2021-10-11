package ListNode

import (
	"testing"

	. "github.com/LinWanCen/LeetCodeGo/assert"
)

// 143. 重排链表 （字节跳动）
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	midLeft := mid(head)
	midRight := midLeft.Next
	// spit
	midLeft.Next = nil
	midRight = rev(midRight)
	merge(head, midRight)
}

func mid(head *ListNode) *ListNode {
	mid, curr := head, head.Next
	for curr != nil && curr.Next != nil {
		curr, mid = curr.Next.Next, mid.Next
	}
	return mid
}

func rev(b *ListNode) *ListNode {
	// nil 1 2 3 4 5 nil
	//   a b t ->  a b
	var a *ListNode = nil
	for b != nil {
		t := b.Next
		b.Next = a
		a, b = b, t
	}
	return a
}

func merge(a, b *ListNode) {
	// 0 1 2 3 4 nil nil
	// a m
	// a b m ->  a   b
	// b n
	// 9 8 7 6 5 nil nil
	var m, n *ListNode
	for a != nil && b != nil {
		m, n = a.Next, b.Next
		a.Next, b.Next = b, m
		a, b = m, n
	}
}

func TestReorderList(t *testing.T) {
	listNode1 := NewListNode([]int{1, 2, 3, 4})
	reorderList(listNode1)
	AssertEqualArr(t, ToArr(listNode1), []int{1, 4, 2, 3})

	listNode2 := NewListNode([]int{1, 2, 3, 4, 5})
	reorderList(listNode2)
	AssertEqualArr(t, ToArr(listNode2), []int{1, 5, 2, 4, 3})

	var listNode0 *ListNode
	reorderList(listNode0)
	AssertEqualArr(t, ToArr(listNode0), []int{})

	listNode := NewListNode([]int{})
	reorderList(listNode)
	AssertEqualArr(t, ToArr(listNode), []int{})
}
