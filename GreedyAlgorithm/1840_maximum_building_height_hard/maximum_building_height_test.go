package maximum_building_height_hard

import (
	"slices"
	"testing"
)

func maxBuilding(n int, restri [][]int) int {
	m := len(restri)
	if m == 0 {
		return n - 1
	}
	slices.SortFunc(restri, func(a, b []int) int {
		return a[0] - b[0]
	})

	h := make([]int, m)
	h[0] = min(restri[0][0]-1, restri[0][1])
	for i := 1; i < m; i++ {
		h[i] = min(h[i-1]+restri[i][0]-restri[i-1][0], restri[i][1])
	}
	for i := m - 2; i >= 0; i-- {
		h[i] = min(h[i], h[i+1]+restri[i+1][0]-restri[i][0])
	}
	ans := max((restri[0][0]-1+h[0])/2, h[m-1]+n-restri[m-1][0])
	for i := range m - 1 {
		ans = max(ans, (restri[i+1][0]-restri[i][0]+h[i]+h[i+1])/2)
	}
	return ans
}

func Test_maximum_building_height(t *testing.T) {
	tests := []struct {
		n            int
		restrictions [][]int
	}{
		{n: 5, restrictions: [][]int{{2, 1}, {4, 1}}},
		{n: 6, restrictions: [][]int{}},
		{n: 1, restrictions: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}},
	}
	for _, test := range tests {
		t.Log(maxBuilding(test.n, test.restrictions))
	}
}
