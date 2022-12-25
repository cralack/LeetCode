package buildingboxes

import "testing"

func minimumBoxes(n int) (ans int) {
	maxN := 0
	for i := 1; maxN+ans+i <= n; i++ {
		ans += i
		maxN += ans
	}
	for j := 1; maxN < n; j++ {
		ans++
		maxN += j
	}
	return
}
func Test_building_boxes(t *testing.T) {
	n := 3
	t.Log(minimumBoxes(n))
	n = 4
	t.Log(minimumBoxes(n))
	n = 16
	t.Log(minimumBoxes(n))
}
