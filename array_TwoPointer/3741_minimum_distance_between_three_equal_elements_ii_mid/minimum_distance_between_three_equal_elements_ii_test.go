package minimum_distance_between_three_equal_elements_ii_mid

import (
	"math"
	"testing"
)

func minimumDistance(nums []int) (ans int) {
	idx := make(map[int][]int)
	for i, num := range nums {
		idx[num] = append(idx[num], i)
	}
	ans = math.MaxInt32
	for _, pos := range idx {
		for i := 2; i < len(pos); i++ {
			ans = min(ans, (pos[i]-pos[i-2])*2)
		}
	}
	if ans == math.MaxInt32 {
		ans = -1
	}
	return
}

func Test_minimum_distance_between_three_equal_elements_ii(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{[]int{1, 2, 1, 1, 3}},
		{[]int{1, 1, 2, 3, 2, 1, 2}},
		{[]int{1}},
	}
	for _, tt := range tests {
		t.Log(minimumDistance(tt.nums))
	}
}
