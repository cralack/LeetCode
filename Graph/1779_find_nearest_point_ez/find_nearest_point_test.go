package findnearestpoint

import (
	"math"
	"testing"
)

func nearestValidPoint(x int, y int, points [][]int) (ans int) {
	minDis := math.MaxInt16
	ans = -1
	for idx, point := range points {
		curX, curY := point[0], point[1]
		if curX == x || curY == y {
			dis := abs(curX-x) + abs(curY-y)
			if dis < minDis {
				minDis = dis
				ans = idx
			}
		}
	}
	return ans
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
func Test_find_nearest_point(t *testing.T) {
	x, y := 3, 4
	points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	t.Log(nearestValidPoint(x, y, points))
	points = [][]int{{3, 4}}
	t.Log(nearestValidPoint(x, y, points))
	points = [][]int{{2, 3}}
	t.Log(nearestValidPoint(x, y, points))
}
