package mergetwosortedlist

import (
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2
	ans := &ListNode{Val: -1}
	cur := ans
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
	}
	if p2 != nil {
		cur.Next = p2
	}
	return ans.Next
}

func (ln *ListNode) Init(arr []int) {
	cur := ln
	cur.Val = arr[0]
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
}

func TestMergeList(t *testing.T) {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	ln1 := &ListNode{}
	ln2 := &ListNode{}
	ln1.Init(arr1)
	ln2.Init(arr2)
	ans := MergeTwoLists(ln1, ln2)
	t.Log(ans)
}
