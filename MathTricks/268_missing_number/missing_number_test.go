package missingnumber

import (
	"testing"
)

func missingNumber(nums []int) int {
	res := 0
	nums = append(nums, 0)
	for i := 0; i < len(nums); i++ {
		res ^= i
		res ^= nums[i]
	}
	return res
}
func Test_missing_number(t *testing.T) {
	nums := []int{3, 0, 1}
	t.Log(missingNumber(nums))
	nums = []int{0, 1}
	t.Log(missingNumber(nums))
	nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	t.Log(missingNumber(nums))
	nums = []int{0}
	t.Log(missingNumber(nums))
}
