package addtwonumber

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Init(arr []int) {
	ln.Val = arr[0]
	cur := ln
	for i := 1; i < len(arr); i++ {
		tmp := &ListNode{Val: arr[i]}
		cur.Next = tmp
		cur = cur.Next
	}
}

func (ln *ListNode) Show() {
	cur := ln
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

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{Val: 0}
	cur := ans
	var sum, flag int
	for l1 != nil || l2 != nil {
		sum, flag = 0, 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += cur.Val
		flag = sum / 10
		cur.Val = sum % 10
		if l1 != nil || l2 != nil || flag != 0 {
			tmp := &ListNode{Val: flag}
			cur.Next = tmp
			cur = cur.Next
		}
	}
	return ans
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{}
	l2 := &ListNode{}
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}
	l1.Init(arr1)
	l2.Init(arr2)
	// l1.Show()
	// l2.Show()
	ret := AddTwoNumbers(l1, l2)
	ret.Show()
}
