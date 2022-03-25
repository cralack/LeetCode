package minimumpathsum

import "testing"

func minPathSum(grid [][]int) int {
	min := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res > v {
				res = v
			}
		}
		return res
	}
	n, m := len(grid), len(grid[0])
	dp := make([]int, m+1)
	for i, n := range grid[0] {
		dp[i+1] = n + dp[i]
	}
	for i := 1; i < n; i++ {
		for j, n := range grid[i] {
			if j == 0 {
				dp[j+1] = dp[j+1] + n
			} else {
				dp[j+1] = min(dp[j], dp[j+1]) + n
			}
		}
	}
	return dp[m]
}
func Test_minimum_path_sum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	t.Log(minPathSum(grid))
}
