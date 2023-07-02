package makingalargeislandhard

import (
	"testing"
)

var parent []int
var size []int

func largestIsland(grid [][]int) (ans int) {
	// init var
	n := len(grid)
	parent = make([]int, n*n+1) // offset i*n+j+1
	size = make([]int, n*n+1)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// init arr
	for i := 1; i <= n*n; i++ {
		parent[i] = i
		size[i] = 1
	}

	// traverse grid
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for _, dir := range dirs {
				// nextNode:(x,y)
				x := i + dir[0]
				y := j + dir[1]
				if n <= x || x < 0 || n <= y || y < 0 ||
					grid[x][y] == 0 {
					continue
				}
				union(i*n+j+1, x*n+y+1) // curNode == 1 && nextNode == 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, size[find(i*n+j+1)])
			} else { // grid[i][j] == 0时作为反转点考察
				tot := 1
				set := make(map[int]bool, 0)
				for _, dir := range dirs {
					// nextNode:(x,y)
					x := i + dir[0]
					y := j + dir[1]
					if n <= x || x < 0 || n <= y || y < 0 ||
						grid[x][y] == 0 {
						continue
					}
					// 未遍历到则tot+=size[nextNode]
					root := find(x*n + y + 1)
					if _, has := set[root]; has {
						continue
					}
					tot += size[root]
					set[root] = true
				}
				ans = max(ans, tot)
			}
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
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func union(p, q int) {
	rootP := find(p)
	rootQ := find(q)
	if rootP == rootQ {
		return
	}
	if size[rootP] > size[rootQ] {
		union(q, p)
	} else {
		size[rootQ] += size[rootP]
		parent[rootP] = parent[rootQ]
	}
}

func Test_making_a_large_island(t *testing.T) {
	grid := [][]int{{1, 0}, {0, 1}}
	t.Log(largestIsland(grid))
	grid = [][]int{{1, 1}, {1, 0}}
	t.Log(largestIsland(grid))
	grid = [][]int{{1, 1}, {1, 1}}
	t.Log(largestIsland(grid))
	grid = [][]int{
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 1},
		{0, 1, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 1, 0, 0}}
	t.Log(largestIsland(grid))
}
