package champagnetower

import (
	"testing"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	//定义dp[i][j]为第i行第j列杯子所经过的水的流量（而不是最终剩余的水量）
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

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Test_champagne_tower(t *testing.T) {
	poured, query_glass, query_row := 1, 1, 1
	t.Log(champagneTower(poured, query_row, query_glass))
	poured, query_glass, query_row = 2, 1, 1
	t.Log(champagneTower(poured, query_row, query_glass))
	poured, query_glass, query_row = 6, 2, 3
	t.Log(champagneTower(poured, query_row, query_glass))
}
