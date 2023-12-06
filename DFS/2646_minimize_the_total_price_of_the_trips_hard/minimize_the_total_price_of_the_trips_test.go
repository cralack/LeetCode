package minimize_the_total_price_of_the_trips_hard

import (
	"testing"
)

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	graph := make([][]int, n)
	for _, e := range edges {
		from, to := e[0], e[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	cnt := make([]int, n)
	var dfs func(int, int, int) bool
	dfs = func(cur, from, tar int) bool {
		cnt[cur]++
		if cur == tar {
			return true
		}
		ok := false
		for _, next := range graph[cur] {
			if next != from {
				ok = dfs(next, cur, tar)
				if ok {
					break
				}
			}
		}
		if !ok {
			cnt[cur]--
		}
		return ok
	}
	// 每个trip经过的节点cnt++
	for _, t := range trips {
		start, end := t[0], t[1]
		dfs(start, -1, end)
	}
	var dfs2 func(int, int) (int, int)
	dfs2 = func(cur, from int) (int, int) {
		a := price[cur] * cnt[cur]
		b := a >> 1
		for _, next := range graph[cur] {
			if next != from {
				// 是否将非相邻节点价格减半
				x, y := dfs2(next, cur)
				a += min(x, y)
				b += x
			}
		}
		return a, b
	}
	a, b := dfs2(0, -1)
	return min(a, b)
}

func Test_minimize_the_total_price_of_the_trips(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
	}{
		{
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}, {1, 3}},
			price: []int{2, 2, 10, 6},
			trips: [][]int{{0, 3}, {2, 1}, {2, 3}},
		},
		{
			n:     2,
			edges: [][]int{{0, 1}},
			price: []int{2, 2},
			trips: [][]int{{0, 0}},
		},
	}
	for _, tt := range tests {
		t.Log(minimumTotalPrice(tt.n, tt.edges, tt.price, tt.trips))
	}
}
