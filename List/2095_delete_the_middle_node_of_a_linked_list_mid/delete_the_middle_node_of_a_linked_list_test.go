package delete_the_middle_node_of_a_linked_list_mid

import (
	"testing"

	. "LeetCode/util/MyList"
)

/**
 * Definition for singly_linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func Test_delete_the_middle_node_of_a_linked_list(t *testing.T) {
	tests := []struct {
		head *ListNode
	}{
		{head: Init([]int{1, 3, 4, 7, 1, 2, 6})},
		{head: Init([]int{1, 2, 3, 4})},
		{head: Init([]int{2, 1})},
		{head: Init([]int{1})},
	}
	for _, test := range tests {
		deleteMiddle(test.head).Show()
	}
}
