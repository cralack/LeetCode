package maximum_difference_between_node_and_ancestor_mid

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
func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode, int, int)
	dfs = func(cur *TreeNode, mn, mx int) {
		if cur == nil {
			ans = max(ans, mx-mn)
			return
		}
		mn = min(mn, cur.Val)
		mx = max(mx, cur.Val)
		dfs(cur.Left, mn, mx)
		dfs(cur.Right, mn, mx)
	}

	dfs(root, root.Val, root.Val)
	return ans
}

func Test_maximum_difference_between_node_and_ancestor(t *testing.T) {
	tests := []struct{ root *TreeNode }{
		{Init([]int{8, 3, 10, 1, 6, -1, 14, -1, -1, 4, 7, 13})},
		{Init([]int{1, -1, 2, -1, 0, 3})},
	}
	for _, tt := range tests {
		t.Log(maxAncestorDiff(tt.root))
	}
}
