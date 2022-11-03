package serializeanddeserializebst

import (
	. "LeetCode/BinaryTree/BinTree"
	"strconv"
	"strings"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	Code string
	Sep  byte
	Root *TreeNode
}

func Constructor() Codec {
	return Codec{
		Sep: ',',
	}
}

// Serializes a tree to a single string.
func (p *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			sb.WriteString("#")
			sb.WriteByte(p.Sep)
			return
		}
		sb.WriteString(strconv.Itoa(cur.Val))
		sb.WriteByte(p.Sep)
		dfs(cur.Left)
		dfs(cur.Right)
	}
	dfs(root)
	p.Code = sb.String()
	return p.Code
}

// Deserializes your encoded data to tree.
func (p *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, string(p.Sep))
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if 0 == len(nodes) {
			return nil
		}
		cur := nodes[0]
		nodes = nodes[1:]
		if cur == "#" {
			return nil
		}
		n, _ := strconv.Atoi(cur)
		root := &TreeNode{Val: n}
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	p.Root = dfs()
	return p.Root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
func Test_serialize_and_deserialize_bst(t *testing.T) {
	root := Init([]int{2, 1, 3, -1, 3, 4, 5})
	codec := Constructor()
	t.Log(codec.serialize(root))
	codec.deserialize(codec.Code)
	root = Init([]int{1, 2, 3, -1, 4})
	t.Log(codec.serialize(root))
	root = Init([]int{2, 1, 3})
	t.Log(codec.serialize(root))
}
