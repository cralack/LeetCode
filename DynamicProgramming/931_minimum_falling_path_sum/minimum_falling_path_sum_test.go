package minimumfallingpathsum

import (
	"math"
	"testing"
)

func minFallingPathSum(matrix [][]int) (ans int) {
	// init
	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+2)
		dp[i][0], dp[i][n+1] = math.MaxInt, math.MaxInt
	}
	copy(dp[0][1:], matrix[0])
	for row := 1; row < n; row++ {
		for col, x := range matrix[row] {
			dp[row][col+1] = min(min(dp[row-1][col], dp[row-1][col+1]), dp[row-1][col+2]) + x
		}
	}
	ans = math.MaxInt
	for _, x := range dp[n-1] {
		ans = min(ans, x)
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_minimum_falling_path_sum(t *testing.T) {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	t.Log(minFallingPathSum(matrix))
	matrix = [][]int{{-19, 57}, {-40, -5}}
	t.Log(minFallingPathSum(matrix))
	matrix = [][]int{{-48}}
	t.Log(minFallingPathSum(matrix))
}
