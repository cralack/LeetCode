package removeelement

import "testing"

func removeElement(nums []int, val int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		for nums[lo] != val {
			lo++
		}
		for nums[hi] == val {
			hi--
		}
		if lo < hi {
			nums[lo], nums[hi] = nums[hi], nums[lo]
		}
	}
	return lo
}
func Test_remove_element(t *testing.T) {
	arr, val := []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	t.Log(removeElement(arr, val))
	t.Log(arr)
}
