package successorlcci

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
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var succ *TreeNode
	//如果节点 pp 的右子树不为空
	//则节点 pp 的后继节点为其右子树中最左侧节点
	if p.Right != nil {
		succ = p.Right
		for succ.Left != nil {
			succ = succ.Left
		}
		return succ
	}
	cur := root
	for cur != nil {
		if cur.Val > p.Val {
			succ = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return succ
}
func Test_successor_lcci(t *testing.T) {
	root := Init([]int{2, 1, 3})
	p := &TreeNode{Val: 1}
	t.Log(inorderSuccessor(root, p).Val)
	root = Init([]int{5, 3, 6, 2, 4, -1, -1, 1})
	p = &TreeNode{Val: 6}
	if inorderSuccessor(root, p) == nil {
		t.Log("nil")
	}
}
