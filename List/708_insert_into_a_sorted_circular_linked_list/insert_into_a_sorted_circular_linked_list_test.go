package insertintoasortedcircularlinkedlist

import (
	"fmt"
	"testing"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

type Node struct {
	Val  int
	Next *Node
}

func Init(arr []int) *Node {
	head := &Node{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &Node{Val: arr[i]}
		cur = cur.Next
	}
	cur.Next = head

	return head
}
func (head *Node) Show() {
	cur := head
	for cur.Next.Val != head.Val {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("%d ", cur.Val)
	fmt.Printf("\n")
}

func insert(head *Node, x int) *Node {
	t := &Node{Val: x}
	t.Next = t
	if head == nil {
		return t
	}
	if head.Next == head {
		head.Next = t
		t.Next = head
		return head
	}
	cur, next := head, head.Next
	for next != head {
		if cur.Val <= x && x <= next.Val {
			break
		}
		if cur.Val > next.Val {
			if cur.Val < x || x < next.Val {
				break
			}
		}
		cur = cur.Next
		next = next.Next
	}
	cur.Next = t
	t.Next = next
	return head
}
func Test_insert_into_a_sorted_circular_linked_list(t *testing.T) {
	arr, insertVal := []int{3, 4, 1}, 2
	head := Init(arr)
	head.Show()
	insert(head, insertVal)
	head.Show()
}
