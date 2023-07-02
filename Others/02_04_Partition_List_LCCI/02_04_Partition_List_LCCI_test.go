package partitionlistlcci

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type List struct {
	Head *ListNode
}

func (ln *List) Init(arr []int) {
	ln.Head = &ListNode{Val: arr[0]}
	cur := ln.Head
	for i := 1; i < len(arr); i++ {
		tmp := &ListNode{Val: arr[i]}
		cur.Next = tmp
		cur = cur.Next
	}
}

func (ln *List) Show() {
	cur := ln.Head
	if cur == nil {
		fmt.Println("list empty")
		return
	}
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}

func patition(head *ListNode, x int) *ListNode {
	cur := head
	pLeft := &ListNode{}
	LeftHead := pLeft
	pRight := &ListNode{}
	RightHead := pRight

	for cur != nil {
		if cur.Val < x {
			pLeft.Next = cur
			pLeft = pLeft.Next
		} else {
			pRight.Next = cur
			pRight = pRight.Next
		}
		cur = cur.Next
	}

	pLeft.Next = RightHead.Next
	pRight.Next = nil
	head = LeftHead.Next
	return head
}

func TestListNode(t *testing.T) {
	arr := [...]int{1, 4, 3, 2, 5, 2}
	mylist := &List{}
	mylist.Init(arr[:])
	// mylist.Show()
	patition(mylist.Head, 3)
	// mylist.Show()
}
