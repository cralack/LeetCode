package mostfrequentsubtreesum

import (
	"testing"

	. "LeetCode/util/BinTree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	cnt := map[int]int{}
	maxCnt := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := dfs(node.Left) + dfs(node.Right) + node.Val
		cnt[sum]++
		if maxCnt < cnt[sum] {
			maxCnt = cnt[sum]
		}
		return sum
	}
	dfs(root)
	ans := []int{}
	for sum, c := range cnt {
		if c == maxCnt {
			ans = append(ans, sum)
		}
	}
	return ans
}
func Test_most_frequent_subtree_sum(t *testing.T) {
	arr := []int{5, 2, -3}
	root := Init(arr)
	t.Log(findFrequentTreeSum(root))
}
