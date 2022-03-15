package advantageshuffle

import (
	"sort"
	"testing"
)

type Node struct {
	Idx, Val int
}

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	n2 := []Node{} //nums2_backup
	for i, n := range nums2 {
		n2 = append(n2, Node{Val: n, Idx: i})
	}
	sort.Slice(n2, func(i, j int) bool { return n2[i].Val > n2[j].Val })
	//升序排序
	sort.Ints(nums1)
	// nums1[left] 是最小值，nums1[right] 是最大值
	left, right := 0, n-1
	res := make([]int, n)

	for len(n2) > 0 {
		pair := n2[0]
		n2 = n2[1:]
		// maxval 是 nums2 中的最大值，i 是对应索引
		i, maxval := pair.Idx, pair.Val
		if maxval < nums1[right] {
			// 如果 nums1[right] 能胜过 maxval，那就自己上
			res[i] = nums1[right]
			right--
		} else {
			// 否则用最小值混一下，养精蓄锐
			res[i] = nums1[left]
			left++
		}
	}
	return res
}
func Test_advantage_shuffle(t *testing.T) {
	a, b := []int{2, 7, 11, 15}, []int{1, 10, 4, 11}
	t.Log(advantageCount(a, b))
	a, b = []int{12, 24, 8, 32}, []int{13, 25, 32, 11}
	t.Log(advantageCount(a, b))
}
