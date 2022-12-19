package findifpathexists

import "testing"

func validPath_dfs(n int, edges [][]int, source int, destination int) (ans bool) {
	graph := make([][]int, n)
	for _, edg := range edges {
		from, to := edg[0], edg[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	vis := make([]bool, n)
	ans = false
	var dfs func(int)
	dfs = func(cur int) {
		if vis[cur] {
			return
		}
		vis[cur] = true
		if cur == destination {
			ans = true
			return
		}
		for _, next := range graph[cur] {
			dfs(next)
		}
	}
	dfs(source)
	return
}

func validPath_uf(n int, edges [][]int, source int, destination int) bool {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	//union
	for _, edg := range edges {
		parent[find(edg[0])] = find(edg[1])
	}
	return find(source) == find(destination)
}
func Test_find_if_path_exists_in_graph(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	destination := 2
	t.Log(validPath_dfs(n, edges, source, destination))
	t.Log(validPath_uf(n, edges, source, destination))
	n = 6
	edges = [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	source = 0
	destination = 5
	t.Log(validPath_dfs(n, edges, source, destination))
	t.Log(validPath_uf(n, edges, source, destination))
	n = 10
	edges = [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4},
		{5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}
	source = 7
	destination = 2
	t.Log(validPath_dfs(n, edges, source, destination))
	t.Log(validPath_uf(n, edges, source, destination))
}
