package maximumbinarytreeiimid

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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		return &TreeNode{Val: val, Left: root}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
func Test_maximum_binary_tree_ii(t *testing.T) {
	root := Init([]int{4, 1, 3, -1, -1, 2})
	val := 5
	root = insertIntoMaxTree(root, val)
	root.Show()
	root = Init([]int{5, 2, 4, -1, 1})
	val = 3
	root = insertIntoMaxTree(root, val)
	root.Show()
	root = Init([]int{5, 2, 3, -1, 1})
	val = 4
	root = insertIntoMaxTree(root, val)
	root.Show()
}
