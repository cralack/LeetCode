package minimumfallingpathsum

import "testing"

func minFallingPathSum(matrix [][]int) int {
	// init
	n := len(matrix)
	dp := make([][]int, n)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := range dp {
		dp[i] = make([]int, n)
		if i == 0 {
			for j, val := range matrix[i] {
				dp[i][j] = val
			}
		}
	}
	minVal := 999999
	// dp
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 { // left side
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == n-1 { // right side
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]),
					dp[i-1][j+1]) + matrix[i][j]
			}
		}
	}
	// check min
	for k := 0; k < n; k++ {
		if minVal > dp[n-1][k] {
			minVal = dp[n-1][k]
		}
	}

	return minVal
}
func Test_minimum_falling_path_sum(t *testing.T) {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	t.Log(minFallingPathSum(matrix))
	matrix = [][]int{{-19, 57}, {-40, -5}}
	t.Log(minFallingPathSum(matrix))
	matrix = [][]int{{-48}}
	t.Log(minFallingPathSum(matrix))
}
