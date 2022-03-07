package linkedlistcycle

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

func hasCycle(head *ListNode) bool {
	pre, cur := head, head
	for cur != nil && cur.Next != nil {
		pre = pre.Next
		cur = cur.Next.Next
		if pre == cur {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	pre, cur := head, head
	for cur != nil && cur.Next != nil {
		cur = cur.Next.Next
		pre = pre.Next
		if cur == pre {
			break
		}
	}
	if cur == nil || cur.Next == nil {
		return nil
	}
	pre = head
	for cur != pre {
		cur = cur.Next
		pre = pre.Next
	}
	return pre
}

func TestListCycle(t *testing.T) {
	arr1 := []int{3, 2, 0, -4}
	list1 := new(ListNode)
	list1.Init(arr1)
	list1.Next.Next.Next.Next = list1.Next
	res1 := detectCycle(list1)
	t.Log(res1.Val)

	arr2 := []int{1, 2}
	list2 := new(ListNode)
	list2.Init(arr2)
	list2.Next.Next = list2
	res2 := detectCycle(list2)
	t.Log(res2.Val)

	list3 := &ListNode{Val: 1}
	if hasCycle(list3) {
		t.Log("yes")
	} else {
		t.Log("no")
	}
}
