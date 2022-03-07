package besttimetobuyandsellstockiii

import (
	"testing"
)

func MaxProfit(prices []int) int {
	max_k := 2
	dp := make([][][2]int, len(prices))
	for i := range dp {
		dp[i] = make([][2]int, max_k+1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(prices); i++ {
		for k := max_k; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0], dp[i][k][1] = 0, -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][max_k][0]
}

func MaxProfit_2(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dpi10, dpi11 := 0, -99999
	dpi20, dpi21 := 0, -99999
	for _, price := range prices {
		dpi20 = max(dpi20, dpi21+price)
		dpi21 = max(dpi21, dpi10-price)
		dpi10 = max(dpi10, dpi11+price)
		dpi11 = max(dpi11, -price)
	}
	return dpi20
}

func Test_best_time_to_buy_and_sell_stock_iii(t *testing.T) {
	stocks := [][]int{{3, 3, 5, 0, 0, 3, 1, 4},
		{1, 2, 3, 4, 5}, {7, 6, 4, 3, 1}}
	for _, prices := range stocks {
		profit := MaxProfit_2(prices)
		t.Log(profit)
	}
}
