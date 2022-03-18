package movezeroes

import "testing"

func moveZeroes(nums []int) {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
func Test_move_zeroes(t *testing.T) {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	t.Log(arr)
}
