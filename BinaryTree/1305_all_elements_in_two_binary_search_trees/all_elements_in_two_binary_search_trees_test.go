package allelementsintwobinarysearchtrees

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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := inorder(root1)
	arr2 := inorder(root2)

	n, m := len(arr1), len(arr2)
	i, j := 0, 0
	res := make([]int, 0, m+n+1)
	for i <= n && j <= m {
		if j == m {
			return append(res, arr1[i:]...)
		}
		if i == n {
			return append(res, arr2[j:]...)
		}
		if arr1[i] <= arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	return res
}
func inorder(root *TreeNode) (res []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}
func Test_all_elements_in_two_binary_search_trees(t *testing.T) {
	root1 := Init([]int{2, 1, 4})
	root2 := Init([]int{1, 0, 3})
	t.Log(getAllElements(root1, root2))
	root1 = Init([]int{1, -1, 8})
	root2 = Init([]int{8, 1})
	t.Log(getAllElements(root1, root2))
}
