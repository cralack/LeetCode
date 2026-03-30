package champagnetower

import (
	"testing"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	// 定义dp[i][j]为第i行第j列杯子所经过的水的流量（而不是最终剩余的水量）
	dp := make([][]float64, query_row+3)
	for i := range dp {
		dp[i] = make([]float64, i+1)
	}
	dp[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][j] > 1 {
				dp[i+1][j] += (dp[i][j] - 1) / 2
				dp[i+1][j+1] += (dp[i][j] - 1) / 2
			}
		}
	}
	return min(1.0, dp[query_row][query_glass])
}

func Test_champagne_tower(t *testing.T) {
	tests := []struct {
		poured      int
		query_glass int
		query_row   int
	}{
		{poured: 1, query_glass: 1, query_row: 1},
		{poured: 2, query_glass: 1, query_row: 1},
		{poured: 6, query_glass: 2, query_row: 3},
		{poured: 100000009, query_glass: 17, query_row: 33},
	}
	for _, test := range tests {
		t.Log(champagneTower(test.poured, test.query_row, test.query_glass))
	}
}
