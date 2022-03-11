package corporateflightbookings

import "testing"

func corpFlightBookings(bookings [][]int, n int) []int {
	dif := make([]int, n)
	for _, update := range bookings {
		i, j := update[0], update[1]
		val := update[2]
		dif[i-1] += val
		if j < n {
			dif[j] -= val
		}
	}
	nums := make([]int, n)
	nums[0] = dif[0]
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] + dif[i]
	}
	return nums
}
func Test_corporate_flight_bookings(t *testing.T) {
	bookings, n := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5
	t.Log(corpFlightBookings(bookings, n))
	bookings, n = [][]int{{1, 2, 10}, {2, 2, 15}}, 2
	t.Log(corpFlightBookings(bookings, n))
}
