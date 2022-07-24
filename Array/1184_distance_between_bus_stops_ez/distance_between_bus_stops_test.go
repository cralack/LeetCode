package distancebetweenbusstopsez

import "testing"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	n := len(distance)
	dis, sumDis := 0, 0
	for i := 0; i < n; i++ {
		if start <= i && i < destination {
			dis += distance[i]
		}
		sumDis += distance[i]
	}
	if dis < sumDis-dis {
		return dis
	}
	return sumDis - dis
}
func Test_distance_between_bus_stops(t *testing.T) {
	distance := []int{1, 2, 3, 4}
	start, destination := 0, 1
	t.Log(distanceBetweenBusStops(distance, start, destination))
	distance = []int{1, 2, 3, 4}
	start, destination = 0, 2
	t.Log(distanceBetweenBusStops(distance, start, destination))
	distance = []int{1, 2, 3, 4}
	start, destination = 0, 3
	t.Log(distanceBetweenBusStops(distance, start, destination))
}
