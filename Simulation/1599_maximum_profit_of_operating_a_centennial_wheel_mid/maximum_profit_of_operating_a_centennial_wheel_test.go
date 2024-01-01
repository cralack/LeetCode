package maximum_profit_of_operating_a_centennial_wheel_mid

import (
	"testing"
)

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) (ans int) {
	ans = -1
	wait, i := 0, 0
	totalProfit, maxProfit := 0, 0
	for wait > 0 || i < len(customers) {
		if i < len(customers) {
			wait += customers[i]
		}
		up := min(4, wait)
		wait -= up
		totalProfit += up*boardingCost - runningCost
		i++
		if totalProfit > maxProfit {
			ans = i
			maxProfit = totalProfit
		}
	}

	return
}

func Test_maximum_profit_of_operating_a_centennial_wheel(t *testing.T) {
	tests := []struct {
		customers    []int
		boardingCost int
		runningCost  int
	}{
		{customers: []int{8, 3}, boardingCost: 5, runningCost: 6},
		{customers: []int{10, 9, 6}, boardingCost: 6, runningCost: 4},
		{customers: []int{3, 4, 0, 5, 1}, boardingCost: 1, runningCost: 92},
	}
	for _, tt := range tests {
		t.Log(minOperationsMaxProfit(tt.customers, tt.boardingCost, tt.runningCost))
	}
}
