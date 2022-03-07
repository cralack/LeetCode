package coursescheduleii

import "testing"

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := func() [][]int {
		graph := make([][]int, numCourses)
		for i := 0; i < numCourses; i++ {
			graph[i] = []int{}
		}
		for _, edge := range prerequisites {
			//修完课程 from 才能修课程 to
			from, to := edge[1], edge[0]
			graph[from] = append(graph[from], to)
		}
		return graph
	}()
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false
	postOrder := []int{}
	var traverse func(s int)
	traverse = func(s int) {
		if onPath[s] {
			hasCycle = true
		}
		if visited[s] || hasCycle {
			return
		}
		visited[s] = true
		onPath[s] = true
		for _, ver := range graph[s] {
			traverse(ver)
		}
		onPath[s] = false
		postOrder = append(postOrder, s)
	}

	for i := 0; i < numCourses; i++ {
		traverse(i)
	}
	if hasCycle {
		return []int{}
	}
	res := []int{}
	//将后序遍历结果反转即可得到拓扑排序
	for i := len(postOrder) - 1; i >= 0; i-- {
		res = append(res, postOrder[i])
	}
	return res
}

func findOrder_BFS(numCourses int, prerequisites [][]int) []int {
	graph := func() [][]int {
		graph := make([][]int, numCourses)
		for i := 0; i < numCourses; i++ {
			graph[i] = []int{}
		}
		for _, edge := range prerequisites {
			//修完课程 from 才能修课程 to
			from, to := edge[1], edge[0]
			graph[from] = append(graph[from], to)
		}
		return graph
	}()
	indgree := make([]int, numCourses)
	for _, edge := range prerequisites {
		indgree[edge[0]]++
	}

	que := []int{}
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			que = append(que, i)
		}
	}

	cnt, res := 0, make([]int, numCourses)

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		res[cnt] = cur
		cnt++
		for _, next := range graph[cur] {
			indgree[next]--
			if indgree[next] == 0 {
				que = append(que, next)
			}
		}
	}
	if cnt != numCourses {
		return []int{}
	}
	return res
}
func Test_course_schedule_ii(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	t.Log(findOrder(numCourses, prerequisites))
	t.Log(findOrder_BFS(numCourses, prerequisites))
}

func Benchmark_course_schedule_ii(b *testing.B) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	b.Run("DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findOrder(numCourses, prerequisites)
		}
		b.StopTimer()
	})
	b.Run("DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findOrder_BFS(numCourses, prerequisites)
		}
		b.StopTimer()
	})
}
