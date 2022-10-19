package jumpgamevi

import (
	"testing"
)

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	//单调递减队列,que[0]为(i-k,i)中最大元素
	que := []int{nums[0]}
	for i := 1; i < n; i++ {
		// dp[i] = Max(dp[i-k, i]) + nums[i]
		cur := que[0] + nums[i]
		dp[i] = cur
		// len(que) > k 时弹出队首
		if k <= i && 0 < len(que) && dp[i-k] == que[0] {
			que = que[1:]
		}
		//保证队尾元素>=cur
		for 0 < len(que) && que[len(que)-1] < cur {
			que = que[:len(que)-1]
		}
		que = append(que, cur)
	}
	return dp[n-1]
}

func Test_jump_game_vi(t *testing.T) {
	nums, k := []int{1, -1, -2, 4, -7, 3}, 2
	t.Log(maxResult(nums, k))
	nums, k = []int{10, -5, -2, 4, 0, 3}, 3
	t.Log(maxResult(nums, k))
	nums, k = []int{1, -5, -20, 4, -1, 3, -6, -3}, 2
	t.Log(maxResult(nums, k))
	nums, k = []int{100, -1, -100, -1, 100}, 2
	t.Log(maxResult(nums, k))
}
