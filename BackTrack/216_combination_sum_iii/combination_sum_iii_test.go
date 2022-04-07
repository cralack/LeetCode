package combinationsumiii

import "testing"

func combinationSum3(n int, tar int) (res [][]int) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	path := make([]int, 0)
	sum := 0
	var dfs func(start int)
	dfs = func(start int) {
		if sum > tar || len(path) > n {
			return
		}
		if sum == tar && len(path) == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			sum += nums[i]
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
			sum -= nums[i]
		}
	}
	dfs(0)
	return res
}
func Test_combination_sum_iii(t *testing.T) {
	k, n := 3, 7
	t.Log(combinationSum3(k, n))
	k, n = 3, 9
	t.Log(combinationSum3(k, n))
	k, n = 4, 1
	t.Log(combinationSum3(k, n))
}
