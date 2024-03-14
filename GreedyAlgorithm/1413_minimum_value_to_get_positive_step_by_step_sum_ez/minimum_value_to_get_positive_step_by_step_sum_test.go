package minimumvaluetogetpositivestepbystepsumez

import "testing"

func minStartValue(nums []int) int {
	sum, minSum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if minSum > sum {
			minSum = sum
		}
	}
	return 1 - minSum
}
func Test_minimum_value_to_get_positive_step_by_step_sum(t *testing.T) {
	nums := []int{-3, 2, -3, 4, 2}
	t.Log(minStartValue(nums))
	nums = []int{1, 2}
	t.Log(minStartValue(nums))
	nums = []int{1, -2, -3}
	t.Log(minStartValue(nums))
	nums = []int{2, 3, 5, -5, -1}
	t.Log(minStartValue(nums))
}
