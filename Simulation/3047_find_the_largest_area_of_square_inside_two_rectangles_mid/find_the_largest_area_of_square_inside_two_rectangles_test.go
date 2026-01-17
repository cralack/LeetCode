package find_the_largest_area_of_square_inside_two_rectangles_mid

import (
	"testing"
)

func largestSquareArea(bottomLeft, topRight [][]int) int64 {
	maxSide := 0
	for i, b1 := range bottomLeft {
		t1 := topRight[i]
		for j, b2 := range bottomLeft[:i] {
			t2 := topRight[j]
			width := min(t1[0], t2[0]) - max(b1[0], b2[0])
			height := min(t1[1], t2[1]) - max(b1[1], b2[1])
			side := min(width, height)
			maxSide = max(maxSide, side)
		}
	}
	return int64(maxSide) * int64(maxSide)
}

func Test_find_the_largest_area_of_square_inside_two_rectangles(t *testing.T) {
	tests := []struct {
		bottomLeft [][]int
		topRight   [][]int
	}{
		{bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}}, topRight: [][]int{{3, 3}, {4, 4}, {6, 6}}},
		{bottomLeft: [][]int{{1, 1}, {2, 2}, {1, 2}}, topRight: [][]int{{3, 3}, {4, 4}, {3, 4}}},
		{bottomLeft: [][]int{{1, 1}, {3, 3}, {3, 1}}, topRight: [][]int{{2, 2}, {4, 4}, {4, 2}}},
	}
	for _, tt := range tests {
		t.Log(largestSquareArea(tt.bottomLeft, tt.topRight))
	}
}
