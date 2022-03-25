package coinchange2

import "testing"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i-coin >= 0 {
				dp[i] = dp[i] + dp[i-coin]
			}
		}
	}
	return dp[amount]
}
func Test_coin_change_2(t *testing.T) {
	amount := 5
	coins := []int{1, 2, 5}
	t.Log(change(amount, coins))
}
