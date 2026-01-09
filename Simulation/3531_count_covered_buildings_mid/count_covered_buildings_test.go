package count_covered_buildings_mid

import (
	"testing"
)

func countCoveredBuildings(n int, buildings [][]int) (ans int) {
	type pair struct{ min, max int }
	row := make([]pair, n+1)
	col := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		row[i].min = 1 << 30
		col[i].min = 1 << 30
	}

	add := func(m []pair, x, y int) {
		m[y].min = min(m[y].min, x)
		m[y].max = max(m[y].max, x)
	}
	isInner := func(m []pair, x, y int) bool {
		return m[y].min < x && x < m[y].max
	}

	for _, point := range buildings {
		x, y := point[0], point[1]
		add(row, x, y)
		add(col, y, x)
	}

	for _, point := range buildings {
		x, y := point[0], point[1]
		if isInner(row, x, y) && isInner(col, y, x) {
			ans++
		}
	}
	return
}

func Test_count_covered_buildings(t *testing.T) {
	tests := []struct {
		n         int
		buildings [][]int
	}{
		{n: 3, buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}}},
		{n: 3, buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{n: 5, buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}}},
	}
	for _, tt := range tests {
		t.Log(countCoveredBuildings(tt.n, tt.buildings))
	}
}
