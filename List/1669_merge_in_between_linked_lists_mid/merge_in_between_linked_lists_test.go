package mergeinbetweenlinkedlists

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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p, q := list1, list1
	for ; a > 1; a-- {
		p = p.Next
	}
	for ; b > 0; b-- {
		q = q.Next
	}
	p.Next = list2
	for p.Next != nil {
		p = p.Next
	}
	p.Next = q.Next
	q.Next = nil
	return list1
}
func Test_merge_in_between_linked_lists(t *testing.T) {
	list1, a, b := Init([]int{0, 1, 2, 3, 4, 5}), 3, 4
	list2 := Init([]int{1000000, 1000001, 1000002})
	mergeInBetween(list1, a, b, list2).Show()
	list1, a, b = Init([]int{0, 1, 2, 3, 4, 5, 6}), 2, 5
	list2 = Init([]int{1000000, 1000001, 1000002, 1000003, 1000004})
	mergeInBetween(list1, a, b, list2).Show()
	list1, a, b = Init([]int{0, 3, 2, 1, 4, 5}), 3, 4
	list2 = Init([]int{1000000, 1000001, 1000002})
	mergeInBetween(list1, a, b, list2).Show()
}
