package constructbinarytreefrominorderandpostordertraversal

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}
	size := len(postorder)
	rootval, idx := postorder[size-1], -1
	for i, v := range inorder {
		if rootval == v {
			idx = i
		}
	}

	return &TreeNode{
		Val:   rootval,
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:size-1]),
	}
}
func Test_build_bin_tree_from_inorder_and_postorder_traversal(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	root.Show()
}
