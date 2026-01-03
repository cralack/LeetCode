package last_day_where_you_can_still_cross_hard

import (
	"testing"
)

type unionFind struct {
	father []int // 代表元
}

func newUnionFind(n int) unionFind {
	father := make([]int, n)
	for i := range father {
		father[i] = i
	}
	return unionFind{father}
}
func (u *unionFind) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind) same(x, y int) bool {
	return u.find(x) == u.find(y)
}
func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	u.father[x] = y
}

var dirs = []struct{ x, y int }{{0, -1}, {0, 1},
	{-1, 0}, {1, 0}}

func latestDayToCross(row, col int, cells [][]int) int {
	top := row * col
	bottom := row*col + 1
	uf := newUnionFind(row*col + 2)

	land := make([][]bool, row)
	for i := range land {
		land[i] = make([]bool, col)
	}

	for day := len(cells) - 1; ; day-- {
		cell := cells[day]
		r, c := cell[0]-1, cell[1]-1
		v := r*col + c
		land[r][c] = true

		if r == 0 {
			uf.merge(v, top)
		}
		if r == row-1 {
			uf.merge(v, bottom)
		}

		for _, dir := range dirs {
			nextX, nextY := r+dir.x, c+dir.y
			if 0 <= nextX && nextX < row && 0 <= nextY &&
				nextY < col && land[nextX][nextY] {
				uf.merge(v, nextX*col+nextY)
			}
		}
		if uf.same(top, bottom) {
			return day
		}
	}
}

func Test_last_day_where_you_can_still_cross(t *testing.T) {
	tests := []struct {
		row   int
		col   int
		cells [][]int
	}{
		{row: 2, col: 2, cells: [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}},
		{row: 2, col: 2, cells: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{row: 3, col: 3, cells: [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			t.Log(latestDayToCross(test.row, test.col, test.cells))
		})
	}
}
