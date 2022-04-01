package nonoverlappingintervals

import (
	"sort"
	"testing"
)

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	//按end升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	//Cnt(不相交区间),前一不相交区间end
	cnt, x_end := 1, intervals[0][1]
	for i := 1; i < n; i++ {
		//若后一个start>=前一个end，则不相交区间+1
		if intervals[i][0] >= x_end {
			cnt++
			x_end = intervals[i][1]
		}
	}
	return n - cnt
}
func Test_non_overlapping_intervals(t *testing.T) {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	t.Log(eraseOverlapIntervals(intervals))
}
