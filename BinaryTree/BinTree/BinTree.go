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
	if len(arr) == 0 {
		return nil
	}
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
	matrix := printTree(root)
	for i := range matrix {
		for _, char := range matrix[i] {
			fmt.Printf("%s", char)
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}

func printTree(root *TreeNode) [][]string {
	maxDep := -1
	//get maxDep
	var dfs_height func(*TreeNode, int)
	dfs_height = func(cur *TreeNode, dep int) {
		if cur == nil {
			return
		}
		if dep > maxDep {
			maxDep = dep
		}
		dfs_height(cur.Left, dep+1)
		dfs_height(cur.Right, dep+1)
	}
	dfs_height(root, 0)
	//init matrix
	m, n := maxDep+1, pow(2, maxDep+1)-1
	matrix := make([][]string, m)
	for i := range matrix {
		matrix[i] = make([]string, n)
		for j := range matrix[i] {
			matrix[i][j] = " "
		}
	}
	//get matrix
	var dfs_matrix func(*TreeNode, int, int)
	dfs_matrix = func(cur *TreeNode, r, c int) {
		if cur == nil {
			return
		}
		matrix[r][c] = strconv.Itoa(cur.Val)

		dfs_matrix(cur.Left, r+1, c-pow(2, maxDep-r-1))
		dfs_matrix(cur.Right, r+1, c+pow(2, maxDep-r-1))
	}
	dfs_matrix(root, 0, (n-1)/2)

	return matrix
}
func pow(x, n int) int {
	ans := 1
	for i := 0; i < n; i++ {
		ans *= x
	}
	return ans
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
