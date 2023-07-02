package pathwithminimumeffort

import (
	"container/heap"
	"testing"
)

const INT_MAX = 1e+6 + 1e+2

type State struct {
	x, y          int
	costFromStart int
}
type Heap []State

func (h Heap) Less(i, j int) bool  { return h[i].costFromStart < h[j].costFromStart }
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	costTo := make([][]int, m)
	for i := range costTo {
		costTo[i] = make([]int, n)
		for j := range costTo[i] {
			costTo[i][j] = INT_MAX
		}
	}

	costTo[0][0] = 0
	pq := &Heap{State{0, 0, 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)
		curCost := cur.costFromStart
		// 到达终点提前结束
		if cur.x == m-1 && cur.y == n-1 {
			return curCost
		}
		if curCost > costTo[cur.x][cur.y] {
			continue
		}
		dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		var nx, ny int
		for _, dir := range dirs {
			nx, ny = cur.x+dir[0], cur.y+dir[1]
			if m <= nx || nx < 0 || n <= ny || ny < 0 {
				continue
			}
			// 一整条路径耗费的体力值是路径上相邻格子之间
			// 高度差 绝对值 的 最大值 决定的。
			nextCost := max(costTo[cur.x][cur.y],
				abs(heights[cur.x][cur.y]-heights[nx][ny]))
			if costTo[nx][ny] > nextCost {
				costTo[nx][ny] = nextCost
				heap.Push(pq, State{nx, ny, nextCost})
			}
		}

	}
	return -1
}
func Test_path_with_maximum_probability(t *testing.T) {
	heights := [][]int{
		{1, 2, 2},
		{3, 8, 2},
		{5, 3, 5}}
	t.Log(minimumEffortPath(heights))
	heights = [][]int{
		{1, 2, 3},
		{3, 8, 4},
		{5, 3, 5}}
	t.Log(minimumEffortPath(heights))
	heights = [][]int{
		{1, 2, 1, 1, 1},
		{1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1},
		{1, 1, 1, 2, 1}}
	t.Log(minimumEffortPath(heights))
}
