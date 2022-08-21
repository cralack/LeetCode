package transformtochessboardhard

import (
	"math/bits"
	"testing"
)

func movesToChessboard(board [][]int) (ans int) {
	n, INF := len(board), 0x3f3f3f3f
	getCnt := func(a, b int) int {
		cnt1, cnt2 := 0, 0
		for i := 0; i < n; i++ {
			cnt1 += ((a >> i) & 1)
			cnt2 += ((b >> i) & 1)
		}
		if cnt1 != cnt2 {
			return INF
		}
		return bits.OnesCount(uint(a^b)) / 2
	}

	bitRow := make(map[int]int, 0)
	bitCol := make(map[int]int, 0)
	//假设起始的行分别为row[0],row[1]
	row := make([]int, 0, 2)
	//起始的列分别为col[0],col[1]
	col := make([]int, 0, 2)
	for i := 0; i < n; i++ {
		a, b := 0, 0
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				a |= (1 << j)
			}
			if board[j][i] == 1 {
				b |= (1 << j)
			}
		}
		if _, has := bitRow[a]; !has {
			bitRow[a] = len(bitRow)
			row = append(row, a)
		}
		if _, has := bitCol[b]; !has {
			bitCol[b] = len(bitCol)
			col = append(col, b)
		}
	}
	//交换行和交换列均不会影响行的种类数量和列的种类数量
	//则若起始棋盘的行/列种类数不为 22，必然无法构造出合法棋盘
	if len(row) != 2 || len(col) != 2 {
		return -1
	}
	mask := (1 << n) - 1 // n个1
	//两者异或结果为必然为mask
	if row[0]^row[1] != mask || col[0]^col[1] != mask {
		return -1
	}
	t := 0
	for i := 0; i < n; i += 2 {
		t += 1 << i
	}
	//
	ans = min(getCnt(row[0], t), getCnt(row[1], t)) +
		min(getCnt(col[0], t), getCnt(col[1], t))

	if ans >= INF {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_transform_to_chessboard(t *testing.T) {
	board := [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}}
	t.Log(movesToChessboard(board))
	board = [][]int{{0, 1}, {1, 0}}
	t.Log(movesToChessboard(board))
	board = [][]int{{1, 0}, {1, 0}}
	t.Log(movesToChessboard(board))
}
