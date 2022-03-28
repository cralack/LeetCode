package cheapestflightswithinkstops

import (
	"container/heap"
	"math"
	"testing"
)

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if res > v {
			res = v
		}
	}
	return res
}
func findCheapestPrice_rec(n int, flights [][]int, src int, dst int, k int) int {
	k++
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -888
		}
	}
	// 将中转站个数转化成边的条数
	indegree := make(map[int][][]int, 0)
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		indegree[to] = append(indegree[to], []int{from, price})
	}
	var dp func(s, k int) int
	// 定义：从 src 出发，k 步之内到达 s 的最短路径权重
	dp = func(s, k int) int {
		// base case
		if s == src {
			return 0
		}
		if k == 0 {
			return -1
		}
		if memo[s][k] != -888 {
			return memo[s][k]
		}
		// 初始化为最大值，方便等会取最小值
		res := math.MaxInt32
		if _, ok := indegree[s]; ok {
			for _, v := range indegree[s] {
				from := v[0]
				price := v[1]
				subProblem := dp(from, k-1)
				if subProblem != -1 {
					res = min(res, subProblem+price)
				}
			}
		}
		if res == math.MaxInt32 {
			memo[s][k] = -1
		} else {
			memo[s][k] = res
		}
		return memo[s][k]
	}
	return dp(dst, k)
}
func findCheapestPrice_ite(n int, flights [][]int, src int, dst int, k int) int {
	const INF = math.MaxInt32
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][src] = 0
	for times := 1; times <= k+1; times++ {
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			dp[times][to] = min(dp[times][to], dp[times-1][from]+price)
		}
	}
	ans := INF
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == INF {
		ans = -1
	}
	return ans
}

type State struct {
	id             int // 图节点的 id
	distFromStart  int // 从 src 节点到当前节点的花费
	nodeNumFromSrc int // 从 src 节点到当前节点经过的节点个数
}
type Heap []State

func (h Heap) Less(i, j int) bool  { return h[i].distFromStart < h[j].distFromStart }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Heap) Len() int            { return len(h) }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func findCheapestPrice_Dijkstra(n int, flights [][]int, src int, dst int, k int) int {
	//init graph
	graph := make(map[int][][]int, n+1)
	for _, edge := range flights {
		from, to, wei := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], []int{to, wei})
	}
	k++
	return dijkstra(graph, src, dst, k, n)
}

// 输入一个起点 src，计算从 src 到其他节点的最短距离
func dijkstra(graph map[int][][]int, src, dst, k, n int) int {
	// 定义：从起点 src 到达节点 i 的最短路径权重为 distTo[i]
	distTO := make([]int, n)
	// 定义：从起点 src 到达节点 i 至少要经过 nodeNumTo[i] 个节点
	nodeNumTO := make([]int, n)
	for i := 0; i < n; i++ {
		distTO[i] = math.MaxInt32
		nodeNumTO[i] = math.MaxInt32
	}
	// base case
	distTO[src] = 0
	nodeNumTO[src] = 0
	// 优先级队列，costFromSrc 较小的排在前面
	pq := &Heap{{id: src, distFromStart: 0, nodeNumFromSrc: 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)
		curNodeid := cur.id
		costFromSrc := cur.distFromStart
		curNodeNumFromSrc := cur.nodeNumFromSrc
		if curNodeid == dst {
			// 找到最短路径
			return costFromSrc
		}
		if curNodeNumFromSrc == k {
			// 中转次数耗尽
			continue
		}
		// 将 curNode 的相邻节点装入队列
		for _, neighbor := range graph[curNodeid] {
			nextNodeId := neighbor[0]
			costToNextNode := costFromSrc + neighbor[1]
			// 中转次数消耗 1
			nextNodeNumFromSrc := curNodeNumFromSrc + 1
			// 更新 dp table
			if distTO[nextNodeId] > costToNextNode {
				distTO[nextNodeId] = costToNextNode
				nodeNumTO[nextNodeId] = nextNodeNumFromSrc
			} // 剪枝，如果中转次数更多，花费还更大，那必然不会是最短路径
			if costToNextNode > distTO[nextNodeId] &&
				nextNodeNumFromSrc > nodeNumTO[nextNodeId] {
				continue
			}
			heap.Push(pq, State{nextNodeId, costToNextNode, nextNodeNumFromSrc})
		}
	}
	return -1
}
func Test_cheapest_flights_within_k_stops(t *testing.T) {
	edges := [][]int{{0, 4, 5}, {0, 3, 8}, {0, 1, 9},
		{4, 5, 6}, {3, 5, 2}, {1, 2, 2}, {2, 5, 1}}
	n, src, dst, k := 6, 0, 5, 3
	t.Log(findCheapestPrice_rec(n, edges, src, dst, k))
	t.Log(findCheapestPrice_ite(n, edges, src, dst, k))
	t.Log(findCheapestPrice_Dijkstra(n, edges, src, dst, k))
}

func Benchmark_cheapest_flights_within_k_stops(b *testing.B) {
	edges := [][]int{{0, 4, 5}, {0, 3, 8}, {0, 1, 9},
		{4, 5, 6}, {3, 5, 2}, {1, 2, 2}, {2, 5, 1}}
	n, src, dst, k := 6, 0, 5, 3
	b.Run("ite", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findCheapestPrice_ite(n, edges, src, dst, k)
		}
		b.StopTimer()
	})
	b.Run("rec", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findCheapestPrice_rec(n, edges, src, dst, k)
		}
		b.StopTimer()
	})
	b.Run("dijkstra", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findCheapestPrice_Dijkstra(n, edges, src, dst, k)
		}
		b.StopTimer()
	})
}
