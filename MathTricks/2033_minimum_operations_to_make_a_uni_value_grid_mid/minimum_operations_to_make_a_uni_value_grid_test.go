package minimum_operations_to_make_a_uni_value_grid_mid

import (
	"slices"
	"testing"
)

func minOperations(grid [][]int, x int) int {
	k := len(grid) * len(grid[0])
	a := make([]int, 0, k)
	target := grid[0][0] % x

	// 1. 判断是否无解
	for _, row := range grid {
		for _, v := range row {
			// 每个数模 x 都必须相等
			if v%x != target {
				return -1
			}
		}
		a = append(a, row...)
	}

	// 2. 计算 grid 的中位数 median
	slices.Sort(a)
	median := a[k/2]

	// 3. 计算操作次数
	ans := 0
	for _, v := range a {
		ans += abs(v - median)
	}
	return ans / x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Test_minimum_operations_to_make_a_uni_value_grid(t *testing.T) {
	tests := []struct {
		grid [][]int
		x    int
	}{
		{grid: [][]int{{2, 4}, {6, 8}}, x: 2},
		{grid: [][]int{{1, 5}, {2, 3}}, x: 1},
		{grid: [][]int{{1, 2}, {3, 4}}, x: 2},
	}
	for _, tt := range tests {
		t.Log(minOperations(tt.grid, tt.x))
	}
}
