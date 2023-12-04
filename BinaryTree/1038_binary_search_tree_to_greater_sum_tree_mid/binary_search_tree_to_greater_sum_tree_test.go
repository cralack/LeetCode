package binarysearchtreetogreatersumtree

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
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Right)
		sum += cur.Val
		cur.Val = sum
		dfs(cur.Left)
	}

	dfs(root)

	return root
}

func Test_binary_search_tree_to_greater_sum_tree(t *testing.T) {
	tests := []struct {
		root *TreeNode
	}{
		{Init([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})},
	}
	for _, tt := range tests {
		bstToGst(tt.root).Show()
	}
}
