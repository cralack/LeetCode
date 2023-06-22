package pondsizeslcci

import (
	"sort"
	"testing"
)

var dir_8 = [8][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1},
	{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
var area int

func pondSizes(land [][]int) (ans []int) {
	n, m := len(land), len(land[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		land[x][y] = 1
		area++
		for i := 0; i < 8; i++ {
			nx := x + dir_8[i][0]
			ny := y + dir_8[i][1]
			if 0 <= nx && nx < n && 0 <= ny && ny < m && land[nx][ny] == 0 {
				dfs(nx, ny)
			}
		}
	}

	for i, col := range land {
		for j := range col {
			if land[i][j] == 0 {
				area = 0
				dfs(i, j)
				ans = append(ans, area)
			}
		}
	}
	sort.Ints(ans)
	return
}
func Test_pond_sizes_lcci(t *testing.T) {
	tests := []struct {
		land [][]int
	}{
		{[][]int{
			{0, 2, 1, 0},
			{0, 1, 0, 1},
			{1, 1, 0, 1},
			{0, 1, 0, 1},
		}},
	}
	for _, tt := range tests {
		t.Log(pondSizes(tt.land))
	}
}
