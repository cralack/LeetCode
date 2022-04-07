package combinationsumii

import (
	"sort"
	"testing"
)

func combinationSum2(candidates []int, target int) (res [][]int) {
	sum := 0
	path := make([]int, 0)
	sort.Ints(candidates)
	var dfs func(start int)
	dfs = func(start int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i + 1)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
func Test_combination_sum_ii(t *testing.T) {
	nums, target := []int{10, 1, 2, 7, 6, 1, 5}, 8
	t.Log(combinationSum2(nums, target))
	nums, target = []int{2, 5, 2, 1, 2}, 5
	t.Log(combinationSum2(nums, target))
}
