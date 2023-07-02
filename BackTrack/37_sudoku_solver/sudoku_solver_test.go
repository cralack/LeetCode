package sudokusolver

import (
	"fmt"
	"testing"
)

func solveSudoku(board [][]byte) {
	DFS(0, 0, board)
}
func DFS(i, j int, board [][]byte) bool {
	m, n := 9, 9
	if j == n {
		// 穷举到最后一列的话就换到下一行重新开始。
		return DFS(i+1, 0, board)
	}
	if i == m {
		// 找到一个可行解，触发 base case
		return true
	}
	if board[i][j] != '.' {
		// 如果不是空格，则不用穷举
		return DFS(i, j+1, board)
	}
	for n := byte('1'); n <= '9'; n++ {
		// 跳过非法数字
		if !isValid(i, j, n, board) {
			continue
		}
		board[i][j] = n
		if DFS(i, j+1, board) {
			return true
		}
		board[i][j] = '.'
	}
	// 穷举完 1~9，依然没有找到可行解，此路不通
	return false
}
func isValid(row, col int, n byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		// check row
		if board[row][i] == n {
			return false
		}
		// check col
		if board[i][col] == n {
			return false
		}
		// check nine square
		if board[row/3*3+i/3][col/3*3+i%3] == n {
			return false
		}

	}
	return true
}
func show(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println("\n\r")
	}
}

func Test_sudoku_solver(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	show(board)

	// empty board for sudoku cheat
	// board := [][]byte{
	// 	//1    2    3    4    5    6    7    8    9
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //1
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //2
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //3
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //4
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //5
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //6
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //7
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //8
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, //9
	// 	//1    2    3    4    5    6    7    8    9
	// }
	// solveSudoku(board)
	// show(board)
}

func Benchmark_sudoku(b *testing.B) {
	b.Run("backtrack", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			board := [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
			solveSudoku(board)
		}
		b.StopTimer()
	})
}
