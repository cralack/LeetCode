package combinationsum

import "testing"

func combinationSum(candidates []int, target int) (res [][]int) {
	sum := 0
	path := make([]int, 0)
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
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
func Test_combination_sum(t *testing.T) {
	nums, tar := []int{2, 3, 5}, 8
	t.Log(combinationSum(nums, tar))
}
