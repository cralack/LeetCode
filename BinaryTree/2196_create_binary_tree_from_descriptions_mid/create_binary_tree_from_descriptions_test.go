package create_binary_tree_from_descriptions_mid

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
func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode, len(descriptions)+1)
	root := 0

	for _, d := range descriptions {
		p, c := d[0], d[1]
		if nodes[p] == nil {
			nodes[p] = &TreeNode{Val: p}
			root ^= p
		}
		if nodes[c] == nil {
			nodes[c] = &TreeNode{Val: c}
			root ^= c
		}
		if d[2] == 1 {
			nodes[p].Left = nodes[c]
		} else {
			nodes[p].Right = nodes[c]
		}
		root ^= c
	}

	return nodes[root]
}

func Test_create_binary_tree_from_descriptions(t *testing.T) {
	tests := []struct {
		descriptions [][]int
	}{
		{descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}},
		{descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}}},
	}
	for _, test := range tests {
		createBinaryTree(test.descriptions).Show()
	}
}
