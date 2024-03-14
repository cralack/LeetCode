package jumpgameii

import "testing"

func jump(nums []int) int {
	n := len(nums)
	end, farthest, jumps := 0, 0, 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, nums[i]+i)
		if end == i {
			jumps++
			end = farthest
		}
	}
	return jumps
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_jump_game_ii(t *testing.T) {
	nums := []int{2, 1, 1, 1, 4}
	t.Log(jump(nums))
	t.Log(jump([]int{3, 1, 4, 2, 5, 5, 5, 7, 1, 1, 1, 1, 1}))
}
