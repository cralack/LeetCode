package battleships_in_a_board_mid

import (
	"testing"
)

var dirs = [][]int{{0, 1}, {1, 0}}

func dfs(board [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) ||
		board[x][y] == '.' {
		return
	}
	board[x][y] = '.'
	for i := range dirs {
		dfs(board, x+dirs[i][0], y+dirs[i][1])
	}
}

func countBattleships(board [][]byte) (ans int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				ans++
				dfs(board, i, j)
			}
		}
	}
	return ans
}

func Test_battleships_in_a_board(t *testing.T) {
	tests := []struct {
		board [][]byte
	}{
		{board: [][]byte{
			{'X', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
		}},
		{board: [][]byte{{'.'}}},
	}
	for _, test := range tests {
		t.Log(countBattleships(test.board))
	}

}
