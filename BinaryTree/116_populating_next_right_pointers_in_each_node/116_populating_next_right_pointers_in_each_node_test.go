package populatingnextrightpointersineachnode

import (
	"fmt"
	"testing"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Init(arr []int) *Node {
	root := &Node{Val: arr[0]}
	queque := make([]*Node, 0)
	queque = append(queque, root)
	for i := 1; i < len(arr); i += 2 {
		cur := queque[0]
		queque = queque[1:]
		if arr[i] != -1 {
			cur.Left = &Node{Val: arr[i]}
			queque = append(queque, cur.Left)
		}
		if arr[i+1] != -1 {
			cur.Right = &Node{Val: arr[i+1]}
			queque = append(queque, cur.Right)
		}
	}
	return root
}

func (root *Node) Show() {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	dep := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", cur.Val)
			if cur.Next == nil {
				fmt.Println("#")
			}
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

// func connect(root *Node) *Node {
// 	var connectTwoNode func(node1, node2 *Node)
// 	connectTwoNode = func(node1, node2 *Node) {
// 		if node1 == nil || node2 == nil {
// 			return
// 		}
// 		node1.Next = node2
// 		connectTwoNode(node1.Left, node1.Right)
// 		connectTwoNode(node2.Left, node2.Right)
// 		connectTwoNode(node1.Right, node2.Left)
// 	}
// 	if root == nil {
// 		return nil
// 	}
// 	connectTwoNode(root.Left, root.Right)
// 	return root
// }

// level-traverse
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

func Test_populating_next_right_pointers_in_each_node(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	list := Init(arr)
	list.Show()
	connect(list)
	list.Show()
}
