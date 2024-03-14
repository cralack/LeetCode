package findvalidmatrixgivenrowandcolumnsums

import "testing"

func restoreMatrix(rowSum []int, colSum []int) (ans [][]int) {
	ans = make([][]int, len(rowSum))

	for i, rs := range rowSum {
		ans[i] = make([]int, len(colSum))
		for j, cs := range colSum {
			ans[i][j] = min(rs, cs)
			rs -= ans[i][j]
			colSum[j] -= ans[i][j]
		}
	}

	return
}
func restoreMatrix_v2(rowSum []int, colSum []int) (ans [][]int) {
	m, n := len(rowSum), len(colSum)
	ans = make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i, j := 0, 0; i < m && j < n; {
		rs, cs := rowSum[i], colSum[j]
		if rs < cs {
			ans[i][j] = rs // 去掉第 i 行
			colSum[j] -= rs
			i++ // 往下走
		} else {
			ans[i][j] = cs // 去掉第 j 列
			rowSum[i] -= cs
			j++ // 往右走
		}
	}
	return
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_find_valid_matrix_given_row_and_column_sums(t *testing.T) {
	tests := []struct {
		rowSum, colSum []int
	}{
		{[]int{3, 8}, []int{4, 7}},
		{[]int{5, 7, 10}, []int{8, 6, 8}},
		{[]int{14, 9}, []int{6, 9, 8}},
		{[]int{1, 0}, []int{1}},
		{[]int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Log(restoreMatrix(append([]int(nil), tt.rowSum...),
				append([]int(nil), tt.colSum...)))
			t.Log(restoreMatrix_v2(append([]int(nil), tt.rowSum...),
				append([]int(nil), tt.colSum...)))

		})
	}
}
func Benchmark_find(b *testing.B) {
	tests := []struct {
		rowSum, colSum []int
	}{
		{[]int{3, 8}, []int{4, 7}},
		{[]int{5, 7, 10}, []int{8, 6, 8}},
		{[]int{14, 9}, []int{6, 9, 8}},
		{[]int{1, 0}, []int{1}},
		{[]int{0}, []int{0}},
	}

	b.Run("V1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, tt := range tests {
				restoreMatrix(append([]int(nil), tt.rowSum...),
					append([]int(nil), tt.colSum...))
			}
		}
		b.StopTimer()
	})
	b.Run("V2", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, tt := range tests {
				restoreMatrix_v2(append([]int(nil), tt.rowSum...),
					append([]int(nil), tt.colSum...))
			}
		}
		b.StopTimer()
	})
}
