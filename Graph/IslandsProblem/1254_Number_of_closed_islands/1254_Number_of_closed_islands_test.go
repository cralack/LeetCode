package numberofclosedislands

import "testing"

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// flodfill up/down side
	for j := 0; j < n; j++ {
		dfs(0, j, grid)
		dfs(m-1, j, grid)
	}
	// flodfill left/right side
	for i := 0; i < m; i++ {
		dfs(i, 0, grid)
		dfs(i, n-1, grid)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
			}
			dfs(i, j, grid)
		}
	}
	return res
}

func dfs(i, j int, grid [][]int) {
	m, n := len(grid), len(grid[0])
	if m <= i || i < 0 || n <= j || j < 0 {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs(i+1, j, grid)
	dfs(i-1, j, grid)
	dfs(i, j+1, grid)
	dfs(i, j-1, grid)
}

func TestNumberOfClosedIslands(t *testing.T) {
	grid1 := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0}}
	grid2 := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0}}
	grid3 := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1}}

	res1 := closedIsland(grid1)
	res2 := closedIsland(grid2)
	res3 := closedIsland(grid3)

	t.Log(res1, res2, res3)

}
