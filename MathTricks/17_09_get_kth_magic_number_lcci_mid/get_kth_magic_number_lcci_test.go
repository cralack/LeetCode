package getkthmagicnumberlccimid

import (
	"container/heap"
	"testing"
)

func getKthMagicNumber_heap(k int) int {
	nums := []int{3, 5, 7}
	pq := &PriorityQ{}
	set := map[int]bool{}
	heap.Push(pq, 1)
	set[1] = true
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(int)
		k--
		if k == 0 {
			return cur
		}
		for _, x := range nums {
			if _, has := set[cur*x]; !has {
				heap.Push(pq, cur*x)
				set[cur*x] = true
			}
		}
	}
	return -1
}

type PriorityQ []int

func (h PriorityQ) Len() int            { return len(h) }
func (h PriorityQ) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h PriorityQ) Less(i, j int) bool  { return h[i] < h[j] }
func (h *PriorityQ) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *PriorityQ) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func getKthMagicNumber(k int) int {
	ans := make([]int, k+1)
	ans[1] = 1
	i3, i5, i7 := 1, 1, 1
	for idx := 2; idx <= k; idx++ {
		a := ans[i3] * 3
		b := ans[i5] * 5
		c := ans[i7] * 7
		mi := min(a, min(b, c))
		if mi == a {
			i3++
		}
		if mi == b {
			i5++
		}
		if mi == c {
			i7++
		}
		ans[idx] = mi
	}
	return ans[k]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_get_kth_magic_number_lcci(t *testing.T) {
	k := 5
	t.Log(getKthMagicNumber(k))
	k = 7
	t.Log(getKthMagicNumber(k))
}
