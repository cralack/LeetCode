package maximumlevelsumofabinarytreemid

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
func maxLevelSum(root *TreeNode) int {
	sum := []int{}
	var dfs func(node *TreeNode, dep int)
	dfs = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		if dep == len(sum) {
			sum = append(sum, node.Val)
		} else {
			sum[dep] += node.Val
		}
		dfs(node.Left, dep+1)
		dfs(node.Right, dep+1)
	}

	dfs(root, 0)
	idx := 0
	for i, s := range sum {
		if sum[idx] < s {
			idx = i
		}
	}
	return idx + 1
}
func Test_maximum_level_sum_of_a_binary_tree(t *testing.T) {
	arr := []int{1, 7, 0, 7, -8, -1, -1}
	t.Log(maxLevelSum(Init(arr)))
	arr = []int{989, -1, 10250, 98693, -89388, -1, -1, -1, -32127}
	t.Log(maxLevelSum(Init(arr)))
	arr = []int{39608, -1, -34332, 84856, 62101, 98129, -1, -1,
		-26118, -1, 57785, -75141, -1, -1, -63491, -63367}
	t.Log(maxLevelSum(Init(arr)))
}
