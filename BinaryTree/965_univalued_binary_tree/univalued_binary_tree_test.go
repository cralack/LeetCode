package univaluedbinarytree

import (
	"testing"

	. "LeetCode/util/BinTree"
)

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	var dfs func(cur *TreeNode) bool
	dfs = func(cur *TreeNode) bool {
		if cur == nil {
			return true
		}
		if cur.Val != val {
			return false
		}
		return dfs(cur.Left) && dfs(cur.Right)
	}
	return dfs(root)
}
func Test_univalued_binary_tree(t *testing.T) {
	root := Init([]int{1, 1, 1, 1, 1, -1, 1})
	t.Log(isUnivalTree(root))
	root = Init([]int{2, 2, 2, 5, 2})
	t.Log(isUnivalTree(root))
}
