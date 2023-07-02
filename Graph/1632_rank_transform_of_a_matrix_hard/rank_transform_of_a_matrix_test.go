package ranktransformofamatrix

import (
	"sort"
	"testing"
)

// unionfind define
type unionFind struct {
	parent, size []int
}

func newUF(n int) *unionFind {
	p := make([]int, n)
	size := make([]int, n)
	for i := range p {
		p[i] = i
		size[i] = 1
	}
	return &unionFind{parent: p, size: size}
}
func (this *unionFind) find(x int) int {
	if this.parent[x] != x {
		this.parent[x] = this.find(this.parent[x])
	}
	return this.parent[x]
}
func (this *unionFind) union(a, b int) {
	pa, pb := this.find(a), this.find(b)
	if pa != pb {
		if this.size[pa] > this.size[pb] {
			this.parent[pb] = pa
			this.size[pa] += this.size[pb]
		} else {
			this.parent[pa] = pb
			this.size[pb] += this.size[pa]
		}
	}
}

// solution
func matrixRankTransform(matrix [][]int) (ans [][]int) {
	// 贪心解法：从小到大遍历元素，并维护每行、每列的最大秩，
	// 该元素的秩即为同行、同列的最大秩加1
	m, n := len(matrix), len(matrix[0])
	type point struct{ i, j int }
	points := make(map[int][]point)
	for i, row := range matrix {
		for j, v := range row {
			// 统计同元素的坐标数
			points[v] = append(points[v], point{i, j})
		}
	}
	// 每行、每列的最大秩
	rowMax := make([]int, m)
	colMax := make([]int, n)
	ans = make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	// 所有元素的集合
	vals := make([]int, 0)
	for p := range points {
		vals = append(vals, p)
	}
	sort.Ints(vals)
	// 从小到大遍历元素
	for _, val := range vals {
		pts := points[val]
		uf := newUF(m + n)
		rank := make([]int, m+n)
		// i,j形成映射关系?
		for _, p := range pts {
			uf.union(p.i, p.j+m)
		}
		for _, p := range pts {
			i, j := p.i, p.j
			rank[uf.find(i)] = max(rank[uf.find(i)], max(rowMax[i], colMax[j]))
		}
		for _, p := range pts {
			i, j := p.i, p.j
			ans[i][j] = 1 + rank[uf.find(i)]
			rowMax[i], colMax[j] = ans[i][j], ans[i][j]
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_rank_transform_of_a_matrix(t *testing.T) {
	matrix := [][]int{
		{1, 2}, // 1,2
		{3, 4}} // 2,3
	t.Log(matrixRankTransform(matrix))
	matrix = [][]int{
		{7, 7}, // 1,1
		{7, 7}} // 1,1
	t.Log(matrixRankTransform(matrix))
	matrix = [][]int{
		{20, -21, 14}, // 4,2,3
		{-19, 4, 19},  // 1,3,4
		{22, -47, 24}, // 5,1,6
		{-19, 4, 19}}  // 1,3,4
	t.Log(matrixRankTransform(matrix))
	matrix = [][]int{
		{7, 3, 6}, // 5,1,4
		{1, 4, 5}, // 1,2,3
		{9, 8, 2}} // 6,3,1
	t.Log(matrixRankTransform(matrix))
	matrix = [][]int{
		{-37, -50, -3, 44},  // 2,1,4,6
		{-37, 46, 13, -32},  // 2,6,5,4
		{47, -42, -3, -40},  // 5,2,4,3
		{-17, -22, -39, 24}} // 4,3,1,5
	t.Log(matrixRankTransform(matrix))
}
