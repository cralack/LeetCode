package graphvalidtree

import "testing"

func validTree(edges [][]int, n int) bool {
	parent, size := make([]int, n), make([]int, n)
	cnt := n
	for i := 0; i < n; i++ {
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
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if connected(u, v) {
			//对于添加的这条边 如果该边的两个节点本来就在同一连通分量里
			// 那么添加这条边会产生环 直接false
			return false
		}
		union(u, v)
	}
	//连通图数量>1则不成独树
	return cnt == 1
}
func Test_graph_valid_tree(t *testing.T) {
	n, edges := 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	t.Log(validTree(edges, n))
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	t.Log(validTree(edges, n))
}
