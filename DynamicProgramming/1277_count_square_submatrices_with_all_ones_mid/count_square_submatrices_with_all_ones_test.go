package countsquaresubmatriceswithallonesmid

import "testing"

func countSquares(matrix [][]int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i, row := range matrix {
		for j, x := range row {
			if x > 0 {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1], f[i][j]) + 1
				ans += f[i+1][j+1]
			}
		}
	}
	return
}
func Test_count_square_submatrices_with_all_ones(t *testing.T) {
	tests := []struct {
		matrix [][]int
	}{
		{
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
		},
		{
			matrix: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
		},
	}

	for _, test := range tests {
		t.Log(countSquares(test.matrix))
	}
}
