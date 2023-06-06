package equalrowandcolumnpairs

import "testing"

func equalPairs(grid [][]int) int {
	help := map[[200]int]int{}
	toKey := func(in []int) [200]int {
		temp := [200]int{}
		for k1, v1 := range in {
			temp[k1] = v1
		}
		return temp
	}
	for _, v := range grid {
		help[toKey(v)] += 1
	}
	n := len(grid)
	ans := 0
	for i := 0; i < n; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			temp[j] = grid[j][i]
		}
		ans += help[toKey(temp)]
	}
	return ans
}
func Test_equal_row_and_column_pairs(t *testing.T) {
	tests := []struct {
		grid [][]int
	}{
		{grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}},
		{grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}}}
	for _, tt := range tests {
		t.Log(equalPairs(tt.grid))
	}
}
