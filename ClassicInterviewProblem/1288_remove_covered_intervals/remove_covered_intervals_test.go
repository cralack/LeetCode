package removecoveredintervals

import (
	"sort"
	"testing"
)

func removeCoveredIntervals(intervals [][]int) int {
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1]
		}
	})
	left, right := intervals[0][0], intervals[0][1]
	cnt := 0
	for i := 1; i < len(intervals); i++ { // 情况一，找到覆盖区间
		if left <= intervals[i][0] && right >= intervals[i][1] {
			cnt++
		} // 情况二，找到相交区间，合并
		if right >= intervals[i][0] && right <= intervals[i][1] {
			right = intervals[i][1]
		} // 情况三，完全不相交，更新起点和终点
		if right < intervals[i][0] {
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	return n - cnt
}
func Test_remove_covered_intervals(t *testing.T) {
	intervals := [][]int{{1, 4}, {3, 6}, {2, 8}}
	t.Log(removeCoveredIntervals(intervals))
	intervals = [][]int{{1, 4}, {1, 2}, {3, 4}}
	t.Log(removeCoveredIntervals(intervals))
}
