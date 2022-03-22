package targetsum

import "testing"

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
func Test_target_sum(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	t.Log(findTargetSumWays_BackTrack(nums, 3))
}
