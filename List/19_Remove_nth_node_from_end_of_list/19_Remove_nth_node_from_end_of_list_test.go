package removenthnodefromendoflist

import (
	. "Learning/LeetCode/List/MyList"
	"testing"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// dummy := &ListNode{0, head}
	dummy := &ListNode{Val: 0, Next: head}
	cur, pre := head, dummy
	for i := 0; i < n; i++ {
		cur = cur.Next
	}

	for ; cur != nil; cur = cur.Next {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next

	return dummy.Next
}

func TestRemoveNthNode(t *testing.T) {
	arrs := [][]int{{1, 2, 3, 4, 5}, {1}, {1, 2}}
	nn := []int{2, 1, 1}
	lists := make([]*ListNode, 0)
	for _, arr := range arrs {
		tmp := Init(arr)
		tmp.Show()
		lists = append(lists, tmp)
	}
	for idx, list := range lists {
		lists[idx] = removeNthFromEnd(list, nn[idx])
	}
	for _, list := range lists {
		list.Show()
	}
}
