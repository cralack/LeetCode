package reachablenodesinsubdividedgraph

import (
	"container/heap"
	"math"
	"testing"
)

func reachableNodes(edges [][]int, maxMoves int, n int) (ans int) {
	graph := make([][]neighbor, n)
	// 建图
	for _, edge := range edges {
		from, to, wei := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], neighbor{to, wei + 1})
		graph[to] = append(graph[to], neighbor{from, wei + 1})
	}

	// 从 0 出发的最短路
	dist := dijkstra(graph, 0)
	//能抵达的原始节点数
	for _, d := range dist {
		if d <= maxMoves {
			ans++
		}
	}
	// 这条边上可以到达的节点数
	for _, edge := range edges {
		from, to, cnt := edge[0], edge[1], edge[2]
		a := max(maxMoves-dist[from], 0)
		b := max(maxMoves-dist[to], 0)
		ans += min(a+b, cnt)
	}
	return
}

type neighbor struct{ to, weight int }

// 返回各点距起点最短距离 dist[]
func dijkstra(graph [][]neighbor, start int) []int {
	dist := make([]int, len(graph))
	for i := range dist {
		//各点dist赋最大值初始化
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	h := &heapPair{{start, 0}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(node)
		idx := cur.idx
		if dist[idx] < cur.dist {
			continue
		}

		for _, edge := range graph[idx] {
			next := edge.to
			if dis := dist[idx] + edge.weight; dis < dist[next] {
				dist[next] = dis
				heap.Push(h, node{next, dis})
			}
		}
	}
	return dist
}

// 当前节点编号及距起点的起始距离
type node struct{ idx, dist int }
type heapPair []node

func (h heapPair) Len() int              { return len(h) }
func (h heapPair) Less(i, j int) bool    { return h[i].dist < h[j].dist }
func (h heapPair) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *heapPair) Push(v interface{})   { *h = append(*h, v.(node)) }
func (h *heapPair) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Test_reachable_nodes_in_subdivided_graph(t *testing.T) {
	edges, maxMoves, n := [][]int{
		{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}, 6, 3
	t.Log(reachableNodes(edges, maxMoves, n)) //13
	edges, maxMoves, n = [][]int{
		{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}}, 10, 4
	t.Log(reachableNodes(edges, maxMoves, n)) //23
	edges, maxMoves, n = [][]int{
		{1, 2, 4}, {1, 4, 5}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5}}, 17, 5
	t.Log(reachableNodes(edges, maxMoves, n)) //1
}
