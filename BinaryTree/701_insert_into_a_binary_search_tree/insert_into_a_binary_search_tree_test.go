package insertintoabinarysearchtree

import (
	. "Learning/Leetcode/BinaryTree/BinTree"
	"testing"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
func Test_insert_into_a_binary_search_tree(t *testing.T) {
	arr := [][]int{{4, 2, 7, 1, 3}, {40, 20, 60, 10, 30, 50, 70}}
	root := Init(arr[1])
	insertIntoBST(root, 25)
	root.Show()
}
