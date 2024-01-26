package minimum_edge_weight_equilibrium_queries_in_a_tree_hard

import (
	"slices"
	"testing"
)

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ to, wt int }
	graph := make([][]edge, n)
	for _, e := range edges {
		from, to, wt := e[0], e[1], e[2]-1
		graph[from] = append(graph[from], edge{to, wt})
		graph[to] = append(graph[to], edge{from, wt})
	}

	const mx = 14 // 2^14 > 10^4
	type pair struct {
		p   int
		cnt [26]int
	}
	pa := make([][mx]pair, n)
	depth := make([]int, n)
	var build func(int, int, int)
	build = func(v, p, d int) {
		pa[v][0].p = p
		depth[v] = d
		for _, e := range graph[v] {
			if w := e.to; w != p {
				pa[w][0].cnt[e.wt] = 1
				build(w, v, d+1)
			}
		}
	}
	build(0, -1, 0)

	// 倍增模板
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1].p = pp.p
				for j := 0; j < 26; j++ {
					pa[v][i+1].cnt[j] = p.cnt[j] + pp.cnt[j]
				}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}

	// 计算 LCA 模板（这里返回最小操作次数）
	// https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/
	f := func(v, w int) int {
		pathLen := depth[v] + depth[w] // 最后减去 depth[lca] * 2
		cnt := [26]int{}
		if depth[v] > depth[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (depth[w]-depth[v])>>i&1 > 0 {
				p := pa[w][i]
				for j := 0; j < 26; j++ {
					cnt[j] += p.cnt[j]
				}
				w = p.p
			}
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
					for j := 0; j < 26; j++ {
						cnt[j] += pv.cnt[j] + pw.cnt[j]
					}
					v, w = pv.p, pw.p
				}
			}
			for j := 0; j < 26; j++ {
				cnt[j] += pa[v][0].cnt[j] + pa[w][0].cnt[j]
			}
			v = pa[v][0].p
		}
		// 现在 v 是 LCA
		return pathLen - depth[v]*2 - slices.Max(cnt[:])
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = f(q[0], q[1])
	}
	return ans
}

func Test_minimum_edge(t *testing.T) {
	tests := []struct {
		n       int
		edges   [][]int
		queries [][]int
	}{
		{n: 7, edges: [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 4, 2},
			{4, 5, 2}, {5, 6, 2}}, queries: [][]int{{0, 3}, {3, 6}, {2, 6}, {0, 6}}},
		{n: 8, edges: [][]int{{1, 2, 6}, {1, 3, 4}, {2, 4, 6}, {2, 5, 3}, {3, 6, 6},
			{3, 0, 8}, {7, 0, 2}}, queries: [][]int{{4, 6}, {0, 4}, {6, 5}, {7, 4}}},
	}
	for _, tt := range tests {
		t.Log(minOperationsQueries(tt.n, tt.edges, tt.queries))
	}
}
