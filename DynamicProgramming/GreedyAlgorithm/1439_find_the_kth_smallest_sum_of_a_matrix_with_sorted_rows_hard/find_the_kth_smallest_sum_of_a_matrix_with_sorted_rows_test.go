package findthekthsmallestsum

import (
	"container/heap"
	"testing"
)

// 373. 查找和最小的 K 对数字
func kSmallestPairs(nums1, nums2 []int, k int) []int {
	n, m := len(nums1), len(nums2)
	ans := make([]int, 0, min(k, n*m)) // 预分配空间
	h := hp{{nums1[0] + nums2[0], 0, 0}}
	for len(h) > 0 && len(ans) < k {
		cur := heap.Pop(&h).(tuple)
		i, j := cur.i, cur.j
		ans = append(ans, nums1[i]+nums2[j]) // 数对和
		if j == 0 && i+1 < n {
			heap.Push(&h, tuple{nums1[i+1] + nums2[0], i + 1, 0})
		}
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

func kthSmallest(mat [][]int, k int) int {
	sum := []int{0}
	for _, row := range mat {
		// 每一行与之前的k个最小sum
		sum = kSmallestPairs(row, sum, k)
	}
	return sum[k-1]
}

type tuple struct{ sum, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func Test_find_the_kth_smallest_sum_of_a_matrix_with_sorted_rows(t *testing.T) {
	tests := []struct {
		mat [][]int
		k   int
	}{
		{mat: [][]int{{1, 3, 11}, {2, 4, 6}}, k: 5},
		{mat: [][]int{{1, 3, 11}, {2, 4, 6}}, k: 9},
		{mat: [][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, k: 7},
		{mat: [][]int{{1, 1, 10}, {2, 2, 9}}, k: 7},
	}
	for _, tt := range tests {
		t.Log(kthSmallest(tt.mat, tt.k))
	}
}
