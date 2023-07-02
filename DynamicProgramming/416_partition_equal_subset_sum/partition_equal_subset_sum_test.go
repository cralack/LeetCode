package partitionequalsubsettar

import "testing"

func canPartition(nums []int) bool {
	// init
	n, tar := len(nums), 0
	for _, n := range nums {
		tar += n
	}
	if tar%2 != 0 {
		return false
	}
	tar /= 2
	dp := make([]bool, tar+1)
	// base case
	dp[0] = true
	// dp
	for i := 0; i < n; i++ {
		for j := tar; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[tar]
}
func Test_partition_equal_subset_tar(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	t.Log(canPartition(nums))
	nums = []int{1, 2, 3, 5}
	t.Log(canPartition(nums))
}
