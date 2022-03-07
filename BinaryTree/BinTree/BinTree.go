package Bintree

import (
	"fmt"
	"strconv"
	"strings"
)

/*
import (
	. "Learning/Leetcode/BinaryTree/BinTree"
	"testing"
)
*/

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

func ByteInit(arr []byte) *TreeNode {
	root := &TreeNode{Val: int(arr[0])}
	queque := []*TreeNode{root}
	for i := 1; i < len(arr); i += 2 {
		cur := queque[0]
		queque = queque[1:]
		if arr[i] != '#' {
			cur.Left = &TreeNode{Val: int(arr[i])}
			queque = append(queque, cur.Left)
		}
		if i+1 < len(arr) && arr[i+1] != '#' {
			cur.Right = &TreeNode{Val: int(arr[i+1])}
			queque = append(queque, cur.Right)
		}
	}
	return root
}
func (root *TreeNode) Show() {
	if root == nil {
		fmt.Println("nil")
		return
	}
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

func (root *TreeNode) Find(n int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == n {
		return root
	}
	left := root.Left.Find(n)
	if left != nil {
		return left
	}
	right := root.Right.Find(n)
	if right != nil {
		return right
	}
	return nil
}

func Equals(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		if Equals(a.Left, b.Left) && Equals(a.Right, b.Right) {
			return true
		}
	}
	return false
}

func Serialize(root *TreeNode) string {
	res := &strings.Builder{}
	var traverse func(node *TreeNode, buf strings.Builder)
	traverse = func(node *TreeNode, buf strings.Builder) {
		if node == nil {
			res.WriteString("#")
			res.WriteByte(',')
			return
		}
		res.WriteString(strconv.Itoa(node.Val))
		res.WriteByte(',')
		traverse(node.Left, *res)
		traverse(node.Right, *res)
	}
	traverse(root, *res)
	return res.String()
}
func Deserialize(data string) *TreeNode {
	str := strings.Split(data, ",")
	var traverse func() *TreeNode
	traverse = func() *TreeNode {
		if str[0] == "#" {
			str = str[1:]
			return nil
		}
		val, _ := strconv.Atoi(str[0])
		str = str[1:]
		return &TreeNode{Val: val, Left: traverse(), Right: traverse()}
	}
	return traverse()
}
