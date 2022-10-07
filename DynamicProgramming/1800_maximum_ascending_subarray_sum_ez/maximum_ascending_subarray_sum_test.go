package maximumascendingsubarraysum

import "testing"

func maxAscendingSum(nums []int) int {
	n := len(nums)
	cur := nums[0]
	max := cur
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		if max < cur {
			max = cur
		}
	}
	return max
}
func Test_maximum_ascending_subarray_sum(t *testing.T) {
	nums := []int{10, 20, 30, 5, 10, 50}
	t.Log(maxAscendingSum(nums))
	nums = []int{10, 20, 30, 40, 50}
	t.Log(maxAscendingSum(nums))
	nums = []int{12, 17, 15, 13, 10, 11, 12}
	t.Log(maxAscendingSum(nums))
	nums = []int{100, 10, 1}
	t.Log(maxAscendingSum(nums))
}
