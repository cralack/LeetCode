package reversenodesinkgroup

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListInit(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func (head *ListNode) Show() {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [left,right) 包含 k 个待反转元素
	left, right := head, head
	for i := 0; i < k; i++ {
		if right == nil {
			// 不足 k 个，不需要反转，base case
			return head
		}
		right = right.Next
	}
	// 反转前 k 个元素
	newHead := reverse(left, right)
	// 递归反转后续链表并连接起来
	left.Next = reverseKGroup(right, k)
	return newHead
}
func reverse(head, tail *ListNode) *ListNode {
	/* 反转区间 [head, tail) 的元素，注意是左闭右开 */
	var pre *ListNode = nil
	cur, next := head, head
	for cur != tail {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func Test_reverse_nodes_in_k_group(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := ListInit(arr)
	list = reverseKGroup(list, 2)
	list.Show()
}
