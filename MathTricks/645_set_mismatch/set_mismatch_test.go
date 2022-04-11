package setmismatch

import (
	"testing"
)

func findErrorNums(nums []int) []int {
	//也可以使用hash表遍历两遍得出cnt
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	n := len(nums)
	dup := -1
	for i := 0; i < n; i++ {
		idx := abs(nums[i]) - 1
		// nums[index] 小于 0 则说明重复访问
		if nums[idx] < 0 {
			dup = abs(nums[i])
		} else {
			nums[idx] *= -1
		}
	}
	missing := -1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			missing = i + 1
		}
	}

	return []int{dup, missing}
}
func Nor_test(nums []int) []int {
	res := 0
	for i := range nums {
		res ^= i
		res ^= nums[i]
	}
	return []int{res}
}
func Test_set_mismatch(t *testing.T) {
	t.Log(findErrorNums([]int{1, 2, 2, 4}))
	t.Log(findErrorNums([]int{1, 1}))
	t.Log(findErrorNums([]int{6, 3, 2, 7, 1, 2, 4}))
	Nor_test([]int{6, 3, 2, 7, 1, 2, 4})
}
