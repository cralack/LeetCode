package liwudezuidajiezhilcof

import "testing"

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = max(dp[i+1][j+1],
				max(dp[i+1][j]+grid[i][j], dp[i][j+1]+grid[i][j]))
		}
	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_li_wu_de_zui_da_jie_zhi_lcof(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	t.Log(maxValue(grid))
}
