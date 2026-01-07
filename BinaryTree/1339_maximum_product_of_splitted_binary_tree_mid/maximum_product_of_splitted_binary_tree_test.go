package maximum_product_of_splitted_binary_tree_mid

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
const mod int = 1e9 + 7

func maxProduct(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return root.Val + dfs(root.Left) + dfs(root.Right)
	}
	totalSum := dfs(root)

	var dfs2 func(root *TreeNode) int
	dfs2 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val + dfs2(root.Left) + dfs2(root.Right)
		ans = max(ans, (totalSum-sum)*sum)
		return sum
	}
	dfs2(root)
	return ans % mod
}

func Test_maximum_product_of_splitted_binary_tree(t *testing.T) {
	tests := []struct {
		root []int
	}{
		{root: []int{1, 2, 3, 4, 5, 6}},
		{root: []int{1, -1, 2, 3, 4, -1, -1, 5, 6}},
		{root: []int{2, 3, 9, 10, 7, 8, 6, 5, 4, 11, 1}},
		{root: []int{1, 1}},
	}
	for _, tt := range tests {
		root := Init(tt.root)
		root.Show()
		t.Log(maxProduct(root))
	}
}
