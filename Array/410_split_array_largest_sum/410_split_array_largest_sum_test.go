package splitarraylargestsum

import "testing"

func splitArray(nums []int, m int) int {
	lo, hi := 1, 0
	for _, n := range nums {
		hi += n
	}
	if hi < lo {
		return 0
	}
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if valid(nums, m, mi) {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo
}
func valid(nums []int, m, tar int) bool {
	sum, cnt := 0, 1
	for _, n := range nums {
		if tar < n {
			return false
		}
		if sum+n <= tar {
			sum += n
		} else {
			cnt++
			sum = n
		}
	}
	return cnt <= m
}
func Test_split_array_largest_sum(t *testing.T) {
	// same with 1011
	nums, m := []int{7, 2, 5, 10, 8}, 2
	t.Log(splitArray(nums, m))
	nums, m = []int{1, 4, 4}, 3
	t.Log(splitArray(nums, m))
	nums, m = []int{0}, 1
	t.Log(splitArray(nums, m))
}
