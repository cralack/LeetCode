package findfirstandlastpositionofelementinsortedarray

import "testing"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lo, hi := 0, len(nums)
	for lo < hi {
		mi := lo + (hi-lo)>>2
		if target <= nums[mi] {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	if lo == len(nums) || nums[lo] != target {
		return []int{-1, -1}
	}
	leftBound := lo
	lo, hi = leftBound, len(nums)
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if target >= nums[mi] {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return []int{leftBound, lo - 1}
}

func LeftBoundBinSearch(nums []int, tar int) int {
	lo, hi := 0, len(nums)

	for lo < hi {
		mi := lo + (hi-lo)>>1
		if tar == nums[mi] {
			hi = mi
		} else if tar < nums[mi] {
			hi = mi
		} else if tar > nums[mi] {
			lo = mi + 1
		}
	}
	if nums[lo] != tar {
		return -1
	}
	return lo
}

func RightBoundBinSearch(nums []int, tar int) int {

	lo, hi := 0, len(nums)
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if tar == nums[mi] {
			lo = mi + 1
		} else if tar < nums[mi] {
			hi = mi
		} else if tar > nums[mi] {
			lo = mi + 1
		}
	}
	lo--
	if nums[lo] != tar {
		return -1
	}
	return lo
}
func TestFindFirstAndLastPositionOfElement(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	tar1 := 8
	tar2 := 6
	res1 := searchRange(nums, tar1)
	res2 := searchRange(nums, tar2)
	t.Log(res1, res2)

	nums3 := []int{2, 2}
	tar3 := 3
	res3 := searchRange(nums3, tar3)
	t.Log(res3)

	// nums := []int{1, 2, 2, 2, 3, 4, 5, 6}
	// tar := 0
	// idx := RightBoundBinSearch(nums, tar)
	// res := nums[idx]
	// t.Log(idx, res)
}
