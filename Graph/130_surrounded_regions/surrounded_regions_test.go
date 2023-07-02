package surroundedregions

import (
	"fmt"
	"testing"
)

func Solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	row, col := len(board), len(board[0])
	dummy := row * col
	parent := make([]int, row*col+1)
	cnt, size := row*col, make([]int, row*col+1)
	for i := range parent {
		// 父节点指针初始指向自己
		parent[i] = i
		size[i] = 1
	}
	find := func(a int) int {
		for parent[a] != a {
			// 每次向树根遍历的同时，顺手将树高缩短了
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
	connected := func(p, q int) bool {
		rootP, rootQ := find(p), find(q)
		return rootP == rootQ
	}
	// 将首列和末列的 O 与 dummy 连通
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			union(i*col, dummy)
		}
		if board[i][col-1] == 'O' {
			union(i*col+col-1, dummy)
		}
	}
	// 将首行和末行的 O 与 dummy 连通
	for j := 0; j < col; j++ {
		if board[0][j] == 'O' {
			union(j, dummy)
		}
		if board[row-1][j] == 'O' {
			union(col*(row-1)+j, dummy)
		}
	}
	// 方向数组 d 是上下左右搜索的常用手法
	d := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	for i := 1; i < row-1; i++ { // 刨掉首列末列
		for j := 1; j < col-1; j++ { // 首行末行
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x, y := i+d[k][0], j+d[k][1]
					if board[x][y] == 'O' {
						union(x*col+y, i*col+j)
					}
				}
			}
		}
	}
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if !connected(dummy, i*col+j) {
				board[i][j] = 'X'
			}
		}
	}
}

func Show(board [][]byte) {
	for _, row := range board {
		for _, x := range row {
			fmt.Printf("%c ", x)
		}
		fmt.Printf("\r\n")
	}
}
func solve_DFS(board [][]byte) {
	row, col := len(board), len(board[0])
	if row == 0 || col == 0 {
		return
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if row <= x || x < 0 || col <= y || y < 0 ||
			board[x][y] != 'O' {
			return
		}
		board[x][y] = '#'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < col; i++ {
		dfs(0, i)
		dfs(row-1, i)

	}
	for j := 0; j < row; j++ {
		dfs(j, 0)
		dfs(j, col-1)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
func Test_surrounded_regions(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X', 'O'},
		{'X', 'X', 'X', 'O', 'X'},
		{'O', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X', 'X'}}
	Show(board)
	t.Logf("\r\n")
	solve_DFS(board)
	Show(board)

}
