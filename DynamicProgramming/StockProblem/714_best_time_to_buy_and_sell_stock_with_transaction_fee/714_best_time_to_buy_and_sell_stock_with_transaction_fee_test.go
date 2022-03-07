package besttimetobuyandsellstockwithtransactionfee

import "testing"

func maxProfit(prices []int, fee int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sell, buy := 0, -prices[0]-fee
	for i := 1; i < len(prices); i++ {
		tmp := sell
		sell = max(sell, buy+prices[i])
		buy = max(buy, tmp-prices[i]-fee)
	}
	return sell
}
func Test_best_time_to_buy_and_sell_stock_with_transaction_fee(t *testing.T) {
	stocks := [][]int{{1, 3, 2, 8, 4, 9}, {1, 3, 7, 5, 10, 3}}
	fees := []int{2, 3}
	for i, price := range stocks {
		profit := maxProfit(price, fees[i])
		t.Log(profit)
	}
}
