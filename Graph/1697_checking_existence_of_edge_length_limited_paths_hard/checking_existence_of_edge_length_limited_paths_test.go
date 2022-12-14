package checkingexistenceofedgelengthlimitedpaths

import (
	"sort"
	"testing"
)

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) (ans []bool) {
	//所有边按照边权从小到大排序
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	//并查集框架
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(p, q int) {
		parent[find(p)] = find(q)
	}
	isConnected := func(p, q int) bool {
		return find(p) == find(q)
	}

	m := len(queries)
	ans = make([]bool, m)
	//将边权排序后对应ans的序列
	qid := make([]int, m)
	for i := range qid {
		qid[i] = i
	}
	sort.Slice(qid, func(i, j int) bool {
		return queries[qid[i]][2] < queries[qid[j]][2]
	})
	j := 0
	for _, i := range qid {
		//对于每个查询,利用并查集的查询操作判断两点是否连通
		qa, qb, limit := queries[i][0], queries[i][1], queries[i][2]
		//将边权严格小于 limit 的所有边加入并查集
		for j < len(edgeList) && edgeList[j][2] < limit {
			u, v := edgeList[j][0], edgeList[j][1]
			union(u, v)
			j++
		}

		ans[i] = isConnected(qa, qb)
	}
	return
}

func Test_checking_existence_of_edge(t *testing.T) {
	n := 3
	edgeList := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	queries := [][]int{{0, 1, 2}, {0, 2, 5}}
	t.Log(distanceLimitedPathsExist(n, edgeList, queries))
	n = 5
	edgeList = [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}
	queries = [][]int{{0, 4, 14}, {1, 4, 13}}
	t.Log(distanceLimitedPathsExist(n, edgeList, queries))
}
