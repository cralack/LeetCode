package palindromelinkedlist

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

func isPalindrome(head *ListNode) bool {
	reverse := func(head *ListNode) *ListNode {
		var pre *ListNode = nil
		cur := head
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
func Test_palindrome_linked_list(t *testing.T) {
	arr := [][]int{{1, 2, 2, 1}, {1, 2}}
	for _, val := range arr {
		list := ListInit(val)
		if isPalindrome(list) {
			t.Log("true")
		} else {
			t.Log("false")
		}
	}

}
