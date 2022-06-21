package painthousemid

import (
	"testing"
)

func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][3]int, n)
	for i := 0; i < 3; i++ {
		dp[0][i] = costs[0][i]
	}
	for i := 1; i < n; i++ {
		// dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		// dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		// dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
		for j := 0; j < 3; j++ {
			dp[i][j] = min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3]) + costs[i][j]
		}
	}
	return min(min(dp[n-1][0], dp[n-1][1]), dp[n-1][2])
}

func minCost_space_comp(costs [][]int) int {
	dp := costs[0]
	for _, cost := range costs[1:] {
		dpNew := make([]int, 3)
		for j, c := range cost {
			dpNew[j] = min(dp[(j+1)%3], dp[(j+2)%3]) + c
		}
		dp = dpNew
	}
	return min(min(dp[0], dp[1]), dp[2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_paint_house(t *testing.T) {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	t.Log(minCost(costs))
}
