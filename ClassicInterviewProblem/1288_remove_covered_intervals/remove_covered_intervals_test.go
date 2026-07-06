package removecoveredintervals

import (
	"cmp"
	"slices"
	"testing"
)

func removeCoveredIntervals(intervals [][]int) (ans int) {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})
	maxRight := 0
	for _, p := range intervals {
		if p[1] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return
}

func Test_remove_covered_intervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
	}{
		{[][]int{{1, 4}, {3, 6}, {2, 8}}},
		{[][]int{{1, 4}, {1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Log(removeCoveredIntervals(tt.intervals))
	}
}
