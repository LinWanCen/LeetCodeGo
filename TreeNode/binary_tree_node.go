package TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Index int
}

const null = -2147483648
const Null = -2147483648

// NewTreeNode
// 数组生成二叉树
func NewTreeNode(input []int) (root *TreeNode) {
	var beforeLine []*TreeNode
	root = &TreeNode{Val: input[0]}
	beforeLine = append(beforeLine, root)
	column := 0
	var thisLine []*TreeNode

	println(`
@startuml TreeNode
digraph g {
           0 [label="`, root.Val, `"]`)
	length := len(input)
	for i := 1; i < length; i++ {
		if input[i] != null {
			thisNode := &TreeNode{Val: input[i], Index: i}
			thisLine = append(thisLine, thisNode)
			beforeIndex := column / 2
			beforeNode := beforeLine[beforeIndex]
			println(beforeNode.Index, ` -> `, i, ";", i, `[label="`, thisNode.Val, `"]`)
			if column%2 == 0 {
				beforeLine[beforeIndex].Right = thisNode
			} else {
				beforeLine[beforeIndex].Left = thisNode
			}
		}
		column++
		if column == 2*len(beforeLine) {
			beforeLine = thisLine
			thisLine = []*TreeNode{}
			column = 0
		}
	}
	println(`}
@enduml
`)
	return
}
