package numberofsubarray

import "testing"

func numSubarrayBoundedMax(nums []int, left, right int) (ans int) {
	// 双指针
	i0, i1 := -1, -1
	for i, x := range nums {
		if x > right {
			i0 = i
		}
		if x >= left {
			i1 = i
		} // 贡献度
		ans += i1 - i0
	}
	return
}

func Test_123(t *testing.T) {
	nums, left, right := []int{2, 1, 4, 3}, 2, 3
	t.Log(numSubarrayBoundedMax(nums, left, right))
	nums, left, right = []int{2, 9, 2, 5, 6}, 2, 8
	t.Log(numSubarrayBoundedMax(nums, left, right))

}
