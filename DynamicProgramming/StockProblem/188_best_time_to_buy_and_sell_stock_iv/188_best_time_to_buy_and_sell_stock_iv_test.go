package besttimetobuyandsellstockiv

import (
	"testing"
)

func maxProfit(max_k int, prices []int) int {
	size := len(prices)
	if size <= 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// if max_k > size/2 {
	// 	tmp := func(prices []int) int {
	// 		dpi10, dpi11 := 0, -99999
	// 		dpi20, dpi21 := 0, -99999
	// 		for _, price := range prices {
	// 			dpi20 = max(dpi20, dpi21+price)
	// 			dpi21 = max(dpi21, dpi10-price)
	// 			dpi10 = max(dpi10, dpi11+price)
	// 			dpi11 = max(dpi11, -price)
	// 		}
	// 		return dpi20
	// 	}
	// 	return tmp(prices)
	// }

	dp := make([][][2]int, size)
	for i := range dp {
		dp[i] = make([][2]int, max_k+1)
	}
	for i, price := range prices {
		for k := max_k; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0], dp[i][k][1] = 0, -price
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+price)
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-price)
		}
	}
	return dp[len(prices)-1][max_k][0]
}
func Test_best_time_to_buy_and_sell_stock_iv(t *testing.T) {
	prices := [][]int{{2, 4, 1}, {3, 2, 6, 5, 0, 3}}
	ks := []int{2, 2}
	// for i, price := range prices {
	// 	profit := maxProfit(ks[i], price)
	// 	t.Log(profit)
	// }

	t.Log(maxProfit(ks[1], prices[1]))
}
