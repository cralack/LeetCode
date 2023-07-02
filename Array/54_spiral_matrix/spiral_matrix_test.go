package spiralmatrix

import "testing"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upperBound, lowerBound := 0, m-1
	leftBound, rightBound := 0, n-1
	res := []int{}
	for len(res) < m*n {
		// →
		if upperBound <= lowerBound {
			for j := leftBound; j <= rightBound; j++ {
				res = append(res, matrix[upperBound][j])
			}
			upperBound++
		}
		// ↓
		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				res = append(res, matrix[i][rightBound])
			}
			rightBound--
		}
		// ←
		if upperBound <= lowerBound {
			for j := rightBound; leftBound <= j; j-- {
				res = append(res, matrix[lowerBound][j])
			}
			lowerBound--
		}
		// ↑
		if leftBound <= rightBound {
			for i := lowerBound; upperBound <= i; i-- {
				res = append(res, matrix[i][leftBound])
			}
			leftBound++
		}
	}

	return res
}
func Test_spiral_matrix(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	t.Log(spiralOrder(matrix))
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	t.Log(spiralOrder(matrix))
}
