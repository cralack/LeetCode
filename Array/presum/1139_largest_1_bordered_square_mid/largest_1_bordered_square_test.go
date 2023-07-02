package largest1borderedsquare

import "testing"

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowSum := make([][]int, m)
	for i := range rowSum {
		rowSum[i] = make([]int, n+1)
	}
	colSum := make([][]int, n)
	for i := range colSum {
		colSum[i] = make([]int, m+1)
	}
	for i, row := range grid {
		for j, x := range row {
			rowSum[i][j+1] = rowSum[i][j] + x // 每行的前缀和
			colSum[j][i+1] = colSum[j][i] + x // 每列的前缀和
		}
	}
	// 枚举正方形最大边长d
	for d := min(m, n); d > 0; d-- {
		// 枚举正方形左上角坐标(i,j)
		for i := 0; i <= m-d; i++ {
			for j := 0; j <= n-d; j++ {
				// 判断当前正方形边长是否成立
				if rowSum[i][j+d]-rowSum[i][j] == d && // 上边
					colSum[j][i+d]-colSum[j][i] == d && // 左边
					rowSum[i+d-1][j+d]-rowSum[i+d-1][j] == d && // 下边
					colSum[j+d-1][i+d]-colSum[j+d-1][i] == d { // 右边
					return d * d
				}
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func Test_largest_1_bordered_square(t *testing.T) {
	grid := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	t.Log(largest1BorderedSquare(grid))
	grid = [][]int{{1, 1, 0, 0}}
	t.Log(largest1BorderedSquare(grid))
}
