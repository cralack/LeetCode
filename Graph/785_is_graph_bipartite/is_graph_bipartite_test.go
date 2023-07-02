package isgraphbipartite

import "testing"

func isBipartite(graph [][]int) bool {
	size := len(graph)
	// 记录图中节点是否被访问过，以及图是否符合二分图性质
	visited, ok := make([]bool, size), true
	// 记录图中节点颜色，false 和 true 代表两种不同颜色
	color := make([]bool, size)

	var traverse func(vertex int)
	// DFS 遍历框架
	traverse = func(vertex int) {
		// 如果已经确定不是二分图了，就不用浪费时间再递归遍历了
		if !ok {
			return
		}
		visited[vertex] = true
		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				// 如果相邻节点neighbor没有被访问过
				// 那么应该给节点neighbor涂上和节点vertex不同的颜色
				if !ok {
					return
				}
				color[neighbor] = !color[vertex]
				// 继续遍历neighbor
				traverse(neighbor)
			} else {
				// 相邻节点neighbor已经被访问过
				// 根据vertex和neighbor的颜色判断是否是二分图
				if color[neighbor] == color[vertex] {
					// 若相同，则此图不是二分图
					ok = false
				}
			}
		}
	}
	// 因为图不一定是联通的，可能存在多个子图
	// 所以要把每个节点都作为起点进行一次遍历
	// 如果发现任何一个子图不是二分图，整幅图都不算二分图
	for v := 0; v < size; v++ {
		if !visited[v] {
			traverse(v)
		}
	}
	return ok
}
func isBipartite_BFS(graph [][]int) bool {
	size := len(graph)
	// 记录图中节点是否被访问过，以及图是否符合二分图性质
	visited, ok := make([]bool, size), true
	// 记录图中节点颜色，false 和 true 代表两种不同颜色
	color := make([]bool, size)

	bfs := func(start int) {
		queue := []int{start}
		visited[start] = true
		for len(queue) > 0 && ok {
			vertex := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[vertex] {
				if !visited[neighbor] {
					color[neighbor] = !color[vertex]
					visited[neighbor] = true
					queue = append(queue, neighbor)
				} else {
					if color[neighbor] == color[vertex] {
						ok = false
					}
				}
			}
		}
	}
	for v := 0; v < size; v++ {
		if !visited[v] {
			bfs(v)
		}
	}
	return ok
}
func Test_is_graph_bipartite(t *testing.T) {
	graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	t.Log(isBipartite_BFS(graph))
	graph2 := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	t.Log(isBipartite(graph2))
}

func Benchmark__is_graph_bipartite(b *testing.B) {
	graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	graph2 := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	b.Run("DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			isBipartite(graph)
			isBipartite(graph2)
		}
		b.StopTimer()
	})
	b.Run("BFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			isBipartite_BFS(graph)
			isBipartite_BFS(graph2)
		}
		b.StopTimer()
	})
}
