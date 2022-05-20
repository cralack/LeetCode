package cutofftreesforgolfevent

import (
	"sort"
	"testing"
)

func cutOffTree(forest [][]int) int {
	type point struct {
		x, y   int
		height int
	}
	trees := []point{}
	for i := range forest {
		for j, h := range forest[i] {
			if h > 1 { //traverse forest
				trees = append(trees, point{i, j, h})
			}
		}
	}
	sort.Slice(trees, //sort by tree's height
		func(i, j int) bool { return trees[i].height < trees[j].height })

	bfs := func(start, tar point) int {
		dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //4 direction
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[start.x][start.y] = true
		step := 0
		que := []point{start}
		for len(que) > 0 {
			size := len(que)
			for i := 0; i < size; i++ {
				//pop cur node
				cur := que[0]
				que = que[1:]
				//base case
				if cur.x == tar.x && cur.y == tar.y {
					return step
				}
				for _, d := range dir { //expand node
					next := point{cur.x + d[0], cur.y + d[1], 0}
					if 0 <= next.x && next.x < m && 0 <= next.y && next.y < n &&
						!vis[next.x][next.y] && forest[next.x][next.y] > 0 {
						que = append(que, next)
						vis[next.x][next.y] = true
					}
				}
			}
			step++
		}
		return -1
	}
	sumStep := 0
	pre := point{0, 0, 0}
	for _, t := range trees {
		stp := bfs(pre, t)
		if stp < 0 {
			return -1
		}
		sumStep += stp
		pre = t
	}
	return sumStep
}
func Test_cut_off_trees_for_golf_event(t *testing.T) {
	forest := [][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}
	t.Log(cutOffTree(forest))
}
