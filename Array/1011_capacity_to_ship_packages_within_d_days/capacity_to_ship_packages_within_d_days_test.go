package capacitytoshippackageswithinddays

import "testing"

func shipWithinDays(weights []int, days int) int {
	lo, hi := 1, 0
	for _, wei := range weights {
		hi += wei
	}
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if valid(weights, days, mi) {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo
}
func valid(weights []int, days, tar int) bool {
	sum, cnt := 0, 1
	for _, wei := range weights {
		if tar < wei {
			return false
		}
		if wei+sum <= tar {
			sum += wei
		} else {
			cnt++
			sum = wei
		}
	}
	return cnt <= days
}
func Test_capacity_to_ship_packages_within_d_days(t *testing.T) {
	weights, days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5
	t.Log(shipWithinDays(weights, days))
	weights, days = []int{3, 2, 2, 4, 1, 4}, 3
	t.Log(shipWithinDays(weights, days))
	weights, days = []int{1, 2, 3, 1, 1}, 4
	t.Log(shipWithinDays(weights, days))
}
