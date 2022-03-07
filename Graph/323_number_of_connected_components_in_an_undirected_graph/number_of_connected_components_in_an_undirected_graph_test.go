package numberofconnectedcomponentsinanundirectedgraph

import (
	"testing"
)

func countComponents(n int, edges [][]int) int {
	//节点 x 的节点是 parent[x]
	parent := make([]int, n)
	cnt, size := n, make([]int, n)
	for i := range parent {
		//父节点指针初始指向自己
		parent[i] = i
		size[i] = 1
	}
	find := func(a int) int {
		for parent[a] != a {
			//每次向树根遍历的同时，顺手将树高缩短了
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}
	union := func(a, b int) {
		pA := find(a)
		pB := find(b)
		if pA == pB {
			return
		}
		if size[pA] > size[pB] {
			parent[pB] = pA
			size[pA] += size[pB]
		} else {
			parent[pA] = pB
			size[pB] += size[pA]
		}
		cnt--
	}
	for i := 0; i < len(edges); i++ {
		union(edges[i][0], edges[i][1])
	}
	return cnt
}
func Test_number_of_connected_components_in_an_undirected_graph(t *testing.T) {
	n, edges := 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}} //ans=1
	t.Log(countComponents(n, edges))
	n, edges = 5, [][]int{{0, 1}, {1, 2}, {3, 4}} //ans=2
	t.Log(countComponents(n, edges))
}
