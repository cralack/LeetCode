package search_in_rotated_sorted_array

import (
	"testing"
)

func search(nums []int, target int) int {
	last := nums[len(nums)-1]
	left, right := -1, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		x := nums[mid]
		if target > last && x <= last {
			right = mid
		} else if x > last && target <= last {
			left = mid
		} else if x >= target {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[right] != target {
		return -1
	}
	return right
}

func Test_search_in_rotated_sorted_array(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
		{nums: []int{1}, target: 0},
	}
	for _, tt := range tests {
		t.Log(search(tt.nums, tt.target))
	}
}
