package specialpositionsinabinarymatrixez

import "testing"

func numSpecial(mat [][]int) (ans int) {
	m, n := len(mat), len(mat[0])
	row, col := make([]int, m), make([]int, n)
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
				ans++
			}
		}
	}
	return
}
func Test_special_positions_in_a_binary_matrix(t *testing.T) {
	mat := [][]int{{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0}}
	t.Log(numSpecial(mat))
	mat = [][]int{{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1}}
	t.Log(numSpecial(mat))
	mat = [][]int{{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0}}
	t.Log(numSpecial(mat))
	mat = [][]int{{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1}}
	t.Log(numSpecial(mat))
}
