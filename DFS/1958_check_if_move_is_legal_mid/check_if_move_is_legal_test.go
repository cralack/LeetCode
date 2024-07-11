package check_if_move_is_legal_mid

import (
	"testing"
)

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	n, m := len(board), len(board[0])
	dirs := []struct{ x, y int }{{1, 0}, {1, -1}, {0, -1}, {-1, -1},
		{-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	for _, dir := range dirs {
		x, y := cMove+dir.x, rMove+dir.y
		if m <= x || x < 0 || n <= y || y < 0 || board[y][x] != color^'B'^'W' {
			continue
		}
		for {
			x += dir.x
			y += dir.y
			if m <= x || x < 0 || n <= y || y < 0 || board[y][x] == '.' {
				break
			}
			if board[y][x] == color {
				return true
			}
		}
	}
	return false
}

func Test_check_if_move_is_legal(t *testing.T) {
	tests := []struct {
		board [][]byte
		rMove int
		cMove int
		color byte
		want  bool
	}{
		{
			board: [][]byte{
				{'.', '.', '.', 'B', '.', '.', '.', '.'},
				{'.', '.', '.', 'W', '.', '.', '.', '.'},
				{'.', '.', '.', 'W', '.', '.', '.', '.'},
				{'.', '.', '.', 'W', '.', '.', '.', '.'},
				{'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'},
				{'.', '.', '.', 'B', '.', '.', '.', '.'},
				{'.', '.', '.', 'B', '.', '.', '.', '.'},
				{'.', '.', '.', 'W', '.', '.', '.', '.'},
			},
			rMove: 4,
			cMove: 3,
			color: 'B',
			want:  true,
		},
		{
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', 'B', '.', '.', 'W', '.', '.', '.'},
				{'.', '.', 'W', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'W', 'B', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', 'B', 'W', '.', '.'},
				{'.', '.', '.', '.', '.', '.', 'W', '.'},
				{'.', '.', '.', '.', '.', '.', '.', 'B'}},
			rMove: 4,
			cMove: 4,
			color: 'W',
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Log(checkMove(tt.board, tt.rMove, tt.cMove, tt.color) == tt.want)
	}
}
