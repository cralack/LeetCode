package longestunivaluepathmid

import (
	. "LeetCode/util/BinTree"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}
		left := dfs(cur.Left)
		right := dfs(cur.Right)
		if cur.Left != nil && cur.Left.Val == cur.Val {
			left++
		} else {
			left = 0
		}
		if cur.Right != nil && cur.Right.Val == cur.Val {
			right++
		} else {
			right = 0
		}
		ans = max(ans, left+right)
		return max(left, right)
	}
	dfs(root)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_longest_univalue_path(t *testing.T) {
	root := Init([]int{5, 4, 5, 1, 1, 5})
	t.Log(longestUnivaluePath(root))
	root = Init([]int{1, 4, 5, 4, 4, 5})
	t.Log(longestUnivaluePath(root))
}
