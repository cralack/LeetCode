package linkedlistcomponents

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
func numComponents(head *ListNode, nums []int) (ans int) {
	vis := make(map[int]struct{}, 0)
	for _, n := range nums {
		vis[n] = struct{}{}
	}

	flag := false
	for cur := head; cur != nil; cur = cur.Next {
		if _, has := vis[cur.Val]; has && !flag {
			ans++
			flag = true
		} else if !has {
			flag = false
		}
	}
	return
}

func Test_linked_list_components(t *testing.T) {
	head := Init([]int{0, 1, 2, 3})
	nums := []int{0, 1, 3}
	t.Log(numComponents(head, nums))
	head = Init([]int{0, 1, 2, 3, 4})
	nums = []int{0, 3, 1, 4}
	t.Log(numComponents(head, nums))
	head = Init([]int{1, 2, 0, 4, 3})
	nums = []int{3, 4, 0, 2, 1}
	t.Log(numComponents(head, nums))
}
