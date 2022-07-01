package numberofenclaves

import (
	"testing"
)

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(grid, 0, j)
		dfs(grid, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if m <= i || i < 0 || n <= j || j < 0 {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
	}

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
func TestNumberOfEnclaves(t *testing.T) {
	grid1 := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0}}
	grid2 := [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0}}

	res1 := numEnclaves(grid1)
	res2 := numEnclaves(grid2)

	t.Log(res1, res2)
}
