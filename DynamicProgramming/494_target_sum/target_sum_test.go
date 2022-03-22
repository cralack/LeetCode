package targetsum

import (
	"strconv"
	"testing"
)

func findTargetSumWays_BackTrack(nums []int, target int) int {
	// base case
	if len(nums) == 0 {
		return 0
	}
	result := 0

	var backtrack func(i, remain int)
	backtrack = func(i, remain int) {
		// base case
		if i == len(nums) {
			if remain == 0 {
				// 说明恰好凑出 target
				result++
			}
			return
		}
		// 给 nums[i] 选择 - 号
		remain += nums[i]
		backtrack(i+1, remain)
		remain -= nums[i]
		// 给 nums[i] 选择 + 号
		remain -= nums[i]
		backtrack(i+1, remain)
		remain += nums[i]
	}
	backtrack(0, target)
	return result
}
func findTargetSumWays_memo(nums []int, target int) int {
	memo := make(map[string]int, 0)
	if len(nums) == 0 {
		return 0
	}
	var dp func(i, remain int) int
	dp = func(i, remain int) int {
		// base case
		if i == len(nums) {
			if remain == 0 {
				return 1
			}
			return 0
		}
		key := strconv.Itoa(i) + "," + strconv.Itoa(remain)
		if _, ok := memo[key]; ok {
			return memo[key]
		}
		result := dp(i+1, remain-nums[i]) + dp(i+1, remain+nums[i])
		memo[key] = result
		return result
	}
	return dp(0, target)
}

func findTargetSumWays_dp(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target) < 0 || (sum+target)%2 != 0 {
		return 0
	}
	//a+b=tar
	//a-b=sum
	//a=(sum+tar)/2
	sum = (sum + target) / 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := sum; j >= nums[i-1]; j-- {
			dp[j] += dp[j-nums[i-1]]
		}
	}
	return dp[sum]
}

func Test_target_sum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	t.Log(findTargetSumWays_BackTrack(nums, 5))
	t.Log(findTargetSumWays_memo(nums, 5))
	t.Log(findTargetSumWays_dp(nums, 5))
}

func Benchmark_target_sum(b *testing.B) {
	nums := []int{1, 1, 1, 1, 1}
	b.Run("BackTrack", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findTargetSumWays_BackTrack(nums, 3)
		}
		b.StopTimer()
	})
	b.Run("memo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findTargetSumWays_memo(nums, 3)
		}
		b.StopTimer()
	})
	b.Run("DP", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findTargetSumWays_dp(nums, 3)
		}
		b.StopTimer()
	})
}
