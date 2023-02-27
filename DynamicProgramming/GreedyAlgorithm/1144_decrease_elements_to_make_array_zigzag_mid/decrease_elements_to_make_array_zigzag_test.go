package decreaseelementstomakearrayzigzag

import (
	"testing"
)

func movesToMakeZigzag(nums []int) int {
	ans := make([]int, 2)
	n := len(nums)
	for i := 0; i < 2; i++ {
		for j := i; j < n; j += 2 {
			d := 0
			if j > 0 {
				d = max(d, nums[j]-nums[j-1]+1)
			}
			if j < n-1 {
				d = max(d, nums[j]-nums[j+1]+1)
			}
			ans[i] += d
		}
	}
	return min(ans[0], ans[1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_decrease_elements_to_make_array_zigzag(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(movesToMakeZigzag(nums))
	nums = []int{9, 6, 1, 6, 2}
	t.Log(movesToMakeZigzag(nums))
}
