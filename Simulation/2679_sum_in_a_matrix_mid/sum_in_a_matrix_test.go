package suminamatrixmid

import (
	"sort"
	"testing"
)

func matrixSum(nums [][]int) (ans int) {
	for _, row := range nums {
		sort.Ints(row)
	}
	for j := range nums[0] {
		max := 0
		for i := range nums {
			if max < nums[i][j] {
				max = nums[i][j]
			}
		}
		ans += max
	}
	return
}
func Test_sum_in_a_matrix(t *testing.T) {
	tests := []struct {
		nums  [][]int
		wanna int
	}{
		{nums: [][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}, wanna: 15},
		{nums: [][]int{{1}}, wanna: 1},
	}
	for _, tt := range tests {
		t.Log(matrixSum(tt.nums) == tt.wanna)
	}
}
