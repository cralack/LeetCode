package courseschedule

import (
	"testing"
)

//dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	//raw data to graph
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
	//DFS 遍历
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
	}

	for i := 0; i < numCourses; i++ {
		traverse(i)
	}
	return !hasCycle
}

func canFinish_BFS(numCourses int, prerequisites [][]int) bool {
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
		to := edge[0]
		//节点 to 的入度加一
		indgree[to]++
	}
	que := []int{}
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			// 节点 i 入度==0 即没有依赖的节点
			// 可以作为拓扑排序的起点，加入队列
			que = append(que, i)
		}
	}
	// 记录遍历的节点个数
	cnt := 0
	// 开始执行 BFS 循环
	for len(que) > 0 {
		// 弹出节点 cur，并将它指向的节点的入度减一
		cur := que[0]
		que = que[1:]
		cnt++
		for _, next := range graph[cur] {
			indgree[next]--
			if indgree[next] == 0 {
				// 如果入度变为 0，说明 next 依赖的节点都已被遍历
				que = append(que, next)
			}
		}
	} // 如果所有节点都被遍历过，说明不成环
	return cnt == numCourses
}
func Test_course_schedule(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	t.Log(canFinish(numCourses, prerequisites))
	t.Log(canFinish_BFS(numCourses, prerequisites))
}
