package triangle_mid

import (
	"testing"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	dp[n-1] = triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j, x := range triangle[i] {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + x
		}
	}
	return dp[0][0]
}

func Test_triangle(t *testing.T) {
	tests := []struct {
		triangle [][]int
	}{
		{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
		{triangle: [][]int{{-10}}},
	}
	for _, tt := range tests {
		t.Log(minimumTotal(tt.triangle))
	}
}
