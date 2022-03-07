package permutations

import (
	"testing"
)

func permute(nums []int) (res [][]int) {
	// res := [][]int{}
	n := len(nums)
	var backtrack func(first int)
	backtrack = func(first int) {
		if first == n {
			tmp := make([]int, n)
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return res
}

func TestPermutations(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	t.Log(ans)
}
