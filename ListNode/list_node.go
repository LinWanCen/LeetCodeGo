package ListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(input []int) (listNode *ListNode) {
	if input == nil || len(input) == 0 {
		return nil
	}
	listNode = &ListNode{Val: input[0]}
	var curr = listNode
	for i := 1; i < len(input); i++ {
		curr.Next = &ListNode{Val: input[i]}
		curr = curr.Next
	}
	return
}

func ToArr(listNode *ListNode) (arr []int) {
	for listNode != nil {
		arr = append(arr, listNode.Val)
		listNode = listNode.Next
	}
	return
}
