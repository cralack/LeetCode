package maximum_total_subarray_value_ii_mid

import (
	"container/heap"
	"math/bits"
	"testing"
)

type pair struct{ min, max int }

func op(a, b pair) pair {
	return pair{min(a.min, b.min), max(a.max, b.max)}
}

type ST [][16]pair

func newST(a []int) ST {
	n := len(a)
	w := bits.Len(uint(n))
	st := make(ST, n)
	for i, x := range a {
		st[i][0] = pair{x, x}
	}
	for j := 1; j < w; j++ {
		for i := range n - 1<<j + 1 {
			st[i][j] = op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}
func (st ST) query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	p := op(st[l][k], st[r-1<<k][k])
	return p.max - p.min
}

func maxTotalValue(nums []int, k int) (ans int64) {
	n := len(nums)
	st := newST(nums)
	h := make(hp, n)
	for i := range h {
		h[i] = tuple{st.query(i, n), i, n}
	}

	for ; k > 0 && h[0].diff > 0; k-- {
		ans += int64(h[0].diff)
		h[0].r--
		h[0].diff = st.query(h[0].l, h[0].r)
		heap.Fix(&h, 0)
	}
	return
}

type tuple struct{ diff, l, r int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].diff > h[j].diff }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func Test_maximum_total_subarray_value_ii(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
	}{
		{nums: []int{1, 3, 2}, k: 2},
		{nums: []int{4, 2, 5, 1}, k: 3},
	}
	for _, tt := range tests {
		t.Log(maxTotalValue(tt.nums, tt.k))
	}
}
