package mincosttoconnectallpoints

import (
	"container/heap"
	"sort"
	"testing"
)

/******************************** kruskal解法 ********************************/
type unionFind struct {
	Parent, Size []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{Parent: parent, Size: size}
}
func (uf *unionFind) find(x int) int {
	for uf.Parent[x] != x {
		uf.Parent[x] = uf.Parent[uf.Parent[x]]
		x = uf.Parent[x]
	}
	return x
}
func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.Size[fx] < uf.Size[fy] {
		fx, fy = fy, fx
	}
	uf.Size[fx] += uf.Size[fy]
	uf.Parent[fy] = fx
	return true
}
func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}
func minCostConnectPoints_kruskal(points [][]int) (ans int) {
	n := len(points)
	type edge struct{ v, w, dis int }
	edges := []edge{}
	for i, p := range points {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(p, points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].dis < edges[j].dis })
	uf := newUnionFind(n)
	left := n - 1
	for _, e := range edges {
		if uf.union(e.v, e.w) {
			ans += e.dis
			left--
			if left == 0 {
				break
			}
		}
	}
	return
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/******************************** Prim解法 ********************************/
type edge struct {
	from   int
	to     int
	weight int
}
type PQ []edge

func (h PQ) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h PQ) Len() int           { return len(h) }
func (h *PQ) Push(x interface{}) {
	*h = append(*h, x.(edge))
}
func (h *PQ) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type Prim struct {
	pq        PQ
	inMST     []bool
	weightSum int
	graph     [][]edge
}

func (p *Prim) newPrim(points [][]int) {
	n := len(points)
	p.graph = func() [][]edge {
		wei := 0
		graph := make([][]edge, n)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				ix, iy := points[i][0], points[i][1]
				jx, jy := points[j][0], points[j][1]
				wei = abs(jx-ix) + abs(jy-iy)
				graph[i] = append(graph[i], edge{i, j, wei})
				graph[j] = append(graph[j], edge{j, i, wei})
			}
		}
		return graph
	}()

	p.inMST = make([]bool, n)
	p.inMST[0] = true
	p.cut(0)
	for len(p.pq) > 0 {
		cur := heap.Pop(&p.pq).(edge)
		to, weight := cur.to, cur.weight
		if p.inMST[to] {
			continue
		}
		p.weightSum += weight
		p.inMST[to] = true
		p.cut(to)
	}
}
func (p *Prim) cut(s int) {
	for _, edge := range p.graph[s] {
		if p.inMST[edge.to] {
			continue
		}
		heap.Push(&p.pq, edge)
	}
}

func minCostConnectPoints_prim(points [][]int) int {
	p := &Prim{}
	p.newPrim(points)
	for _, v := range p.inMST {
		if !v {
			return -1
		}
	}
	return p.weightSum
}

func Test_min_cost_to_connect_all_points(t *testing.T) {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	t.Log(minCostConnectPoints_kruskal(points))
	t.Log(minCostConnectPoints_prim(points))
}
func Benchmark_min_cost_to_connect_all_points(b *testing.B) {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	b.Run("Kruskal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {

			minCostConnectPoints_kruskal(points)
		}
		b.StopTimer()
	})
	b.Run("Prim", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			minCostConnectPoints_prim(points)
		}
		b.StopTimer()
	})
}
