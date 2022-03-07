package searchinabinarysearchtree

import (
	. "Learning/Leetcode/BinaryTree/BinTree"
	"testing"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
func Test_search_in_a_binary_search_tree(t *testing.T) {
	arr := []int{4, 2, 7, 1, 3}
	root := Init(arr)
	node := searchBST(root, 2)
	node.Show()
}
