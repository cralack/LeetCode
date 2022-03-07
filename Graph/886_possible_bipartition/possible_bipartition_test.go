package possiblebipartition

import "testing"

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := func() [][]int {
		graph := make([][]int, n+1)
		for _, edge := range dislikes {
			// 「无向图」相当于「双向图」
			// from, to := edge[0], edge[1]
			graph[edge[0]] = append(graph[edge[0]], edge[1])
			graph[edge[1]] = append(graph[edge[1]], edge[0])
		}
		return graph
	}()

	visited, color := make([]bool, n+1), make([]bool, n+1)
	ok := true
	var dfs func(vertex int)
	dfs = func(vertex int) {
		if !ok {
			return
		}
		visited[vertex] = true
		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				color[neighbor] = !color[vertex]
				dfs(neighbor)
			} else {
				if color[neighbor] == color[vertex] {
					ok = false
					return
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	return ok
}

func possibleBipartition_BFS(n int, dislikes [][]int) bool {
	graph := func() [][]int {
		graph := make([][]int, n+1)
		for _, edge := range dislikes {
			// 「无向图」相当于「双向图」
			// from, to := edge[0], edge[1]
			graph[edge[0]] = append(graph[edge[0]], edge[1])
			graph[edge[1]] = append(graph[edge[1]], edge[0])
		}
		return graph
	}()

	visited, color := make([]bool, n+1), make([]bool, n+1)
	ok := true

	bfs := func(start int) {
		que := []int{start}
		visited[start] = true
		for len(que) > 0 && ok {
			vertex := que[0]
			que = que[1:]
			for _, neighbor := range graph[vertex] {
				if !visited[neighbor] {
					color[neighbor] = !color[vertex]
					visited[neighbor] = true
					que = append(que, neighbor)
				} else {
					if color[vertex] == color[neighbor] {
						ok = false
						return
					}
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		if !visited[i] {
			bfs(i)
		}
	}
	return ok
}
func Test_possible_bipartition(t *testing.T) {
	n, dislikes := 5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	t.Log(possibleBipartition(n, dislikes))
	n, dislikes = 4, [][]int{{1, 2}, {1, 3}, {2, 4}}
	t.Log(possibleBipartition_BFS(n, dislikes))
}

func Benchmark_possible_bipartition(b *testing.B) {
	n, dislikes := 5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	b.Run("DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			possibleBipartition(n, dislikes)
		}
		b.StopTimer()
	})
	b.Run("BFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			possibleBipartition_BFS(n, dislikes)
		}
		b.StopTimer()
	})
}
