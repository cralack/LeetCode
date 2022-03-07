package middleofthelinkedlist

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Init(arr []int) {
	cur := head
	cur.Val = arr[0]
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
}

func (head *ListNode) Show() {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}

func MiddleNode(head *ListNode) *ListNode {
	pre, cur := head, head
	for cur != nil && cur.Next != nil {
		pre = pre.Next
		cur = cur.Next.Next
	}
	return pre
}

func TestMidNode(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	list1 := new(ListNode)
	list1.Init(arr1)
	node1 := MiddleNode(list1)
	t.Log(node1.Val)

	arr2 := []int{1, 2, 3, 4, 5, 6}
	list2 := new(ListNode)
	list2.Init(arr2)
	node2 := MiddleNode(list2)
	t.Log(node2.Val)
}
