package jumpgame

import "testing"

func canJump(nums []int) bool {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	farthest := 0
	for i, num := range nums {
		farthest = max(farthest, i+num)
		if farthest <= i {
			return i == n-1 || false
		}
	}
	return farthest >= n-1
}
func Test_jump_game(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	t.Log(canJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	t.Log(canJump(nums))
	t.Log(canJump([]int{3, 0, 8, 2, 0, 0, 1}))
	t.Log(canJump([]int{2, 0, 0}))
	t.Log(canJump([]int{0}))

}
