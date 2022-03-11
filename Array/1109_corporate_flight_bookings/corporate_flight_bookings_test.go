package corporateflightbookings

import "testing"

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	return nums
}
func Test_corporate_flight_bookings(t *testing.T) {
	bookings, n := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5
	t.Log(corpFlightBookings(bookings, n))
	bookings, n = [][]int{{1, 2, 10}, {2, 2, 15}}, 2
	t.Log(corpFlightBookings(bookings, n))
}
