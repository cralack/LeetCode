package combinations

import "testing"

func combine(n int, k int) (res [][]int) {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	path := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
func Test_combinations(t *testing.T) {
	n, k := 4, 2
	t.Log(combine(n, k))
}
