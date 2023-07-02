package add_two_numbers

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	ans = &ListNode{}
	cur, carry := ans, 0
	for ; l1 != nil || l2 != nil || carry > 0; cur = cur.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}
	return ans.Next
}
func Test_add_two_numbers(t *testing.T) {
	tests := []struct {
		l1 *ListNode
		l2 *ListNode
	}{
		{Init([]int{2, 4, 3}), Init([]int{5, 6, 4})},
		{Init([]int{0}), Init([]int{0})},
		{Init([]int{9, 9, 9, 9, 9, 9, 9}), Init([]int{9, 9, 9, 9})},
	}
	for _, tt := range tests {
		addTwoNumbers(tt.l1, tt.l2).Show()
	}
}
