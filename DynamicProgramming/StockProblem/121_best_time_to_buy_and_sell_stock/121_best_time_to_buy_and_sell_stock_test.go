package besttimetobuyandsellstock

import (
	"testing"
)

func maxProfit(prices []int) int {
	size := len(prices)
	// dp := [size][2]int{}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sell, buy := 0, -999999
	for i := 0; i < size; i++ {
		//今天未持有 1.昨天未持有 今天rest 2.昨天持有 今天sell
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		sell = max(sell, buy+prices[i])
		//今天持有 1.昨天持有 今天rest 2.昨天未持有 今天buy
		// dp[i][1] = max(dp[i-1][1], -prices[i])
		buy = max(buy, -prices[i])
	}
	return sell
}
func Test_best_time_to_buy_and_sell_stock(t *testing.T) {
	stocks := [][]int{{7, 1, 5, 3, 6, 4}, {7, 6, 4, 3, 1}}
	for _, stock := range stocks {
		profit := maxProfit(stock)
		t.Log(profit)
	}
}
