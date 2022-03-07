package pathwithmaximumprobability

import (
	"container/heap"
	"testing"
)

type Node struct {
	id   int
	prob float64
}
type PQ []Node

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].prob > h[j].prob }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func maxProbability(n int, edges [][]int, succProb []float64, start, end int) float64 {
	graph := make([][][]float64, n)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		weight := succProb[i]
		graph[from] = append(graph[from], []float64{float64(to), weight})
		graph[to] = append(graph[to], []float64{float64(from), weight})
	}
	probTo := make([]float64, n)
	for i := range probTo {
		probTo[i] = -1
	}
	probTo[start] = 1
	pq := &PQ{Node{start, 1}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Node)
		if cur.id == end {
			return cur.prob
		}
		if cur.prob < probTo[cur.id] {
			continue
		}
		for _, nextNode := range graph[cur.id] {
			nextNodeID := int(nextNode[0])
			nextProb := probTo[cur.id] * nextNode[1]
			if probTo[nextNodeID] < nextProb {
				probTo[nextNodeID] = nextProb
				heap.Push(pq, Node{nextNodeID, nextProb})
			}
		}
	}
	return 0.0
}
func maxProbability_better(n int, edges [][]int, succProb []float64, start, end int) float64 {
	prob := make([]float64, n)
	prob[start] = 1
	ok := true
	for ok {
		ok = false
		for i := 0; i < len(edges); i++ {
			cur := edges[i][0]
			next := edges[i][1]
			if prob[cur] < prob[next]*succProb[i] {
				prob[cur] = prob[next] * succProb[i]
				ok = true
			}
			if prob[next] < prob[cur]*succProb[i] {
				prob[next] = prob[cur] * succProb[i]
				ok = true
			}
		}
	}
	return prob[end]
}

func Test_path_with_maximum_probability(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2
	t.Log(maxProbability(n, edges, succProb, start, end))
	n = 3
	edges = [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb = []float64{0.5, 0.5, 0.3}
	start = 0
	end = 2
	t.Log(maxProbability_better(n, edges, succProb, start, end))
	n = 3
	edges = [][]int{{0, 1}}
	succProb = []float64{0.5}
	start = 0
	end = 2
	t.Log(maxProbability_better(n, edges, succProb, start, end))
}
