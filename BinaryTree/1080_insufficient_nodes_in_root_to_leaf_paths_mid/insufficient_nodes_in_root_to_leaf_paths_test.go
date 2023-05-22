package insufficientnodesinroottoleafpaths

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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	limit -= root.Val
	if root.Left == root.Right {
		if limit > 0 {
			return nil
		}
		return root
	}
	root.Left = sufficientSubset(root.Left, limit)
	root.Right = sufficientSubset(root.Right, limit)
	if root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}

func Test_insufficient_nodes_in_root_to_leaf_paths(t *testing.T) {
	tests := []struct {
		arr   []int
		limit int
	}{
		{[]int{1, 2, 3, 4, -99, -99, 7, 8, 9, -99, -99, 12, 13, -99, 14}, 1},
		{[]int{5, 4, 8, 11, -1, 17, 4, 7, 1, -1, -1, 5, 3}, 22},
		{[]int{1, 2, -3, -5, -1, 4, -1}, -1},
	}
	for _, tt := range tests {
		root := Init(tt.arr)
		root.Show()
		root = sufficientSubset(root, tt.limit)
		root.Show()
	}
}
