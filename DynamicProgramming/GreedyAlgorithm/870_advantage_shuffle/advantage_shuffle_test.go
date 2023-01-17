package advantageshuffle

import (
	"sort"
	"testing"
)

type Node struct {
	Idx, Val int
}

func advantageCount(nums1, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	sort.Ints(nums1)

	//nums2的索引
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	//对nums2的索引排序
	sort.Slice(ids, func(i, j int) bool { return nums2[ids[i]] < nums2[ids[j]] })
	left, right := 0, n-1
	//能比直接比,不能比就放队尾
	for _, x := range nums1 {
		if x > nums2[ids[left]] {
			ans[ids[left]] = x // 用下等马比下等马
			left++
		} else {
			ans[ids[right]] = x // 用下等马比上等马
			right--
		}
	}
	return ans
}

func Test_advantage_shuffle(t *testing.T) {
	a, b := []int{2, 11, 7, 15}, []int{1, 10, 4, 11}
	t.Log(advantageCount(a, b))
	a, b = []int{12, 24, 8, 32}, []int{13, 25, 32, 11}
	t.Log(advantageCount(a, b))
}
