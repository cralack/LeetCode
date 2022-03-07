package networkdelaytime

import (
	"container/heap"
	"testing"
)

const INT_MAX = int(^uint(0) >> 1)

type State struct {
	id            int
	distFromStart int
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func networkDelayTime(times [][]int, n int, k int) int {
	// 构造图
	graph := make([][][]int, n+1)
	for _, edge := range times {
		from, to := edge[0], edge[1]
		weight := edge[2]
		// from -> List<(to, weight)>
		// 邻接表存储图结构，同时存储权重信息
		graph[from] = append(graph[from], []int{to, weight})
	}
	// 启动 dijkstra 算法计算以节点 k 为起点到其他节点的最短路径
	distTO := dijkstra(k, graph)
	// 找到最长的那一条最短路径
	res := 0
	for i := 1; i < len(distTO); i++ {
		if distTO[i] == INT_MAX {
			// 有节点不可达，返回 -1
			return -1
		}
		res = max(res, distTO[i])
	}
	return res
}

// 输入一个起点 start，计算从 start 到其他节点的最短距离
func dijkstra(start int, graph [][][]int) []int {
	// 定义：distTo[i] 的值就是起点 start 到达节点 i 的最短路径权重
	distTo := make([]int, len(graph))
	for i := 0; i < len(distTo); i++ {
		// base case，start 到 start 的最短距离就是 0
		distTo[i] = INT_MAX
	}
	distTo[start] = 0
	// 优先级队列，distFromStart 较小的排在前面
	// 从起点 start 开始进行 BFS
	pq := &Heap{{id: start, distFromStart: 0}}

	for pq.Len() > 0 {
		curState := heap.Pop(pq).(State)
		curNodeId, curDistFromStart := curState.id, curState.distFromStart
		if curDistFromStart > distTo[curNodeId] {
			continue
		}
		// 将 curNode 的相邻节点装入队列
		for _, neighbor := range graph[curNodeId] {
			nextNodeID := neighbor[0]
			distToNextNode := distTo[curNodeId] + neighbor[1]
			// 更新 dp table
			if distTo[nextNodeID] > distToNextNode {
				distTo[nextNodeID] = distToNextNode
				heap.Push(pq, State{nextNodeID, distToNextNode})
			}
		}
	}
	return distTo
}
func Test_network_delay_time(t *testing.T) {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n, k := 4, 2
	t.Log(networkDelayTime(times, n, k))
}
