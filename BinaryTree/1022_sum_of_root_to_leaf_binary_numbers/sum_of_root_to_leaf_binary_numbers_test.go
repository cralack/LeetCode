package sumofroottoleafbinarynumbers

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
func sumRootToLeaf(root *TreeNode) int {
	var dfs func(cur *TreeNode, val int) int
	dfs = func(cur *TreeNode, val int) int {
		if cur == nil {
			return 0
		}
		val = val<<1 | cur.Val
		if cur.Left == nil && cur.Right == nil {
			return val
		}
		return dfs(cur.Left, val) + dfs(cur.Right, val)
	}
	return dfs(root, 0)
}
func Test_sum_of_root_to_leaf_binary_numbers(t *testing.T) {
	root := Init([]int{1, 0, 1, 0, 1, 0, 1})
	t.Log(sumRootToLeaf(root))
}
