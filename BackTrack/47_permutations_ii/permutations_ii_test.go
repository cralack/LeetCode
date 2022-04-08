package permutationsii

import (
	"sort"
	"testing"
)

func permuteUnique(nums []int) (res [][]int) {
	var dfs func(start int)
	dfs = func(start int) {
		//base case
		if start == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			//fmt.Println(tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			nums[start], nums[i] = nums[i], nums[start]
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	sort.Ints(nums)
	dfs(0)
	return res
}
func Test_permutations_ii(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(permuteUnique(nums))
	nums = []int{1, 1, 2}
	t.Log(permuteUnique(nums))
}
