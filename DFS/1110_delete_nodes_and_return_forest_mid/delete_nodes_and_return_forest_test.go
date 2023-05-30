package deletenodesandreturnforest

import (
	. "LeetCode/BinaryTree/BinTree"
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

func delNodes(root *TreeNode, to_delete []int) (ans []*TreeNode) {
	dic := make(map[int]struct{})
	for _, del := range to_delete {
		dic[del] = struct{}{}
	}
	var dfs func(*TreeNode) *TreeNode
	dfs = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}
		cur.Left = dfs(cur.Left)
		cur.Right = dfs(cur.Right)
		if _, has := dic[cur.Val]; !has {
			return cur
		}
		if cur.Left != nil {
			ans = append(ans, cur.Left)
		}
		if cur.Right != nil {
			ans = append(ans, cur.Right)
		}
		return nil
	}
	if dfs(root) != nil {
		ans = append(ans, root)
	}
	return
}
func Test_delete_nodes_and_return_forest(t *testing.T) {
	tests := []struct {
		root      *TreeNode
		to_delete []int
	}{
		{root: Init([]int{1, 2, 3, 4, 5, 6, 7}), to_delete: []int{3, 5}},
		{root: Init([]int{1, 2, 4, -1, 3}), to_delete: []int{3}},
	}
	for _, tt := range tests {
		for _, tree := range delNodes(tt.root, tt.to_delete) {
			tree.Show()
		}
	}
}
