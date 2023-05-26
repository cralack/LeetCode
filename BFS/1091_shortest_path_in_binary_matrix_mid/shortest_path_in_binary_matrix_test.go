package shortestpathinbinarymatrix

import "testing"

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	que := [][]int{{0, 0}}
	n := len(grid)
	grid[0][0] = 1

	for step := 1; len(que) > 0; step++ {
		// 或size:=len(que)后 i:=0;i<size;i++
		for i := len(que); i > 0; i-- {
			cur := que[0]
			que = que[1:]
			if cur[0] == n-1 && cur[1] == n-1 {
				return step
			}
			for j := 0; j < len(dir); j++ {
				nextX := cur[0] + dir[j][0]
				nextY := cur[1] + dir[j][1]
				if 0 <= nextX && nextX < n && 0 <= nextY && nextY < n &&
					grid[nextX][nextY] != 1 {
					grid[nextX][nextY] = 1
					que = append(que, []int{nextX, nextY})
				}
			}
		}
	}
	return -1
}

func Test_shortest_path_in_binary_matrix(t *testing.T) {
	tests := []struct {
		grid  [][]int
		wanna int
	}{
		{grid: [][]int{{0, 1}, {1, 0}}, wanna: 2},
		{grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, wanna: 4},
		{grid: [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, wanna: -1},
		{grid: [][]int{{0, 0, 0}, {1, 0, 0}, {1, 1, 0}}, wanna: 3},
		{grid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, wanna: 2},
		{grid: [][]int{
			{0, 1, 1, 0, 0, 0},
			{0, 1, 0, 1, 1, 0},
			{0, 1, 1, 0, 1, 0},
			{0, 0, 0, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0}},
			wanna: 14},
	}
	for _, tt := range tests {
		flag := []string{"√", "x"}
		if shortestPathBinaryMatrix(tt.grid) == tt.wanna {
			t.Log(flag[0])
		} else {
			t.Log(flag[1])
		}
	}
}
