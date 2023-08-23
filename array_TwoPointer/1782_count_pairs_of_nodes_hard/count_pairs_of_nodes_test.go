package countpairsofnodes

import (
	"sort"
	"testing"
)

func countPairs(n int, edges [][]int, queries []int) []int {
	// deg[i] 表示与点 i 相连的边的数目
	deg := make([]int, n+1) // 节点编号从 1 到 n
	type edge struct{ x, y int }
	cntE := map[edge]int{}
	for _, e := range edges {
		x, y := e[0], e[1]
		if x > y {
			x, y = y, x
		}
		deg[x]++
		deg[y]++
		// 统计每条边的出现次数，注意 1-2 和 2-1 算同一条边
		cntE[edge{x, y}]++
	}

	ans := make([]int, len(queries))
	sortedDeg := append([]int(nil), deg...)
	sort.Ints(sortedDeg) // 排序，为了双指针
	for j, q := range queries {
		left, right := 1, n // 相向双指针
		for left < right {
			if sortedDeg[left]+sortedDeg[right] <= q {
				left++
			} else {
				ans[j] += right - left
				right--
			}
		}
		// 考虑边重复统计
		for edge, cnt := range cntE {
			sum := deg[edge.x] + deg[edge.y]
			if sum > q && sum-cnt <= q {
				ans[j]--
			}
		}
	}
	return ans
}

func Test_count_pairs_of_nodes(t *testing.T) {
	tests := []struct {
		n       int
		edges   [][]int
		queries []int
	}{
		{n: 4, edges: [][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1}}, queries: []int{2, 3}},
		{n: 5, edges: [][]int{{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5}}, queries: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Log(countPairs(tt.n, tt.edges, tt.queries))
	}
}
