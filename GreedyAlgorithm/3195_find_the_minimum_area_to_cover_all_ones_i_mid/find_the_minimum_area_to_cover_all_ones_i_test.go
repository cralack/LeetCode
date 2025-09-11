package findtheminimumareatocoverallonesimid

import "testing"

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	minX, minY, maxX, maxY := m, n, 0, 0
	for x, row := range grid {
		for y, g := range row {
			if g == 1 {
				minX = min(minX, x)
				minY = min(minY, y)
				maxX = max(maxX, x)
				maxY = max(maxY, y)
			}
		}
	}
	return (maxX - minX + 1) * (maxY - minY + 1)
}
func Test_find_the_minimum_area_to_cover_all_ones_i(t *testing.T) {
	tests := []struct {
		grid [][]int
	}{
		{grid: [][]int{{0, 1, 0}, {1, 0, 1}}},
		{grid: [][]int{{0, 0}, {1, 0}}},
	}
	for _, tt := range tests {
		t.Log(minimumArea(tt.grid))
	}
}
