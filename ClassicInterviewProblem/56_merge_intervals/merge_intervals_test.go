package mergeintervals

import (
	"sort"
	"testing"
)

func merge(intervals [][]int) (res [][]int) {
	// n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right && right <= intervals[i][1] {
			right = intervals[i][1]
		}
		if right < intervals[i][0] {
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	res = append(res, []int{left, right})
	return
}
func Test_merge_intervals(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	t.Log(merge(intervals))
	intervals = [][]int{{1, 4}, {4, 5}}
	t.Log(merge(intervals))
}
