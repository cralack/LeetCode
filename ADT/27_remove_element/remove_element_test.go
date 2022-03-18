package removeelement

import "testing"

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
func Test_remove_element(t *testing.T) {
	arr, val := []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	arr = arr[:removeElement(arr, val)]
	t.Log(arr)
}
