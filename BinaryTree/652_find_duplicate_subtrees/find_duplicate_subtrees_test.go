package findduplicatesubtrees

import (
	"fmt"
	"strconv"
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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var traverse func(root *TreeNode) string
	var memo = make(map[string]int)
	var res = make([]*TreeNode, 0)
	traverse = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		subtree := left + "," + right + "," + strconv.Itoa(root.Val)
		if memo[subtree] == 0 {
			memo[subtree]++
		} else if memo[subtree] == 1 {
			res = append(res, root)
			memo[subtree]++
		}
		return subtree
	}
	traverse(root)
	return res
}

func Test_find_duplicate_subtrees(t *testing.T) {
	// arr := []int{1, 2, 3, 4, -1, 2, 4, -1, -1, 4}
	// arr := []int{2, 1, 1}
	arr := []int{2, 2, 2, 3, -1, 3, -1}
	root := Init(arr)
	//root.Show()
	ans := findDuplicateSubtrees(root)
	for _, v := range ans {
		v.Show()
	}
}
