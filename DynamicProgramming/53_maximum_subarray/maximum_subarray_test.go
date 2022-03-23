package maximumsubarray

import "testing"

func maxSubArray(nums []int) int {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res < v {
				res = v
			}
		}
		return res
	}
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}
	return max(dp...)
}
func Test_maximum_subarray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(nums))
}
