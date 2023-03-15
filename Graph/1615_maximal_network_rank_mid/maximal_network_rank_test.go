package maximalnetworkrankmid

import (
	"testing"
)

func maximalNetworkRank(n int, roads [][]int) (ans int) {
	graph := make([][]int, n)
	cnt := make([]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}

	for _, r := range roads {
		from, to := r[0], r[1]
		graph[from][to] = 1
		graph[to][from] = 1
		cnt[from]++
		cnt[to]++
	}
	for from := 0; from < n; from++ {
		for to := from + 1; to < n; to++ {
			ans = max(ans, cnt[from]+cnt[to]-graph[from][to])
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_maximal_network_rank(t *testing.T) {
	tests := []struct {
		n     int
		roads [][]int
	}{
		{n: 4, roads: [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}},
		{n: 5, roads: [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}},
		{n: 8, roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}},
	}
	for _, tt := range tests {
		t.Log(maximalNetworkRank(tt.n, tt.roads))
	}
}
