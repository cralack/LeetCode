package besttimetobuyandsellstockii

import "testing"

func maxProfit(prices []int) int {
	sell, buy := 0, -999999
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(prices); i++ {
		tmp := sell
		sell = max(sell, buy+prices[i])
		buy = max(buy, tmp-prices[i])
	}
	return sell
}

func Test_best_time_to_buy_and_sell_stock_ii(t *testing.T) {
	prices := [][]int{{7, 1, 5, 3, 6, 4}, {1, 2, 3, 4, 5}, {7, 6, 4, 3, 1}}
	for _, stock := range prices {
		profit := maxProfit(stock)
		t.Log(profit)
	}
}
