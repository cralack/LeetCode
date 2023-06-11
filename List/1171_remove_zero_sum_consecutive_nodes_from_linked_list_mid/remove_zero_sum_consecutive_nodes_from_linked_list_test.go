package removezerosumconsecutivenodes

import (
	. "LeetCode/List/MyList"
	"testing"
)

/**
 * Definition for singly_linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	dum := &ListNode{Val: 0, Next: head}
	preSum := make(map[int]*ListNode)
	var cur *ListNode
	var sum int

	for cur, sum = dum, 0; cur != nil; cur = cur.Next {
		sum += cur.Val
		preSum[sum] = cur
	}
	for cur, sum = dum, 0; cur != nil; cur = cur.Next {
		sum += cur.Val
		cur.Next = preSum[sum].Next
	}
	return dum.Next
}
func Test_remove_zero_sum_consecutive_nodes_from_linked_list(t *testing.T) {
	tests := []struct {
		head *ListNode
	}{
		{head: Init([]int{1, 2, -3, 3, 1})},
		{head: Init([]int{1, 2, 3, -3, 4})},
		{head: Init([]int{1, 2, 3, -3, -2})},
	}

	for _, tt := range tests {
		removeZeroSumSublists(tt.head).Show()
	}
}
