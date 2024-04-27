package find_the_width_of_columns_of_a_grid

import (
	"testing"
)

func findColumnWidth(grid [][]int) (ans []int) {
	ans = make([]int, len(grid[0]))
	for j := range grid[0] {
		mn, mx := grid[0][j], grid[0][j]
		for _, row := range grid {
			mn = min(mn, row[j])
			mx = max(mx, row[j])
		}
		xLen := 1
		for x := max(mx/10, -mn); x > 0; x /= 10 {
			xLen++
		}
		ans[j] = xLen
	}
	return
}

func Test_find_the_width_of_columns_of_a_grid(t *testing.T) {
	grid := [][]int{{1}, {22}, {333}}
	t.Log(findColumnWidth(grid))
	grid = [][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}}
	t.Log(findColumnWidth(grid))
}
