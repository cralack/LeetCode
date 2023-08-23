package subarraysumequalsk

import (
	"testing"
)

func subarraySum(nums []int, k int) int {
	n := len(nums)
	preSum := make(map[int]int, 0)
	preSum[0] = 1
	cnt, sum0_i := 0, 0
	for i := 0; i < n; i++ {
		sum0_i += nums[i]
		sum0_j := sum0_i - k
		if _, ok := preSum[sum0_j]; ok {
			cnt += preSum[sum0_j]
		}
		preSum[sum0_i]++
	}
	return cnt
}
func Test_subarray_sum_equals_k(t *testing.T) {
	nums, k := []int{3, 5, 2, -2, 4, 1}, 5
	t.Log(subarraySum(nums, k))
	nums, k = []int{1, 1, 1}, 2
	t.Log(subarraySum(nums, k))
	nums, k = []int{1, 2, 3}, 3
	t.Log(subarraySum(nums, k))
}
