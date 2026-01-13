package separate_squares_i_mid

import (
	"math/bits"
	"testing"
)

func separateSquares(squares [][]int) float64 {
	totArea := 0
	maxY := 0
	for _, sq := range squares {
		l := sq[2]
		totArea += l * l
		maxY = max(maxY, sq[1]+l)
	}

	check := func(y float64) bool {
		area := 0.
		for _, sq := range squares {
			yi := float64(sq[1])
			if yi < y {
				l := float64(sq[2])
				area += l * min(y-yi, l)
			}
		}
		return area >= float64(totArea)/2
	}

	left, right := 0., float64(maxY)
	for range bits.Len(uint(maxY * 1e5)) {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return (left + right) / 2
}

func Test_separate_squares_i(t *testing.T) {
	tests := []struct {
		squares [][]int
	}{
		{squares: [][]int{{0, 0, 1}, {2, 2, 1}}},
		{squares: [][]int{{0, 0, 2}, {1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Logf("%.5f", separateSquares(tt.squares))
	}
}
