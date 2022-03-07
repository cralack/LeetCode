package besttimetobuyandsellstockwithcooldown

import (
	"testing"
)

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(prices); i++ {
		if i-1 == -1 {
			dp[i][0], dp[i][1] = 0, -prices[i]
			continue
		}
		if i-2 == -1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i])
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func Test_best_time_to_buy_and_sell_stock_with_cooldown(t *testing.T) {
	stock := []int{1, 2, 3, 0, 2}
	res := maxProfit(stock)
	t.Log(res)
}
