package removeduplicatesfromsortedlist

import (
	. "Learning/LeetCode/List/MyList"
	"testing"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
func Test_remove_duplicates_from_sorted_list(t *testing.T) {
	arr := []int{1, 1, 2, 3, 3}
	head := &ListNode{}
	head.Init(arr)
	deleteDuplicates(head)
	head.Show()
}
