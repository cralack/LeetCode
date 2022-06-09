package findkthsmallestpairdistance

import (
	"sort"
	"testing"
)

func smallestDistancePair(nums []int, k int) int {
	// 为了二分查找，需要将nums排序，找到每次比较的最大距离
	sort.Ints(nums)
	// lo, hi分别是最大距离和最小距离，这个最小距离可以按照0开始
	lo, hi := 0, nums[len(nums)-1]-nums[0]
	for lo < hi {
		mi := lo + (hi-lo)>>1
		// 假如距离为mi的排位在k后面的话，结果在左边，hi需要前移
		if check(nums, mi) >= k {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	// 循环结束时lo==hi
	return hi
}

// 求dis在nums中，距离排第几。（滑动窗口）
func check(nums []int, dis int) int {
	cnt, left := 0, 0
	for right := 1; right < len(nums); right++ {
		for nums[right]-nums[left] > dis {
			left++
		}
		cnt += right - left
	}
	return cnt
}

func Test_find_k_th_smallest_pair_distance(t *testing.T) {
	nums, k := []int{1, 3, 1}, 1
	t.Log(smallestDistancePair(nums, k))
	nums, k = []int{1, 1, 1}, 2
	t.Log(smallestDistancePair(nums, k))
	nums, k = []int{1, 6, 1}, 3
	t.Log(smallestDistancePair(nums, k))
}
