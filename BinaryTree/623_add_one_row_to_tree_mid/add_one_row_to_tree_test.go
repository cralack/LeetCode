package addonerowtotreemid

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
func addOneRow(root *TreeNode, val int, tarDepth int) *TreeNode {
	if tarDepth == 1 {
		return &TreeNode{Val: val, Left: root, Right: nil}
	}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == tarDepth-1 {
			node.Left = &TreeNode{Val: val,
				Left: node.Left}
			node.Right = &TreeNode{Val: val,
				Right: node.Right}
			return
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	return root
}
func Test_add_one_row_to_tree(t *testing.T) {
	root := Init([]int{4, 2, 6, 3, 1, 5})
	val, depth := 1, 2
	addOneRow(root, val, depth)
	root.Show()

	root = Init([]int{4, 2, -1, 3, 1})
	val, depth = 1, 3
	addOneRow(root, val, depth)
	root.Show()

	root = Init([]int{4, 2, 6, 3, 1, 5})
	val, depth = 1, 1
	addOneRow(root, val, depth)
	root.Show()
}
