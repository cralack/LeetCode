package spicalarrayez

import (
	"sort"
	"testing"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		x := n - i
		if (i == 0 || nums[i-1] < x) && x <= nums[i] {
			return x
		}
	}
	return -1
}
func Test_spical_array(t *testing.T) {
	nums := []int{3, 5}
	t.Log(specialArray(nums))
	nums = []int{0, 0}
	t.Log(specialArray(nums))
	nums = []int{0, 4, 3, 0, 4}
	t.Log(specialArray(nums))
	nums = []int{3, 6, 7, 7, 0}
	t.Log(specialArray(nums))
}
