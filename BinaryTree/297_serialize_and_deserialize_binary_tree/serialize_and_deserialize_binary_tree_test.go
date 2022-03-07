package serializeanddeserializebinarytree

import (
	. "Learning/Leetcode/BinaryTree/BinTree"
	"strconv"
	"strings"
	"testing"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}
func (cod *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := &strings.Builder{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		//stack.pop
		cur := queue[0]
		queue = queue[1:]
		if cur != nil {
			res.WriteString(strconv.Itoa(cur.Val))
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)

		} else {
			res.WriteString("#")
		}
		if len(queue) > 0 {
			res.WriteByte(',')
		}
	}
	return res.String()
}
func (cod *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 0 {
		return nil
	}
	str := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(str[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	for i := 1; i < len(str); i += 2 {
		cur := queue[0]
		queue = queue[1:]
		if str[i] != "#" {
			leftVal, _ := strconv.Atoi(str[i])
			cur.Left = &TreeNode{Val: leftVal}
			queue = append(queue, cur.Left)
		}
		if i+1 < len(str) && str[i+1] != "#" {
			rightVal, _ := strconv.Atoi(str[i+1])
			cur.Right = &TreeNode{Val: rightVal}
			queue = append(queue, cur.Right)
		}
	}
	return root
}

func serialize(root *TreeNode) string {
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
func deserialize(data string) *TreeNode {
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

// func Test_serialize_and_deserialize_binary_tree(t *testing.T) {
// 	arr := []byte{1, 4, 3, 2, 4, 2, 5, '#', '#', '#', '#', '#', '#', 4, 6}
// 	rootSrc := ByteInit(arr)
// 	code := &Codec{}
// 	str1 := code.serialize(rootSrc)
// 	rootDes1 := code.deserialize(str1)
// 	// root2.Show()
// 	t.Log(Equals(rootSrc, rootDes1))

// 	str2 := serialize(rootSrc)
// 	rootDes2 := deserialize(str2)
// 	t.Log(Equals(rootDes2, rootSrc))
// }

func Benchmark_serialize_and_deserialize_binary_tree(b *testing.B) {
	des := []byte{1, 4, 3, 2, 4, 2, 5, '#', '#', '#', '#', '#', '#', 4, 6}
	rootDes := ByteInit(des)
	b.Run("迭代", func(b *testing.B) {
		b.ResetTimer()
		code := &Codec{}
		for i := 0; i < b.N; i++ {
			str := code.serialize(rootDes)
			rootSrc := code.deserialize(str)
			Equals(rootDes, rootSrc)
		}
		b.StopTimer()
	})
	b.Run("递归", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			str := serialize(rootDes)
			rootSrc := deserialize(str)
			Equals(rootDes, rootSrc)
		}
	})
}
