package rangeaddition

import "testing"

func getModifiedArray(length int, updates [][]int) []int {
	//all 0 this time
	nums := make([]int, length)
	dif := make([]int, len(nums))
	// dif[0] = nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	dif[i] = nums[i] - nums[i-1]
	// }
	for _, update := range updates {
		left, right := update[0], update[1]
		val := update[2]
		dif[left] += val
		if right+1 < len(dif) {
			dif[right+1] -= val
		}
	}
	nums[0] = dif[0]
	for i := 1; i < len(dif); i++ {
		nums[i] = nums[i-1] + dif[i]
	}
	return nums
}
func Test_range_addition(t *testing.T) {
	length := 5
	updates := [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}
	t.Log(getModifiedArray(length, updates))
}
