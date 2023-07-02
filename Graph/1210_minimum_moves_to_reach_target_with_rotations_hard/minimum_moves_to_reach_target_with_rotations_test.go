package minimummovestoreachtarget

import "testing"

type tuple struct{ x, y, s int }

var dirs = []tuple{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

func minimumMoves(g [][]int) int {
	n := len(g)
	vis := make([][][2]bool, n)
	for i := range vis {
		vis[i] = make([][2]bool, n)
	}
	vis[0][0][0] = true // 初始位置
	que := []tuple{{}}
	for step := 1; len(que) > 0; step++ {
		tmp := que
		que = nil
		for _, cur := range tmp {
			for _, d := range dirs {
				// next Node
				x, y, s := cur.x+d.x, cur.y+d.y, cur.s^d.s
				// 蛇头坐标
				x2, y2 := x+s, y+(s^1)
				// 蛇头在graph内且vis[蛇尾]==false
				if x2 < n && y2 < n && !vis[x][y][s] &&
					// 蛇身不能在障碍物上
					g[x][y] == 0 && g[x2][y2] == 0 &&
					// 无需旋转 || 旋转所需空间[x+1][y+1]无障碍
					(d.s == 0 || g[x+1][y+1] == 0) {
					// 终点特判
					if x == n-1 && y == n-2 { // 此时蛇头一定在 (n-1,n-1)
						return step
					}
					vis[x][y][s] = true
					que = append(que, tuple{x, y, s})
				}
			}
		}
	}
	return -1
}

func Test_minimum_moves_to_reach_target(t *testing.T) {
	grid := [][]int{{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 0}}
	t.Log(minimumMoves(grid))
	grid = [][]int{{0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1},
		{1, 1, 1, 0, 0, 1},
		{1, 1, 1, 0, 0, 1},
		{1, 1, 1, 0, 0, 0}}
	t.Log(minimumMoves(grid))
	grid = [][]int{{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 0}}
	t.Log(minimumMoves(grid))
	grid = [][]int{{0, 0, 0},
		{0, 0, 1},
		{0, 0, 0}}
	t.Log(minimumMoves(grid))
	grid = [][]int{{0, 0, 0},
		{0, 0, 0},
		{0, 1, 0}}
	t.Log(minimumMoves(grid))
}
