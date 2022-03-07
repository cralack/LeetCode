package connectingcitieswithminimumcost

import (
	"container/heap"
	"sort"
	"testing"
)

/******************************** kruskal解法 ********************************/
func minimumCost_kruskal(n int, connections [][]int) int {
	mst := 0
	Less := func(i, j int) bool {
		return connections[i][2] <= connections[j][2]
	}
	sort.Slice(connections, Less)
	parent, size := make([]int, n+1), make([]int, n+1)
	cnt := n + 1
	for i := 1; i < n+1; i++ {
		parent[i] = i
		size[i] = 1
	}
	find := func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}
	union := func(p, q int) {
		rootP := find(p)
		rootQ := find(q)
		if rootP == rootQ {
			return
		}
		if size[rootP] > size[rootQ] {
			parent[rootQ] = rootP
			size[rootP] += size[rootQ]
		} else {
			parent[rootP] = rootQ
			size[rootQ] += size[rootP]
		}
		cnt--
	}
	connected := func(p, q int) bool {
		rootP := find(p)
		rootQ := find(q)
		return rootP == rootQ
	}
	for _, h := range connections {
		u, v := h[0], h[1]
		weight := h[2]
		if connected(u, v) {
			continue
		}
		mst += weight
		union(u, v)
	}
	if cnt == 2 {
		return mst
	}
	return -1
}

/******************************** Prim解法 ********************************/
type PriorityQueue []edge
type edge struct {
	from   int
	to     int
	weight int
}
type Prim struct {
	// 核心数据结构，存储「横切边」的优先级队列
	pq PriorityQueue
	// 类似 visited 数组的作用，记录哪些节点已经成为最小生成树的一部分
	inMST []bool
	// 记录最小生成树的权重和
	weightSum int
	// graph 是用邻接表表示的一幅图，
	// graph[s] 记录节点 s 所有相邻的边，
	// 三元组 int[]{from, to, weight} 表示一条边
	graph [][]edge
	// 按照边的权重从小到大排序
}

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(edge))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (p *Prim) newPrim(graph [][]edge) {
	p.graph = graph
	n := len(graph)
	p.inMST = make([]bool, n)
	// 假设从节点 0 开始切分
	p.inMST[0] = true
	p.cut(0)
	// 不断进行切分，向最小生成树中添加边
	for len(p.pq) > 0 {
		ed := heap.Pop(&p.pq)
		to, weight := ed.(edge).to, ed.(edge).weight
		if p.inMST[to] {
			// 节点 to 已经在最小生成树中，跳过
			// 否则这条边会产生环
			continue
		}
		// 将边 edge 加入最小生成树
		p.weightSum += weight
		p.inMST[to] = true
		// 节点 to 加入后，进行新一轮切分，会产生更多横切边
		p.cut(to)
	}
}

// 将 s 的横切边加入优先队列
func (p *Prim) cut(s int) {
	// 遍历 s 的邻边
	for _, ed := range p.graph[s] {
		to := ed.to
		if p.inMST[to] {
			// 相邻接点 to 已经在最小生成树中，跳过
			// 否则这条边会产生环
			continue
		}
		// 加入横切边队列
		heap.Push(&p.pq, ed)
	}
}
func (p *Prim) weightS() int {
	return p.weightSum
}
func (p *Prim) allConnected() bool {
	for i := 0; i < len(p.inMST); i++ {
		if !p.inMST[i] {
			return false
		}
	}
	return true
}
func minimumCost_prim(n int, connections [][]int) int {
	graph := func() [][]edge {
		graph := make([][]edge, n)
		for _, conn := range connections {
			u, v := conn[0]-1, conn[1]-1
			weight := conn[2]
			graph[u] = append(graph[u], edge{from: u, to: v, weight: weight})
			graph[v] = append(graph[v], edge{from: v, to: u, weight: weight})
		}
		return graph
	}()
	p := &Prim{}
	p.newPrim(graph)
	if !p.allConnected() {
		return -1
	}
	return p.weightS()
}

func Test_connecting_cities_with_minimum_cost(t *testing.T) {
	n, conections := 3, [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}
	t.Log(minimumCost_prim(n, conections))
	n, conections = 4, [][]int{{1, 2, 3}, {3, 4, 4}}
	t.Log(minimumCost_kruskal(n, conections))
	t.Log(minimumCost_prim(n, conections))
}
