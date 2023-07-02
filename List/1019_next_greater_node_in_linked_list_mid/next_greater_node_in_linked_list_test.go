package nextgreaternodeinlinkedlist

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
func nextLargerNodes(head *ListNode) (ans []int) {
	stack := make([]int, 0)
	var dfs func(int, *ListNode)
	dfs = func(idx int, cur *ListNode) {
		if cur == nil {
			ans = make([]int, idx)
			return
		}
		dfs(idx+1, cur.Next)
		for len(stack) > 0 && stack[len(stack)-1] <= cur.Val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[idx] = stack[len(stack)-1]
		}
		stack = append(stack, cur.Val)
	}
	dfs(0, head)
	return
}
func Test_next_greater_node_in_linked_list(t *testing.T) {
	head := Init([]int{2, 1, 5})
	t.Log(nextLargerNodes(head))
	head = Init([]int{2, 7, 4, 3, 5})
	t.Log(nextLargerNodes(head))
	head = Init([]int{2, 7, 2, 4, 3, 5})
	t.Log(nextLargerNodes(head))
}
