package findlargestvalueineachtreerowmid

import (
	. "Learning/LeetCode/BinaryTree/BinTree"
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
func largestValues(root *TreeNode) (ans []int) {
	var dfs func(cur *TreeNode, curHeight int)
	dfs = func(cur *TreeNode, curHeight int) {
		if cur == nil {
			return
		}
		if curHeight == len(ans) {
			ans = append(ans, cur.Val)
		} else {
			ans[curHeight] = max(ans[curHeight], cur.Val)
		}
		dfs(cur.Left, curHeight+1)
		dfs(cur.Right, curHeight+1)
	}
	dfs(root, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_find_largest_value_in_each_tree_row(t *testing.T) {
	arr := []int{1, 3, 2, 5, 3, -1, 9}
	root := Init(arr)
	t.Log(largestValues(root))
}
