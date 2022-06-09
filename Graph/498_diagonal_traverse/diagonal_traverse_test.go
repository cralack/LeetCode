package diagonaltraverse

import "testing"

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	i, j := 0, 0
	for lim := 0; lim < m+n-1; lim++ {
		if lim%2 == 0 { //↗
			i = min(lim, m-1)
			j = lim - i
			for ; 0 <= i && j < n; i, j = i-1, j+1 {
				ans = append(ans, mat[i][j])
			}
		} else { //↙
			j = min(lim, n-1)
			i = lim - j
			for ; 0 <= j && i < m; i, j = i+1, j-1 {
				ans = append(ans, mat[i][j])
			}
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_diagonal_traverse(t *testing.T) {
	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	t.Log(findDiagonalOrder(mat))
	mat = [][]int{{1, 2}, {3, 4}}
	t.Log(findDiagonalOrder(mat))
}
