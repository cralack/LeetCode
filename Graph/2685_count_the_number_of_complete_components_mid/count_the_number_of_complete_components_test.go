package count_the_number_of_complete_components_mid

import (
	"testing"
)

func countCompleteComponents(n int, edges [][]int) (ans int) {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	vis := make([]bool, n)
	var vertex, edge int

	var dfs func(int)
	dfs = func(cur int) {
		vis[cur] = true
		vertex++
		edge += len(graph[cur])
		for _, next := range graph[cur] {
			if !vis[next] {
				dfs(next)
			}
		}
	}

	for i, b := range vis {
		if !b {
			vertex, edge = 0, 0
			dfs(i)
			if edge == vertex*(vertex-1) {
				ans++
			}
		}
	}
	return
}

func Test_count_the_number_of_complete_components(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
	}{
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}},
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}},
	}
	for _, test := range tests {
		t.Log(countCompleteComponents(test.n, test.edges))
	}
}
