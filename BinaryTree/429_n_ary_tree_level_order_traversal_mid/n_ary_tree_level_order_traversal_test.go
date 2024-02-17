package n_ary_tree_level_order_traversal_mid

import (
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func Deserialize(arr []interface{}) *Node {
	if len(arr) == 0 {
		return nil
	}

	root := &Node{Val: arr[0].(int)} // assuming the first element is always the root and is an int
	queue := []*Node{root}
	index := 2 // starting after the root and its null separator

	for index < len(arr) {
		parent := queue[0]
		queue = queue[1:]
		var children []*Node
		for index < len(arr) && arr[index] != nil {
			child := &Node{Val: arr[index].(int)}
			children = append(children, child)
			queue = append(queue, child)
			index++
		}
		parent.Children = children
		index++ // skip the null separator
	}

	return root
}

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	que := []*Node{root}
	for len(que) > 0 {
		n := len(que)
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			cur := que[0]
			que = que[1:]

			tmp[i] = cur.Val

			for _, next := range cur.Children {
				que = append(que, next)
			}
		}
		ans = append(ans, tmp)
	}
	return
}

func Test_n_ary_tree_level_order_traversal(t *testing.T) {
	root1 := Deserialize([]interface{}{1, nil, 3, 2, 4, nil, 5, 6})
	t.Log(levelOrder(root1))
	root2 := Deserialize([]interface{}{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14})
	t.Log(levelOrder(root2))
}
