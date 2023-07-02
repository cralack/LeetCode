package stonegame

import "testing"

type ele struct {
	First, Second int
}

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]ele, n)
	for i := range dp {
		dp[i] = make([]ele, n)
		dp[i][i].First = piles[i]
		dp[i][i].Second = 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 先手选择最左边或最右边的分数
			left := piles[i] + dp[i+1][j].Second
			right := piles[j] + dp[i][j-1].Second
			// 套用状态转移方程
			// 先手肯定会选择更大的结果，后手的选择随之改变
			if left > right {
				dp[i][j].First = left
				dp[i][j].Second = dp[i+1][j].First
			} else {
				dp[i][j].First = right
				dp[i][j].Second = dp[i][j-1].First
			}
		}
	}
	res := dp[0][n-1]
	return res.First-res.Second > 0

	// 只要先手，必胜
	// return true
}
func Test_stone_game(t *testing.T) {
	piles := []int{5, 3, 4, 5}
	t.Log(stoneGame(piles))
}
