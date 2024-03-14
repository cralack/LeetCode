package largest_element_in_an_array_after_merge_operations_mid

import (
	"testing"
)

func maxArrayValue(nums []int) int64 {
	n := len(nums)
	sum := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= sum {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
	}
	return int64(sum)
}

func Test_largest_element_in_an_array_after_merge_operations(t *testing.T) {
	tests := []struct{ nums []int }{
		{[]int{2, 3, 7, 9, 3}},
		{[]int{5, 3, 3}},
		{[]int{77}},
	}
	for _, tt := range tests {
		t.Log(maxArrayValue(tt.nums))
	}
}
