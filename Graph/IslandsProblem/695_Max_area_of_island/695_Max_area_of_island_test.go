package maxareaofisland

import "testing"

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := dfs(i, j, grid)
			if tmp > maxArea {
				maxArea = tmp
			}
		}
	}
	return maxArea
}

func dfs(i, j int, grid [][]int) int {
	area := 0
	m, n := len(grid), len(grid[0])

	if m <= i || i < 0 || n <= j || j < 0 {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	area++
	area += dfs(i+1, j, grid)
	area += dfs(i-1, j, grid)
	area += dfs(i, j+1, grid)
	area += dfs(i, j-1, grid)
	return area
}

func TestMaxAre(t *testing.T) {
	grid1 := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	grid2 := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}

	res1 := maxAreaOfIsland(grid1)
	res2 := maxAreaOfIsland(grid2)

	t.Log(res1, res2)

}
