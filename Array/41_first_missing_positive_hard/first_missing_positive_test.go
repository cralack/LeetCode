package firstmissingpositivehard

import "testing"

func firstMissingPositive(nums []int) int {
	// 缺失的最小正数只有两种情况，n:=len(nums)
	// 1.在[0,n]之间  2.为n+1

	// 将每个nums[i]放到对应位置
	for i := 0; i < len(nums); i++ {
		for 0 < nums[i] && nums[i] < len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	// 遍历修改后的切片，如果nums[i]!=i+1,i+1即为缺失的正数，如果每个元素都满足，则为n+1
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
func Test_first_missing_positive(t *testing.T) {
	nums := []int{1, 2, 0}
	t.Log(firstMissingPositive(nums))
	nums = []int{3, 4, -1, 1}
	t.Log(firstMissingPositive(nums))
	nums = []int{7, 8, 9, 11, 12}
	t.Log(firstMissingPositive(nums))
}
