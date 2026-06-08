package partition_array_according_to_given_pivot_mid

import (
	"testing"
)

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	ans := make([]int, n)

	// 三个指针，分别指向：
	// left:  存放小于 pivot 的下一个位置（从 0 开始）
	// right: 存放大于 pivot 的下一个位置（从末尾往前，所以倒着填）
	left := 0
	right := n - 1

	// 第一趟遍历：双指针同时从两头推进
	// 完美的相对顺序：小于 pivot 的从左往右填，大于 pivot 的从右往左填
	for i := 0; i < n; i++ {
		if nums[i] < pivot {
			ans[left] = nums[i]
			left++
		}
		// 注意：为了保持大于 pivot 元素的原始顺序，
		// 既然我们是从右往左填入 ans，那我们就应该从右往左遍历 nums！
		if nums[n-1-i] > pivot {
			ans[right] = nums[n-1-i]
			right--
		}
	}

	// 第二趟：把中间空出来的部分全部填满 pivot
	for left <= right {
		ans[left] = pivot
		left++
	}

	return ans
}

func Test_partition_array_according_to_given_pivot(t *testing.T) {
	tests := []struct {
		nums  []int
		pivot int
	}{
		{nums: []int{9, 12, 5, 10, 14, 3, 10}, pivot: 10},
		{nums: []int{-3, 4, 3, 2}, pivot: 2},
	}
	for _, tt := range tests {
		t.Log(pivotArray(tt.nums, tt.pivot))
	}
}
