package shortestpath

import "testing"

const (
	red  = iota //0
	blue        //1
)

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) (ans []int) {
	//build graph
	graph := make([][][]int, 2)
	for i := range graph {
		graph[i] = make([][]int, n)
	}
	for _, rE := range redEdges {
		from, to := rE[0], rE[1]
		graph[red][from] = append(graph[red][from], to)
	}
	for _, bE := range blueEdges {
		from, to := bE[0], bE[1]
		graph[blue][from] = append(graph[blue][from], to)
	}
	//init val
	vis := make([][2]bool, n)
	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	//define node & que
	type node struct{ idx, color int }
	que := []node{}
	que = append(que, node{0, 0}, node{0, 1})
	step := 0
	//BFS
	for len(que) > 0 {
		for k := len(que); k > 0; k-- {
			cur := que[0]
			que = que[1:]
			idx, color := cur.idx, cur.color
			if ans[idx] == -1 {
				ans[idx] = step
			}
			vis[idx][color] = true
			color ^= 1 //颜色取反
			for _, next := range graph[color][idx] {
				if !vis[next][color] {
					que = append(que, node{next, color})
				}
			}
		}
		step++
	}
	return ans
}
func Test_shortest_path_with_alternating_colors(t *testing.T) {
	n := 3
	red_edges := [][]int{{0, 1}, {1, 2}}
	blue_edges := [][]int{}
	t.Log(shortestAlternatingPaths(n, red_edges, blue_edges))
	n = 3
	red_edges = [][]int{{0, 1}}
	blue_edges = [][]int{{2, 1}}
	t.Log(shortestAlternatingPaths(n, red_edges, blue_edges))
	n = 3
	red_edges = [][]int{{1, 0}}
	blue_edges = [][]int{{2, 1}}
	t.Log(shortestAlternatingPaths(n, red_edges, blue_edges))
	n = 3
	red_edges = [][]int{{0, 1}}
	blue_edges = [][]int{{1, 2}}
	t.Log(shortestAlternatingPaths(n, red_edges, blue_edges))
	n = 3
	red_edges = [][]int{{0, 1}, {0, 2}}
	blue_edges = [][]int{{1, 0}}
	t.Log(shortestAlternatingPaths(n, red_edges, blue_edges))
}
