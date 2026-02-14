package minimum_cost_path_with_edge_reversals_mid

import (
	"container/heap"
	"testing"
)

// Dijkstra模板题
func minCost(n int, edges [][]int) int {
	type edge struct {
		to, weit int
	}
	graph := make([][]edge, n) // 邻接表
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		graph[x] = append(graph[x], edge{y, w})
		graph[y] = append(graph[y], edge{x, w * 2}) // 本题反转边
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = 1 << 31
	}
	dis[0] = 0 // 起点到自己的距离是 0
	// 堆中保存 (起点到节点 x 的最短路长度，节点 x)
	h := &hp{{0, 0}} // push start

	for h.Len() > 0 {
		cur := heap.Pop(h).(pair)
		disX, x := cur.dis, cur.x
		if disX > dis[x] { // x 之前出堆过
			continue
		}
		if x == n-1 {
			return disX
		}
		for _, neigh := range graph[x] {
			y := neigh.to
			disY := disX + neigh.weit
			if disY < dis[y] {
				dis[y] = disY // 更新 x 的邻居的最短路
				// 懒更新堆：只插入数据，不更新堆中数据
				// 相同节点可能有多个不同的 newDisY，除了最小的 newDisY，其余值都会触发上面的 continue
				heap.Push(h, pair{disY, y})
			}
		}
	}
	return -1
}

type pair struct {
	dis int // 至起点距离
	x   int // 节点编号
}
type hp []pair

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].dis < (*h)[j].dis }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)       { a := *h; v = a[len(a)-1]; *h = a[:len(a)-1]; return }

func Test_minimum_cost_path_with_edge_reversals(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
	}{
		{n: 4, edges: [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}}},
		{n: 4, edges: [][]int{{0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 3}}},
	}
	for _, test := range tests {
		t.Log(minCost(test.n, test.edges))
	}
}
