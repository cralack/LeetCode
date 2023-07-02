package numberofdistinctislands

import (
	"bytes"
	"testing"
)

func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	islands := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				var buf bytes.Buffer
				dfs(grid, i, j, 's', &buf)
				islands[buf.String()] = true
			}
		}
	}
	cnt := len(islands)

	return cnt
}
func dfs(grid [][]int, i, j int, dir byte, buf *bytes.Buffer) {
	m, n := len(grid), len(grid[0])
	if m <= i || i < 0 || n <= j || j < 0 {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	tmp := []byte{}
	tmp = append(tmp, dir, ',')
	buf.WriteString(string(tmp))
	// fmt.Printf("%x\n", unsafe.Pointer(&buf))

	dfs(grid, i+1, j, 'd', buf)
	dfs(grid, i-1, j, 'u', buf)
	dfs(grid, i, j+1, 'r', buf)
	dfs(grid, i, j-1, 'l', buf)

	tmp2 := []byte{}
	tmp2 = append(tmp2, '_', dir, ',')
	buf.WriteString(string(tmp2))
	// fmt.Printf("%x\n", unsafe.Pointer(&buf))
}

func TestNumberOfDistinctIslands(t *testing.T) {
	grid1 := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0}}
	res := numDistinctIslands(grid1)
	t.Log(res)
}
