package countsubislands

import "testing"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid2), len(grid2[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(grid2, i, j)
			}
		}
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				cnt++
				dfs(grid2, i, j)
			}
		}
	}
	return cnt
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if m <= i || i < 0 || n <= j || j < 0 {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func TestCountSubIslands(t *testing.T) {
	grid1 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1}}
	grid2 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0}}
	grid3 := [][]int{
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1}}
	grid4 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1}}

	res := countSubIslands(grid1, grid2)
	t.Log(res)

	res2 := countSubIslands(grid3, grid4)
	t.Log(res2)
}
