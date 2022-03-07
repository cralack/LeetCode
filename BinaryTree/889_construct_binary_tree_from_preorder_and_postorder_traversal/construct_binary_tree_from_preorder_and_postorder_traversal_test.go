package constructbinarytreefrompreorderandpostordertraversal

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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	for i, v := range postorder {
		if preorder[1] == v {
			root.Left = constructFromPrePost(preorder[1:i+2],
				postorder[:i+1])
			root.Right = constructFromPrePost(preorder[i+2:],
				postorder[i+1:len(postorder)-1])
			break
		}
	}
	return root
}
func Test_build_bin_tree_from_preorder_and_postorder_traversal(t *testing.T) {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	// pre := []int{2, 1, 3}
	// post := []int{3, 1, 2}
	root := constructFromPrePost(pre, post)
	root.Show()
}
