package path_existence_queries_in_a_graph_ii_hard

import (
	"math/bits"
	"slices"
	"testing"
)

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return nums[i] - nums[j] })

	// rank[i] 表示 nums[i] 是 nums 中的第几小，或者说节点 i 在 idx 中的下标
	rank := make([]int, n)
	for i, j := range idx {
		rank[j] = i
	}

	// 双指针，从第 i 小的数开始，向左一步，最远能跳到第 left 小的数
	pa := make([][]int, n)
	mx := bits.Len(uint(n))
	left := 0
	for i, j := range idx {
		for nums[j]-nums[idx[left]] > maxDiff {
			left++
		}
		pa[i] = make([]int, mx)
		pa[i][0] = left
	}

	// 倍增
	for i := range mx - 1 {
		for x := range pa {
			p := pa[x][i]
			pa[x][i+1] = pa[p][i]
		}
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		if l == r { // 不用跳
			continue
		}
		l, r = rank[l], rank[r]
		if l > r { // 保证 l 在 r 左边
			l, r = r, l
		}
		// 从 r 开始，向左跳到 l
		res := 0
		for k := mx - 1; k >= 0; k-- {
			if pa[r][k] > l {
				res |= 1 << k
				r = pa[r][k]
			}
		}
		if pa[r][0] > l { // 无法跳到 l
			ans[qi] = -1
		} else {
			ans[qi] = res + 1 // 再跳一步就能到 l
		}
	}
	return ans
}

func Test_path_existence_queries_in_a_graph_ii(t *testing.T) {
	tests := []struct {
		n       int
		nums    []int
		maxDiff int
		queries [][]int
	}{
		{n: 5, nums: []int{1, 8, 3, 4, 2}, maxDiff: 3, queries: [][]int{{0, 3}, {2, 4}}},
		{n: 5, nums: []int{5, 3, 1, 9, 10}, maxDiff: 2, queries: [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}}},
		{n: 3, nums: []int{3, 6, 1}, maxDiff: 1, queries: [][]int{{0, 0}, {0, 1}, {1, 2}}},
	}
	for _, tt := range tests {
		t.Log(pathExistenceQueries(tt.n, tt.nums, tt.maxDiff, tt.queries))
	}
}
