package shortestbridge

import "testing"

func shortestBridge(grid [][]int) (ans int) {
	n := len(grid)
	dirs := []int{-1, 0, 1, 0, -1}
	type pair struct{ i, j int }
	que := []pair{}
	// func part
	var dfs func(int, int)
	dfs = func(i, j int) {
		grid[i][j] = 2
		que = append(que, pair{i, j})
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k], j+dirs[k+1]
			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 1 {
				dfs(x, y)
			}
		}
	}
	// find 1st island && 1->2
	for i, x := 0, 1; i < n && x == 1; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				x = 0
				break
			}
		}
	}
	// bfs find 2nd island
	for {
		for i := len(que); i > 0; i-- {
			cur := que[0]
			que = que[1:]
			for k := 0; k < 4; k++ {
				x, y := cur.i+dirs[k], cur.j+dirs[k+1]
				if x >= 0 && x < n && y >= 0 && y < n {
					if grid[x][y] == 1 {
						return
					}
					if grid[x][y] == 0 {
						grid[x][y] = 2
						que = append(que, pair{x, y})
					}
				}
			}
		}
		ans++
	}
}

func Test_shortest_bridge(t *testing.T) {
	grid := [][]int{
		{0, 1},
		{1, 0}}
	t.Log(shortestBridge(grid))

	grid = [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 1}}
	t.Log(shortestBridge(grid))

	grid = [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1}}
	t.Log(shortestBridge(grid))
}
