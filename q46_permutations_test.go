package assert

import (
	"testing"

	. "github.com/LinWanCen/LeetCodeGo/assert"
)

// 46. 全排列
func permute(nums []int) [][]int {
	n := len(nums)
	length := 1
	for i := 1; i <= n; i++ {
		length *= i
	}
	res := make([][]int, 0, length)
	arr := make([]int, 0, n)
	b := make([]bool, n)
	backtrack(&nums, &b, &arr, &res)
	return res
}

func backtrack(nums *[]int, b *[]bool, arr *[]int, res *[][]int) {
	if len(*nums) == len(*arr) {
		*res = append(*res, append([]int(nil), *arr...))
		return
	}
	for i := 0; i < len(*nums); i++ {
		if (*b)[i] {
			continue
		}
		*arr = append(*arr, (*nums)[i])
		(*b)[i] = true
		backtrack(nums, b, arr, res)
		(*b)[i] = false
		*arr = (*arr)[:(len(*arr) - 1)]
	}
}

func TestPermute(t *testing.T) {
	AssertEqualMatrix(t, permute([]int{1, 2, 3}), [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
	})
	AssertEqualMatrix(t, permute([]int{0, 1}), [][]int{
		{0, 1}, {1, 0},
	})
	AssertEqualMatrix(t, permute([]int{1}), [][]int{
		{1},
	})
}
