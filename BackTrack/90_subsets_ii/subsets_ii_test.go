package subsetsii

import (
	"fmt"
	"sort"
	"testing"
)

func subsetsWithDup(nums []int) (res [][]int) {
	path := make([]int, 0)
	sort.Ints(nums)
	var dfs func(start int)
	dfs = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		fmt.Println(tmp)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
func Test_subsets_ii(t *testing.T) {
	nums := []int{1, 2, 2}
	t.Log(subsetsWithDup(nums))
}
