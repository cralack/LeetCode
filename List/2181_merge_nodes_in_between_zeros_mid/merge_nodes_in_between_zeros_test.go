package merge_nodes_in_between_zeros_mid

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
func mergeNodes(head *ListNode) *ListNode {
	tail := head
	for cur := head.Next; cur.Next != nil; cur = cur.Next {
		if cur.Val != 0 {
			tail.Val += cur.Val
		} else {
			tail = tail.Next
			tail.Val = 0
		}
	}
	tail.Next = nil
	return head
}

func Test_merge_nodes_in_between_zeros(t *testing.T) {
	tests := []struct {
		head1 *ListNode
	}{
		{Init([]int{0, 3, 1, 0, 4, 5, 2, 0})},
		{Init([]int{0, 1, 0, 3, 0, 2, 2, 0})},
	}
	for _, tt := range tests {
		mergeNodes(tt.head1)
		tt.head1.Show()
	}
}
