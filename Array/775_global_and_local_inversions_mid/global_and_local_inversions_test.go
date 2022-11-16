package globalandlocalinversions

import "testing"

// 维护前缀最大值
// 局部倒置一定是全局倒置,但全局倒置不一定是局部倒置
// 即全局倒置的数量>=局部倒置的数量
func isIdealPermutation(nums []int) bool {
	mx := 0
	//枚举nums[i]，其中 2<=i<=n-1
	for i := 2; i < len(nums); i++ {
		//非相邻数字是否满足递增
		mx = max(mx, nums[i-2])
		if mx > nums[i] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_global_and_local_inversions(t *testing.T) {
	nums := []int{1, 0, 2}
	t.Log(isIdealPermutation(nums))
	nums = []int{1, 2, 0}
	t.Log(isIdealPermutation(nums))
	nums = []int{0, 2, 3, 4, 1}
	t.Log(isIdealPermutation(nums))
}
