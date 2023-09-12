package coursescheduleivmid

import (
	"testing"
)

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) (ans []bool) {
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
	}
	graph := make([][]int, n)
	indeg := make([]int, n)
	for _, pre := range prerequisites {
		from, to := pre[0], pre[1]
		graph[from] = append(graph[from], to)
		indeg[to]++
	}
	q := make([]int, 0)
	for i, x := range indeg {
		if x == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range graph[cur] {
			f[cur][next] = true
			for x := 0; x < n; x++ {
				// 枚举x,x连通cur||next
				f[x][next] = f[x][next] || f[x][cur]
			}
			indeg[next]--
			if indeg[next] == 0 {
				q = append(q, next)
			}
		}
	}
	for _, q := range queries {
		ans = append(ans, f[q[0]][q[1]])
	}
	return
}

func Test_course_schedule_iv(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
	}{
		{numCourses: 2, prerequisites: [][]int{{1, 0}}, queries: [][]int{{0, 1}, {1, 0}}},
		{numCourses: 2, prerequisites: [][]int{}, queries: [][]int{{1, 0}, {0, 1}}},
		{numCourses: 3, prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}}, queries: [][]int{{1, 0}, {1, 2}}},
	}
	for _, tt := range tests {
		t.Log(checkIfPrerequisite(tt.numCourses, tt.prerequisites, tt.queries))
	}
}
