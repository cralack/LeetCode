package permutationsii

import (
	"sort"
	"testing"
)

func permuteUnique(nums []int) (res [][]int) {
	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		// base case
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
			// fmt.Println(tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 剪枝逻辑，固定相同的元素在排列中的相对位置
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()
	return res
}
func Test_permutations_ii(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(permuteUnique(nums))
	nums = []int{1, 1, 2}
	t.Log(permuteUnique(nums))
	nums = []int{0, 1, 0, 0, 9}
	t.Log(permuteUnique(nums))
}
