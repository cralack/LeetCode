package subsets

import (
	"testing"
)

func subsets(nums []int) (res [][]int) {
	path := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		//前序遍历
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			// 做选择
			path = append(path, nums[i])
			// 通过 start 参数控制树枝的遍历，避免产生重复的子集
			dfs(i + 1)
			// 撤销选择
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(subsets(nums))
}
