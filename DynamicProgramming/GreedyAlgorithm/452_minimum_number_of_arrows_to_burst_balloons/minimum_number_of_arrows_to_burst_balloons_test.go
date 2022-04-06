package minimumnumberofarrowstoburstballoons

import (
	"sort"
	"testing"
)

func findMinArrowShots(points [][]int) int {
	n := len(points)
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	cnt, end := 1, points[0][1]
	for i := 1; i < n; i++ {
		if points[i][0] > end {
			cnt++
			end = points[i][1]
		}
	}
	return cnt
}
func Test_minimum_number_of_arrows_to_burst_balloons(t *testing.T) {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	t.Log(findMinArrowShots(points))

	points = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	t.Log(findMinArrowShots(points))

	points = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	t.Log(findMinArrowShots(points))
}
