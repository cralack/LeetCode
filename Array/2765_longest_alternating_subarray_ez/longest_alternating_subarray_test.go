package longest_alternating_subarray_ez

import (
	"testing"
)

func alternatingSubarray(nums []int) int {
	ans := -1
	i, n := 0, len(nums)
	for i < n-1 {
		if nums[i+1]-nums[i] != 1 {
			i++
			continue
		}
		i0 := i
		i += 2
		for i < n && nums[i] == nums[i0]+(i-i0)%2 {
			i++
		}
		ans = max(ans, i-i0)
		i--
	}
	return ans
}

func Test_longest_alternating_subarray(t *testing.T) {
	tests := []struct{ nums []int }{
		{[]int{2, 3, 4, 3, 4}},
		{[]int{4, 5, 6}},
		{[]int{3, 4, 3, 4, 5, 4, 5}},
	}
	for _, tt := range tests {
		t.Log(alternatingSubarray(tt.nums))
	}
}
