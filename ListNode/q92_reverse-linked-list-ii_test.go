package ListNode

import (
	"testing"

	. "github.com/LinWanCen/LeetCodeGo/assert"
)

// 92. 反转链表 II （字节跳动）
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	headLeft := &ListNode{Val: -1, Next: head}
	pre := headLeft
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	//   1 2 3 4 5
	// pre
	//     a t
	// pre t a
	//       a t
	// pre t   a
	a := pre.Next
	for i := left; i < right; i++ {
		t := a.Next
		a.Next, t.Next, pre.Next = t.Next, pre.Next, t
	}
	return headLeft.Next
}

func TestReverseBetween(t *testing.T) {
	AssertEqualArr(t,
		ToArr(reverseBetween(NewListNode([]int{1, 2, 3, 4, 5}), 2, 4)),
		[]int{1, 4, 3, 2, 5},
	)

	AssertEqualArr(t,
		ToArr(reverseBetween(NewListNode([]int{3, 5}), 1, 2)),
		[]int{5, 3},
	)

	AssertEqualArr(t,
		ToArr(reverseBetween(NewListNode([]int{5}), 1, 1)),
		[]int{5},
	)
}
