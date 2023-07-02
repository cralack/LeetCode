package intersectionoftwolinkedlists

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	// 拼接后对齐
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
func TestIntersectionOfTwoLinkedList(t *testing.T) {
	arr1 := []int{4, 1}
	arr2 := []int{5, 6, 1, 8, 4, 5}
	list1 := new(ListNode)
	list1.Init(arr1)
	list2 := new(ListNode)
	list2.Init(arr2)
	list1.Next.Next = list2.Next.Next.Next
	res := getIntersectionNode(list1, list2)
	t.Log(res.Val)
}
