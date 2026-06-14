package maximum_twin_sum_of_a_linked_list_mid

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
func pairSum(head *ListNode) (ans int) {
	left := head

	var dfs func(*ListNode)
	dfs = func(right *ListNode) {
		if right.Next != nil {
			dfs(right.Next)
		}
		ans = max(ans, left.Val+right.Val)
		left = left.Next
	}
	dfs(head)
	return
}

func Test_maximum_twin_sum_of_a_linked_list(t *testing.T) {
	tests := []struct {
		head *ListNode
	}{
		{head: Init([]int{5, 4, 2, 1})},
		{head: Init([]int{4, 2, 2, 3})},
		{head: Init([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
		{head: Init([]int{1, 100000})},
	}
	for _, test := range tests {
		t.Log(pairSum(test.head))
	}
}
