package numberofpairsofinterchangeablerectangles

import (
	"testing"
)

func interchangeableRectangles(rectangles [][]int) (ans int64) {
	rate := make(map[[2]int]int64, 0)
	for _, rect := range rectangles {
		width, height := rect[0], rect[1]
		g := gcd(width, height)
		rate[[2]int{width / g, height / g}]++
	}
	for _, cnt := range rate {
		ans += cnt * (cnt - 1) / 2
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func Test_number_of_pairs_of_interchangeable_rectangles(t *testing.T) {
	rectangles := [][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}
	t.Log(interchangeableRectangles(rectangles))
	rectangles = [][]int{{4, 5}, {7, 8}}
	t.Log(interchangeableRectangles(rectangles))
}
