package equalsumarrayswithminimumnumberofoperations

import (
	"testing"
)

func minOperations(nums1 []int, nums2 []int) (ans int) {
	if 6*len(nums1) < len(nums2) || 6*len(nums2) < len(nums1) {
		return -1 // 优化
	}
	// 数组元素和的差，我们要让这个差变为 0
	dif := 0
	for _, x := range nums2 {
		dif += x
	}
	for _, x := range nums1 {
		dif -= x
	}
	// 统一让 nums1 的数变大，nums2 的数变小
	if dif < 0 {
		dif = -dif
		nums1, nums2 = nums2, nums1
	}
	// 统计每个数的最大变化量
	cnt := [6]int{}
	// nums1 的变成 6
	for _, x := range nums1 {
		cnt[6-x]++
	}
	// nums2 的变成 1
	for _, x := range nums2 {
		cnt[x-1]++
	}
	// 从大到小枚举最大变化量 5 4 3 2 1
	for i := 5; ; i-- {
		// 可以让 d 变为 0
		if i*cnt[i] >= dif {
			return ans + (dif+i-1)/i //向上取整
		}
		// 需要所有最大变化量为 i 的数
		ans += cnt[i]
		dif -= i * cnt[i]
	}
}

func Test_minimum_number_of_operations(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{1, 1, 2, 2, 2, 2}
	t.Log(minOperations(nums1, nums2))
	nums1 = []int{6, 6}
	nums2 = []int{1}
	t.Log(minOperations(nums1, nums2))
}
