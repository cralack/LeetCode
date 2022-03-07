package reverselinkedlistii

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
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// var successor *ListNode = nil

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		return head
	}
	last := reverseN(head.Next, n-1)
	headNext := head.Next
	head.Next = headNext.Next
	headNext.Next = head
	return last
}

func Test_reverse_linked_list_ii(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	left, right := 2, 4
	list1 := &ListNode{}
	list1.Init(arr1)
	list2 := reverseBetween(list1, left, right)
	// list2 := reverseN(list1, 3)
	list2.Show()
}
