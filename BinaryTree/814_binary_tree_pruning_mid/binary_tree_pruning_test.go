package binarytreepruningmid

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
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
func Test_binary_tree_pruning(t *testing.T) {
	root := pruneTree(Init([]int{1, -1, 0, 0, 1}))
	root.Show()
	root = pruneTree(Init([]int{1, 0, 1, 0, 0, 0, 1}))
	root.Show()
	root = pruneTree(Init([]int{1, 1, 0, 1, 1, 0, 1, 0}))
	root.Show()
}
