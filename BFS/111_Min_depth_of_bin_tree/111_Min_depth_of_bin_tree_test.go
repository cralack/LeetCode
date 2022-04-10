package mindepthofbintree

import (
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	dep := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				return dep
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		dep++
	}
	return dep
}

func TestMinDepth(t *testing.T) {
	arr := []int{3, 9, 20, -1, -1, 15, 7}
	root := Init(arr)
	dep := minDepth(root)
	t.Log(dep)

	arr1 := []int{2, -1, 3, -1, 4, -1, 5, -1, 6}
	root2 := Init(arr1)
	dep2 := minDepth(root2)
	t.Log(dep2)
}
