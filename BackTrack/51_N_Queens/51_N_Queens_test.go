package nqueens

import (
	"testing"
)

func solveNQueens(n int) (res [][]string) {
	if n == 1 {
		res = append(res, []string{"Q"})
		return res
	}
	if 1 < n && n <= 3 {
		return res
	}
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}

	var trackback func(int, [][]bool)
	trackback = func(row int, board [][]bool) {
		if row == n {
			//generate board result
			tmp := []string{}
			for i := 0; i < n; i++ {
				row := make([]byte, n)
				for j := 0; j < n; j++ {
					if board[i][j] {
						row[j] = 'Q'
					} else {
						row[j] = '.'
					}
				}
				tmp = append(tmp, string(row))
			}
			res = append(res, tmp)
			return
		}

		for col := 0; col < n; col++ {
			if !isValid(row, col, board) {
				continue
			}
			board[row][col] = true
			trackback(row+1, board)
			board[row][col] = false
		}

	}

	trackback(0, board)
	return res
}

func isValid(row, col int, board [][]bool) bool {
	n := len(board)

	//check upside
	for i := row; i >= 0; i-- {
		if board[i][col] == true {
			return false
		}
	}
	//check up-right side
	for i, j := row, col; 0 <= i && j < n; i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}
	//check up-left side
	for i, j := row, col; 0 <= i && 0 <= j; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}
	return true
}

func TestNQueens(t *testing.T) {
	n := 4
	board := solveNQueens(n)
	cnt := 0
	for _, res := range board {
		t.Log(res)
		cnt++
	}
	t.Log(cnt)
}
