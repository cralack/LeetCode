package maximumbinarytree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	idx, max := -1, -1
	for i, v := range nums {
		if v > max {
			max = v
			idx = i
		}
	}
	root := &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
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

func Test_maximum_binary_tree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	root.Show()
}
