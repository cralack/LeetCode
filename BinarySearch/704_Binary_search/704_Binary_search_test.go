package binarysearch

import "testing"

func TestBinSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	tar1, tar2 := 9, 2
	res1 := search(nums, tar1)
	res2 := search(nums, tar2)
	t.Log(res1, res2)
}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] == target {
			return mi
		} else if nums[mi] < target {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return -1
}
