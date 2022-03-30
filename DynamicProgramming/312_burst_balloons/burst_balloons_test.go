package burstballoons

import (
	"testing"
)

func maxCoins(nums []int) int {
	n := len(nums)
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1 //边界
	for i, n := range nums {
		points[i+1] = n
	}
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}
	return dp[0][n+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Test_burst_balloons(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	t.Log(maxCoins(nums))
}
