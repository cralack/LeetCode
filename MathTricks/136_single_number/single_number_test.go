package singlenumber

import "testing"

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
func Test_single_number(t *testing.T) {
	nums := []int{2, 2, 1}
	t.Log(singleNumber(nums))
	nums = []int{4, 1, 2, 1, 2}
	t.Log(singleNumber(nums))
}
