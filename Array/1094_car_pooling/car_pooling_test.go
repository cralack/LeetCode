package carpooling

import "testing"

func carPooling(trips [][]int, capacity int) bool {
	n := 1001
	end := 0
	cap := make([]int, n)
	for i := range cap {
		cap[i] = capacity
	}
	dif := make([]int, n)
	for _, trip := range trips {
		i, j := trip[1], trip[2]
		if end < j {
			end = j
		}
		val := trip[0]
		dif[i] -= val
		if j < n {
			dif[j] += val
		}
	}
	cap[0] += dif[0]
	if cap[0] < 0 {
		return false
	}
	for i := 1; i <= end; i++ {
		cap[i] = cap[i-1] + dif[i]
		if cap[i] < 0 {
			return false
		}
	}
	return true
}
func Test_car_pooling(t *testing.T) {
	trips, cap := [][]int{{2, 1, 5}, {3, 3, 7}}, 4
	t.Log(carPooling(trips, cap))
	trips, cap = [][]int{{2, 1, 5}, {3, 3, 7}}, 5
	t.Log(carPooling(trips, cap))
	trips, cap = [][]int{{2, 1, 5}, {3, 5, 7}}, 3
	t.Log(carPooling(trips, cap))
	trips, cap = [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11
	t.Log(carPooling(trips, cap))
	trips, cap = [][]int{{7, 5, 6}, {6, 7, 8}, {10, 1, 6}}, 16
	t.Log(carPooling(trips, cap))
	trips, cap = [][]int{{9, 0, 1}, {3, 3, 7}}, 4
	t.Log(carPooling(trips, cap))
}
