package reverselinkedlist

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
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
func ReverseListNode(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func Test_reverse_linked_list(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	list1 := &ListNode{}
	list1.Init(arr1)
	list2 := ReverseListNode(list1)
	list2.Show()
}
