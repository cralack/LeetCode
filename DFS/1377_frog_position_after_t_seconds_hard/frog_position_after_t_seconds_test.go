package frogpositionaftertseconds

import "testing"

func frogPosition(n int, edges [][]int, t int, target int) (ans float64) {
	graph := make([][]int, n+1)
	graph[1] = []int{0}
	for _, edge := range edges {
		// 无向图邻接包括父节点与子节点
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	var dfs func(int, int, int, int) bool
	dfs = func(cur, father, leftT, rate int) bool {
		if cur == target &&
			// t秒后到达target 或 targe 是叶子停在原地
			(leftT == 0 || len(graph[cur]) == 1) {
			ans = 1 / float64(rate)
			return true
		}
		if cur == target || leftT == 0 {
			return false
		}
		// 子节点规模
		size := len(graph[cur]) - 1
		for _, next := range graph[cur] {
			if next != father &&
				dfs(next, cur, leftT-1, rate*size) {
				return true
			}
		}
		return false
	}

	dfs(1, 0, t, 1)
	return
}

func Test_frog_position_after_t_seconds(t *testing.T) {
	tests := []struct {
		n      int
		edges  [][]int
		t      int
		target int
	}{
		{n: 7, edges: [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}},
			t: 2, target: 4},
		{n: 7, edges: [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}},
			t: 1, target: 7},
		{n: 3, edges: [][]int{{2, 1}, {3, 2}}, t: 1, target: 2},
	}
	for _, tt := range tests {
		t.Log(frogPosition(tt.n, tt.edges, tt.t, tt.target))
	}
}
