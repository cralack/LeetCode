package kthsmallestelementinabst

import "testing"

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

func kthSmallest(root *TreeNode, k int) int {
	idx := 0
	kVal := -1
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		idx++
		if idx == k {
			kVal = root.Val
			return
		}
		traverse(root.Right)
	}
	traverse(root)
	return kVal
}
func Test_kth_smallest_element_in_a_bst(t *testing.T) {
	// arr := []int{3, 1, 4, -1, 2}
	arr := []int{5, 3, 6, 2, 4, -1, -1, 1}
	root := Init(arr)
	k := kthSmallest(root, 3)
	t.Log(k)
}
