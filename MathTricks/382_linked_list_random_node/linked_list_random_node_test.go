package linkedlistrandomnode

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	. "LeetCode/util/MyList"
)

/**
 * Definition for singly_linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (p *Solution) GetRandom() (ans int) {
	for cur, i := p.Head, 1; cur != nil; cur = cur.Next {
		if rand.Intn(i) == 0 {
			ans = cur.Val
		}
		i++
	}
	return
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
func Test_linked_list_random_node(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}
	head := Init(arr)
	sol := Solution{Head: head}
	waiter := sync.WaitGroup{}
	n := 20

	waiter.Add(n)
	cnt := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 500; j++ {
				cnt[sol.GetRandom()]++
			}
			waiter.Done()
		}()

		time.Sleep(1 * time.Nanosecond)
	}
	waiter.Wait()
	weight := make([]float32, 10)
	for i := range cnt {
		weight[i-1] = float32(cnt[i]) / float32(1000)
	}
	t.Log(weight)
}
