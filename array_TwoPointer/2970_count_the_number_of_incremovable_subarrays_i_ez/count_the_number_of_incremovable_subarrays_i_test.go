package count_the_number_of_incremovable_subarrays_i_ez

import (
	"testing"
)

func incremovableSubarrayCount(nums []int) int {
	n := len(nums)
	i := 0
	for i < n-1 && nums[i] < nums[i+1] {
		i++
	}
	if i == n-1 { // 每个非空子数组都可以移除
		return n * (n + 1) / 2
	}
	ans := i + 2 // 不保留后缀的情况，一共 i+2 个
	// 枚举保留的后缀为 nums[j:]
	for j := n - 1; j == n-1 || nums[j] < nums[j+1]; j-- {
		for i >= 0 && nums[i] >= nums[j] {
			i--
		}
		// 可以保留前缀 nums[:i+1], nums[:i], ..., nums[:0] 一共 i+2 个
		ans += i + 2
	}
	return ans
}

func Test_count_the_number_of_incremovable_subarrays_i(t *testing.T) {
	tests := []struct{ nums []int }{
		{nums: []int{1, 2, 3, 4}},
		{nums: []int{6, 5, 7, 8}},
		{nums: []int{8, 7, 6, 6}},
		{nums: []int{1, 3, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Log(incremovableSubarrayCount(tt.nums))
	}
}
