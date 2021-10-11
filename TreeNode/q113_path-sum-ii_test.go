package TreeNode

import (
	"testing"

	. "github.com/LinWanCen/LeetCodeGo/assert"
)

// 113. 路径总和 II （欢聚时代）
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Right, left)
		dfs(node.Left, left)
	}
	dfs(root, targetSum)
	return
}

func TestPathSum(t *testing.T) {
	root := NewTreeNode([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1})
	AssertEqualMatrix(t, pathSum(
		root, 22),
		[][]int{
			{5, 4, 11, 2},
			{5, 8, 4, 5},
		})
	root2 := NewTreeNode([]int{1, 2, 3})
	AssertEqualMatrix(t, pathSum(
		root2, 3),
		[][]int{
			{1, 2},
		})
	AssertEqualMatrix(t, pathSum(
		root2, 5),
		[][]int{})
	AssertEqualMatrix(t, pathSum(
		root2, 0),
		[][]int{})
}
