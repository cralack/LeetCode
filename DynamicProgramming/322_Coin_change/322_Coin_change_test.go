package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	coins := [][]int{{1, 2, 5}, {2}, {1}, {1}, {1}}
	amount := []int{11, 3, 0, 1, 2}
	for i, v := range coins {
		n := coinChange(v, amount[i])
		t.Log(n)
	}
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) //初始化dp数组
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1 //初始值设为大值
	}

	dp[0] = 0 //base case
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(dp); i++ { //遍历所有状态取值
		for _, coin := range coins { //遍历所有选择
			if i-coin < 0 { //子问题无解则跳过
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
