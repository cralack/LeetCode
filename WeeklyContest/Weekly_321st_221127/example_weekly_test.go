package weekly_contest

import (
	. "LeetCode/util/MyList"
	"math"
	"testing"
)

/************ 1st test************/
func pivotInteger(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return x
	}
	return -1
}
func Test_1st(t *testing.T) {
	n := 8
	t.Log(pivotInteger(n)) //6
	n = 1
	t.Log(pivotInteger(n)) //1
	n = 4
	t.Log(pivotInteger(n)) //-1
}

/************ 2nd test************/
func appendCharacters(s string, t string) int {
	i, n := 0, len(s)
	for j := range t {
		for i < n && s[i] != t[j] {
			i++
		}
		if i == n {
			return len(t) - j
		}
		i++
	}
	return 0
}
func Test_2nd(t *testing.T) {
	s := "coaching"
	txt := "coding"
	t.Log(appendCharacters(s, txt))
}

/************ 3rd test************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	next := removeNodes(head.Next)
	if next == nil {
		return head
	}
	if head.Val < next.Val {
		return next
	}
	head.Next = next
	return head
}
func Test_3rd(t *testing.T) {
	head := Init([]int{5, 2, 13, 3, 8})
	head = removeNodes(head)
	head.Show()
}

/************ 4th test************/

func Test_4th(t *testing.T) {

}
