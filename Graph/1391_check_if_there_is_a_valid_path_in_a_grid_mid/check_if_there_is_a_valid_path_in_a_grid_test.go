package check_if_there_is_a_valid_path_in_a_grid_mid

import (
	"testing"
)

var dirs = [7][2][2]int{
	{},
	{{0, -1}, {0, 1}},
	{{-1, 0}, {1, 0}},
	{{0, -1}, {1, 0}},
	{{1, 0}, {0, 1}},
	{{0, -1}, {-1, 0}},
	{{0, 1}, {-1, 0}},
}

// contains 判断街道 street 是否包含移动方向 dir
func contains(street int, dir [2]int) bool {
	return dirs[street][0] == dir || dirs[street][1] == dir
}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if x == m-1 && y == n-1 {
			return true
		}
		vis[x][y] = true
		for _, dir := range dirs[grid[x][y]] {
			i, j := x+dir[0], y+dir[1]
			// 核心判断逻辑：
			// 1. 0 <= i < m && 0 <= j < n: 确保下一个格子没有越界
			// 2. !vis[i][j]: 确保下一个格子没有被走过
			// 3. contains(...): 确保下一个格子有一个**反向的开口**来迎接我们。
			// if 0 <= i && i < m && 0 <= j && j < n && !vis[i][j] &&
			// 	contains(grid[i][j], [2]int{-dir[0], -dir[1]}) && dfs(i, j) {
			// 	return true
			// }

			if i < 0 || i >= m || j < 0 || j >= n || vis[i][j] {
				continue
			}
			if !contains(grid[i][j], [2]int{-dir[0], -dir[1]}) {
				continue
			}
			if dfs(i, j) {
				return true
			}
		}
		return false
	}
	return dfs(0, 0)
}

func Test_check_if_there_is_a_valid_path_in_a_grid(t *testing.T) {
	tests := []struct {
		grid [][]int
	}{
		{grid: [][]int{{2, 4, 3}, {6, 5, 2}}},
		{grid: [][]int{{1, 2, 1}, {1, 2, 1}}},
		{grid: [][]int{{1, 1, 1, 1, 1, 1, 3}}},
		{grid: [][]int{{2}, {2}, {2}, {2}, {2}, {2}, {6}}},
	}
	for _, tt := range tests {
		t.Log(hasValidPath(tt.grid))
	}
}
