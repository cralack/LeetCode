package linkedlistcycle

import (
	"testing"

	. "LeetCode/util/MyList"
)

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
	list1 := Init([]int{3, 2, 0, -4})
	list1.Next.Next.Next.Next = list1.Next
	res1 := detectCycle(list1)
	t.Log(res1.Val)

	list2 := Init([]int{1, 2})
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
