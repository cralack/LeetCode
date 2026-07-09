package path_existence_queries_in_a_graph_i_mid

import (
	"testing"
)

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	p := make([]int, n)
	for i := 1; i < n; i++ {
		p[i] = i
		if nums[i]-nums[i-1] <= maxDiff {
			p[i] = p[i-1]
		}
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		u, v := q[0], q[1]
		ans[i] = p[u] == p[v]
	}
	return ans
}

func Test_path_existence_queries_in_a_graph_i(t *testing.T) {
	tests := []struct {
		n       int
		nums    []int
		maxDiff int
		queries [][]int
	}{
		{n: 2, nums: []int{1, 3}, maxDiff: 1, queries: [][]int{{0, 0}, {0, 1}}},
		{n: 4, nums: []int{2, 5, 6, 8}, maxDiff: 2, queries: [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}},
	}
	for _, test := range tests {
		t.Log(pathExistenceQueries(test.n, test.nums, test.maxDiff, test.queries))
	}
}
