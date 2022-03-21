package twosumivinputisabst

import (
	. "Learning/Leetcode/BinaryTree/BinTree"
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
func findTarget(root *TreeNode, k int) bool {
	vis := make(map[int]bool, 0)
	que := []*TreeNode{root}
	for len(que) > 0 {
		//pop
		cur := que[0]
		que = que[1:]
		if _, ok := vis[k-cur.Val]; ok {
			return true
		}
		vis[cur.Val] = true
		if cur.Left != nil {
			que = append(que, cur.Left)
		}
		if cur.Right != nil {
			que = append(que, cur.Right)
		}
	}
	return false
}
func Test_two_sum_iv_input_is_a_bst(t *testing.T) {
	arr := []int{5, 3, 6, 2, 4, -1, 7}
	root := Init(arr)
	t.Log(findTarget(root, 9))
	t.Log(findTarget(root, 28))
}
