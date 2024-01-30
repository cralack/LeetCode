package minimum_seconds_to_equalize_a_circular_array_mid

import (
	"testing"
)

func minimumSeconds(nums []int) int {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	n := len(nums)
	ans := n
	for _, a := range pos {
		mx := n - a[len(a)-1] + a[0] // dis[tail~head]
		for i := 1; i < len(a); i++ {
			mx = max(mx, a[i]-a[i-1]) // max dis
		}
		ans = min(ans, mx/2) // 最后再除 2
	}

	return ans
}

func Test_minimum_seconds_to_equalize_a_circular_array(t *testing.T) {
	tests := []struct{ nums []int }{
		{[]int{1, 2, 1, 2}},
		{[]int{2, 1, 3, 3, 2}},
		{[]int{5, 5, 5, 5}},
	}
	for _, tt := range tests {
		t.Log(minimumSeconds(tt.nums))
	}
}
