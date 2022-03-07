package invertbinarytree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Init(arr []int) *TreeNode {
	root := &TreeNode{Val: arr[0]}
	queque := make([]*TreeNode, 0)
	queque = append(queque, root)
	for i := 1; i < len(arr); i += 2 {
		cur := queque[0]
		queque = queque[1:]
		if arr[i] != -1 {
			cur.Left = &TreeNode{Val: arr[i]}
			queque = append(queque, cur.Left)
		}
		if arr[i+1] != -1 {
			cur.Right = &TreeNode{Val: arr[i+1]}
			queque = append(queque, cur.Right)
		}
	}
	return root
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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func Test_invert_binary_tree(t *testing.T) {
	arr := []int{4, 2, 7, 1, 3, 6, 9}
	root := Init(arr)
	invertTree(root)
	root.Show()
}
