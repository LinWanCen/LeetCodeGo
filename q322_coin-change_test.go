package assert

import (
	"math"
	"testing"

	. "github.com/LinWanCen/LeetCodeGo/assert"
)

// 322. 零钱兑换
func coinChange(coins []int, amount int) int {
	// 数组⼤⼩为 amount + 1，初始值也为 amount + 1
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	// base case
	dp[0] = 0
	// 循环在求所有选择的最⼩值
	for _, coin := range coins {
		// 循环在遍历所有状态的所有取值
		for i := 0; i <= amount; i++ {
			// ⼦问题⽆解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func TestCoinChange(t *testing.T) {
	AssertEqual(t, coinChange([]int{1, 2, 5}, 11), 3)
	AssertEqual(t, coinChange([]int{2}, 3), -1)
	AssertEqual(t, coinChange([]int{1}, 0), 0)
	AssertEqual(t, coinChange([]int{1}, 1), 1)
	AssertEqual(t, coinChange([]int{1}, 2), 2)
}
