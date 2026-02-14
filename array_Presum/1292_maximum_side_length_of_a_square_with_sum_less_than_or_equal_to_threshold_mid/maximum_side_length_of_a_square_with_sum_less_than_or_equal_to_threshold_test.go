package maximum_side_length_of_a_square_with_sum_less_than_or_equal_to_threshold_mid

import (
	"testing"
)

func maxSideLength(mat [][]int, threshold int) (ans int) {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		sum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			sum[i][j] = mat[i-1][j-1] + sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1]
		}
	}
	// 返回左上角在 (r1, c1)，右下角在 (r2, c2) 的子矩阵元素和
	query := func(r1, c1, r2, c2 int) int {
		return sum[r2+1][c2+1] - sum[r2+1][c1] - sum[r1][c2+1] + sum[r1][c1]
	}

	for i := range m {
		for j := range n {
			for i+ans < m && j+ans < n &&
				query(i, j, i+ans, j+ans) <= threshold {
				ans++
			}
		}
	}
	return
}

func Test_maximum_side_length_of_a_square_with_sum_less_than_or_equal_to_threshold(t *testing.T) {
	tests := []struct {
		mat       [][]int
		threshold int
	}{
		{mat: [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, threshold: 4},
		{mat: [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, threshold: 1},
	}
	for _, tt := range tests {
		t.Log(maxSideLength(tt.mat, tt.threshold))
	}
}
