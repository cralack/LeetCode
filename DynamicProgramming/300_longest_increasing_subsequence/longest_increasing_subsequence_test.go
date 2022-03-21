package longestincreasingsubsequence

import "testing"

func lengthOfLIS(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i := range dp {
		res = max(res, dp[i])
	}
	return res
}
func Test_longest_increasing_subsequence(t *testing.T) {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	t.Log(lengthOfLIS(arr))
}
