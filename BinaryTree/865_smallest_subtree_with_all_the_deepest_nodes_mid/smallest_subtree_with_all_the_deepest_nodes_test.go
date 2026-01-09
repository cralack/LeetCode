package smallest_subtree_with_all_the_deepest_nodes_mid

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
func subtreeWithAllDeepest(root *TreeNode) (ans *TreeNode) {
	maxDeepth := -1
	var dfs func(root *TreeNode, depth int) int
	dfs = func(cur *TreeNode, depth int) int {
		if cur == nil {
			maxDeepth = max(maxDeepth, depth)
			return depth
		}
		leftMaxDeepth := dfs(cur.Left, depth+1)
		rightMaxDeepth := dfs(cur.Right, depth+1)

		if leftMaxDeepth == rightMaxDeepth && leftMaxDeepth == maxDeepth {
			ans = cur
		}
		return max(leftMaxDeepth, rightMaxDeepth)
	}
	dfs(root, 0)
	return
}

func Test_smallest_subtree_with_all_the_deepest_nodes(t *testing.T) {
	tests := []struct {
		root []int
	}{
		{root: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}},
		{root: []int{1}},
		{root: []int{0, 1, 3, -1, 2}},
	}
	for _, test := range tests {
		root := Init(test.root)
		root.Show()
		t.Log()
		subtreeWithAllDeepest(root).Show()
		t.Log()
	}
}
