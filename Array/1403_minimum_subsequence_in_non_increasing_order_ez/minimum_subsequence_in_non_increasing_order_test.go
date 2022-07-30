package minimumsubsequenceinnonincreasingorderez

import (
	"sort"
	"testing"
)

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	tar, sum := 0, 0
	for i := range nums {
		tar += nums[i]
	}
	for i, v := range nums {
		sum += v
		if tar-sum < sum {
			return nums[:i+1]
		}
	}
	return nums
}
func Test_minimum_subsequence_in_non_increasing_order(t *testing.T) {
	nums := []int{4, 3, 10, 9, 8}
	t.Log(minSubsequence(nums))
	nums = []int{4, 4, 7, 6, 7}
	t.Log(minSubsequence(nums))
	nums = []int{6}
	t.Log(minSubsequence(nums))
}
