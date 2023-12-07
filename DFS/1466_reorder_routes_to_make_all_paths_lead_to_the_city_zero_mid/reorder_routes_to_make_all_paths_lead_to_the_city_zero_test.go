package reorder_routes_to_make_all_paths_lead_to_the_city_zero_mid

import (
	"testing"
)

func minReorder(n int, connections [][]int) int {
	graph := make([][][2]int, n)
	for _, c := range connections {
		from, to := c[0], c[1]
		graph[from] = append(graph[from], [2]int{to, 1})
		graph[to] = append(graph[to], [2]int{from, 0})
	}
	var dfs func(int, int) int
	dfs = func(cur int, from int) (ans int) {
		for _, edge := range graph[cur] {
			if next, flag := edge[0], edge[1]; next != from {
				ans += flag + dfs(next, cur)
			}
		}
		return
	}
	return dfs(0, -1)
}

func Test_reorder_routes_to_make_all_paths_lead_to_the_city_zero(t *testing.T) {
	tests := []struct {
		n           int
		connections [][]int
	}{
		{n: 6, connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}},
		{n: 5, connections: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}},
		{n: 3, connections: [][]int{{1, 0}, {2, 0}}},
	}
	for _, tt := range tests {
		t.Log(minReorder(tt.n, tt.connections))
	}
}
