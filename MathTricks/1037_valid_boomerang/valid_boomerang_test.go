package validboomerang

import "testing"

func isBoomerang(points [][]int) bool {
	v1 := [2]int{points[1][0] - points[0][0], points[1][1] - points[0][1]}
	v2 := [2]int{points[2][0] - points[0][0], points[2][1] - points[0][1]}
	return v1[0]*v2[1]-v1[1]*v2[0] != 0
}
func Test_valid_boomerang(t *testing.T) {
	points := [][]int{{1, 1}, {2, 3}, {3, 2}}
	t.Log(isBoomerang(points))
	points = [][]int{{1, 1}, {2, 2}, {3, 3}}
	t.Log(isBoomerang(points))
}
