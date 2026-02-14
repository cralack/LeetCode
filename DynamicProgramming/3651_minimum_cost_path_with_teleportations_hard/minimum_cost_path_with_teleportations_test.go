package minimum_cost_path_with_teleportations_hard

import (
	"math"
	"slices"
	"testing"
)

func minCost(grid [][]int, k int) int {
	n := len(grid[0])
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}
	sufMinF := make([]int, mx+2)
	for i := range sufMinF {
		sufMinF[i] = math.MaxInt
	}
	minF := make([]int, mx+1)
	f := make([]int, n+1)

	for range k + 1 {
		for i := range minF {
			minF[i] = math.MaxInt
		}
		for i := range f {
			f[i] = math.MaxInt / 2
		}
		f[1] = -grid[0][0]
		for _, row := range grid {
			for j, x := range row {
				// 1. f[j] + x: 从左边过来
				// 2. f[j+1] + x: 从上边过来
				// 3. sufMinF[x]: 从上一轮 >= x 的值跳过来
				f[j+1] = min(f[j]+x, f[j+1]+x, sufMinF[x])
				minF[x] = min(minF[x], f[j+1])
			}
		}

		for i := mx; i >= 0; i-- {
			sufMinF[i] = min(sufMinF[i+1], minF[i])
		}
	}
	return f[n]
}

func Test_minimum_cost_path_with_teleportations(t *testing.T) {
	tests := []struct {
		grid [][]int
		k    int
	}{
		{grid: [][]int{{1, 3, 3}, {2, 5, 4}, {4, 3, 5}}, k: 2},
		{grid: [][]int{{1, 2}, {2, 3}, {3, 4}}, k: 1},
	}
	for _, tt := range tests {
		t.Log(minCost(tt.grid, tt.k))
	}
}
