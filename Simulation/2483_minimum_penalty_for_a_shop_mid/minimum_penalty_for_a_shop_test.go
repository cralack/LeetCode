package minimum_penalty_for_a_shop_mid

import (
	"testing"
)

func bestClosingTime(customers string) (ans int) {
	// cost := strings.Count(customers, "Y")  // not necessary
	// mincost := cost
	cost, mincost := 0, 0
	for i, customer := range customers {
		if customer == 'Y' {
			cost--
		} else {
			cost++
		}
		if cost < mincost {
			mincost = cost
			ans = i + 1
		}
	}
	return ans
}

func Test_minimum_penalty_for_a_shop(t *testing.T) {
	tests := []struct {
		customers string
	}{
		{customers: "YYNY"},
		{customers: "NNNNN"},
		{customers: "YYYY"},
	}
	for _, test := range tests {
		t.Log(bestClosingTime(test.customers))
	}
}
