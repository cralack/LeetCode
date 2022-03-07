package convertbsttogreatertree

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
	queque := []*TreeNode{root}
	for i := 1; i < len(arr); i += 2 {
		cur := queque[0]
		queque = queque[1:]
		if arr[i] != -1 {
			cur.Left = &TreeNode{Val: arr[i]}
			queque = append(queque, cur.Left)
		}
		if i+1 < len(arr) && arr[i+1] != -1 {
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

//中序遍历递归写法
func ConvertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}
	traverse(root)
	return root
}

//中序遍历迭代写法
func ConverToBST(root *TreeNode) *TreeNode {
	curMax := 0
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curMax += cur.Val
			cur.Val = curMax
			cur = cur.Left
		}
	}
	return root
}
func Test_convert_bst_to_greater_tree(t *testing.T) {
	arr1 := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	root := Init(arr1)
	// greater := convertBST(root)
	greater := ConverToBST(root)
	greater.Show()
}
