package partitiontokequalsumsubsetsmid

import (
	"sort"
	"testing"
)

func canPartitionKSubsets(nums []int, k int) bool {
	tar := 0
	for _, n := range nums {
		tar += n
	}
	if tar%k != 0 {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	tar /= k
	n := len(nums)
	used := make([]bool, n)

	var dfs func(idx, sum, valid int) bool
	dfs = func(idx, sum, valid int) bool {
		if valid == k {
			return true
		}
		if sum == tar {
			return dfs(0, 0, valid+1)
		}
		for i := idx; i < n; i++ {
			if used[i] || sum+nums[i] > tar {
				continue
			}
			used[i] = true
			if dfs(i+1, sum+nums[i], valid) {
				return true
			}
			used[i] = false
			if sum == 0 {
				return false
			}
		}
		return false
	}

	return dfs(0, 0, 0)
}
func Test_partition_to_k_equal_sum_subsets(t *testing.T) {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	t.Log(canPartitionKSubsets(nums, k))
	nums = []int{1, 2, 3, 4}
	k = 3
	t.Log(canPartitionKSubsets(nums, k))
	nums = []int{2, 2, 2, 2, 3, 4, 5}
	k = 4
	t.Log(canPartitionKSubsets(nums, k))
}
