package findpeakelement

import "testing"

func findPeakElement(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] < nums[mi+1] {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}
func Test_find_peak_element(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	t.Log(findPeakElement(nums))
	nums = []int{1, 2, 1, 3, 5, 6, 4}
	t.Log(findPeakElement(nums))
}
