package spiralmatrixii

import "testing"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	upB, loB := 0, n-1
	leB, riB := 0, n-1
	cnt := 1
	for cnt <= n*n {
		//→
		if upB <= loB {
			for j := leB; j <= riB; j++ {
				matrix[upB][j] = cnt
				cnt++
			}
			upB++
		}
		//↓
		if leB <= riB {
			for i := upB; i <= loB; i++ {
				matrix[i][riB] = cnt
				cnt++
			}
			riB--
		}
		//←
		if upB <= loB {
			for j := riB; leB <= j; j-- {
				matrix[loB][j] = cnt
				cnt++
			}
			loB--
		}
		//↑
		if leB <= riB {
			for i := loB; upB <= i; i-- {
				matrix[i][leB] = cnt
				cnt++
			}
			leB++
		}
	}
	return matrix
}
func Test_spiral_matrix_ii(t *testing.T) {
	n := 3
	matrix := generateMatrix(n)
	t.Log(matrix)
	n = 1
	matrix = generateMatrix(n)
	t.Log(matrix)
}
