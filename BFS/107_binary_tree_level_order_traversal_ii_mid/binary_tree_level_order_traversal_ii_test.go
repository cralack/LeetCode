package binarytreelevelordertraversalii

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
func levelOrderBottom(root *TreeNode) (ans [][]int) {
	if root == nil {
		return ans
	}
	que := []*TreeNode{root}
	dep := 0
	//bfs
	for len(que) > 0 {
		n := len(que)
		if len(ans) == dep {
			ans = append(ans, []int{})
		}
		for i := 0; i < n; i++ {
			cur := que[0]
			que = que[1:]
			ans[dep] = append(ans[dep], cur.Val)
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}
		dep++
	}
	//reverse
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}
	return
}

func Test_binary_tree_level_order_traversal_ii(t *testing.T) {
	root := Init([]int{3, 9, 20, -1, -1, 15, 7})
	t.Log(levelOrderBottom(root))
	root = Init([]int{1})
	t.Log(levelOrderBottom(root))
	root = Init([]int{})
	t.Log(levelOrderBottom(root))
}
