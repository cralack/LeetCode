package delete_nodes_from_linked_list_present_in_array_mid

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
func modifiedList(nums []int, head *ListNode) *ListNode {
	var dummy = &ListNode{Val: -1, Next: head}
	dic := map[int]bool{}
	for _, n := range nums {
		dic[n] = true
	}
	pre, cur := dummy, dummy.Next
	for cur != nil {
		if dic[cur.Val] {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}

func Test_delete_nodes_from_linked_list_present_in_array(t *testing.T) {
	tests := []struct {
		nums []int
		head []int
	}{
		{nums: []int{1, 2, 3}, head: []int{1, 2, 3, 4, 5}},
		{nums: []int{1}, head: []int{1, 2, 1, 2, 1, 2}},
		{nums: []int{5}, head: []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		modifiedList(test.nums, Init(test.head)).Show()
	}
}
