package relocate_marbles_mid

import (
	"sort"
	"testing"
)

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) (ans []int) {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	for i, cur := range moveFrom {
		delete(set, cur)
		set[moveTo[i]] = struct{}{}
	}
	ans = make([]int, 0, len(set))
	for k := range set {
		ans = append(ans, k)
	}

	sort.Ints(ans)
	return
}

func Test_relocate_marbles(t *testing.T) {
	tests := []struct {
		nums     []int
		moveFrom []int
		moveTo   []int
	}{
		{nums: []int{1, 6, 7, 8}, moveFrom: []int{1, 7, 2}, moveTo: []int{2, 9, 5}},
		{nums: []int{1, 1, 3, 3}, moveFrom: []int{1, 3}, moveTo: []int{2, 2}},
	}
	for _, tt := range tests {
		t.Log(relocateMarbles(tt.nums, tt.moveFrom, tt.moveTo))
	}
}
