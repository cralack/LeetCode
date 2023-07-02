package cellswithoddvaluesinamatrixez

import "testing"

func oddCells(m int, n int, indices [][]int) int {
	col, row := make([]int, m), make([]int, n)
	for _, ind := range indices {
		col[ind[0]]++
		row[ind[1]]++
	}
	oddCol, oddRow := 0, 0
	for _, n := range col { // Odd↓
		oddCol += n & 1
	}
	for _, n := range row { // Odd→
		oddRow += n & 1
	} // Even→ * Odd↓ + Even↓ * Odd→
	return (n-oddRow)*oddCol + (m-oddCol)*oddRow
}
func Test_cells_with_odd_values_in_a_matrix(t *testing.T) {
	m, n := 2, 3
	indices := [][]int{{0, 1}, {1, 1}}
	t.Log(oddCells(m, n, indices))
	m, n = 2, 2
	indices = [][]int{{0, 0}, {1, 1}}
	t.Log(oddCells(m, n, indices))
}
