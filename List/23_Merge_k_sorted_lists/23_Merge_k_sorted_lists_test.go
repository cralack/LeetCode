package mergeksortedlists

import (
	"container/heap"
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Init(arr []int) {
	cur := head
	cur.Val = arr[0]
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
}

func (head *ListNode) Show() {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}

type minHeap []*ListNode

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	size := old.Len()
	x := old[size-1]
	*h = old[0 : size-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := new(minHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}
	return dummyHead.Next
}

func TestMergeList(t *testing.T) {
	lists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	listz := make([]*ListNode, 0)
	for _, list := range lists {
		tmp := &ListNode{}
		tmp.Init(list)
		listz = append(listz, tmp)
	}
	res := mergeKLists(listz)
	res.Show()
}
