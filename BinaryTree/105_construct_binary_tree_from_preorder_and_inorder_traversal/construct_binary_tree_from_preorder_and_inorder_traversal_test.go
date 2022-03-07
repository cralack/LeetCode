package constructbinarytreefrompreorderandinordertraversal

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) Show() {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	dep := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		dep++
		fmt.Printf("\r\n")
	}
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}
	rootval, idx := preorder[0], -1
	for i, val := range inorder {
		if rootval == val {
			idx = i
		}
	}
	root := &TreeNode{
		Val:   rootval,
		Left:  buildTree(preorder[1:1+idx], inorder[:idx]),
		Right: buildTree(preorder[1+idx:], inorder[idx+1:]),
	}
	return root
}
func Test_build_bin_tree_from_preorder_and_inorder_traversal(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	root.Show()
}
