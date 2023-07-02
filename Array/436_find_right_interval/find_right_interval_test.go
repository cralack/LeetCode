package findrightinterval

import (
	"sort"
	"testing"
)

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	type pair struct {
		pos, idx int
	}
	starts := make([]pair, n)
	ends := make([]pair, n)
	for i, p := range intervals {
		starts[i] = pair{pos: p[0], idx: i}
		ends[i] = pair{pos: p[1], idx: i}
	}
	sort.Slice(starts, func(i, j int) bool { return starts[i].pos < starts[j].pos })
	sort.Slice(ends, func(i, j int) bool { return ends[i].pos < ends[j].pos })

	res := make([]int, n)
	// sort
	j := 0
	for _, end := range ends {
		for j < n && starts[j].pos < end.pos {
			j++
		}
		if j < n {
			res[end.idx] = starts[j].idx
		} else {
			res[end.idx] = -1
		}
	}
	return res
}
func Test_find_right_interval(t *testing.T) {
	intervals := [][]int{{1, 2}}
	t.Log(findRightInterval(intervals))
	intervals = [][]int{{3, 4}, {2, 3}, {1, 2}}
	t.Log(findRightInterval(intervals))
	intervals = [][]int{{1, 4}, {2, 3}, {3, 4}}
	t.Log(findRightInterval(intervals))
}
