package maxtaxiearning

import (
	"testing"
)

func maxTaxiEarnings(n int, rides [][]int) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// [end][idx][start/tip]
	ends := make([][][2]int, n+1)
	for _, ride := range rides {
		start, end, tip := ride[0], ride[1], ride[2]
		ends[end] = append(ends[end], [2]int{start, tip})
	}
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1]
		for _, end := range ends[i] {
			start, tip := end[0], end[1]
			dp[i] = max(dp[i], dp[start]+i-start+tip)
		}
	}
	return int64(dp[n])
}

func TestMaxTaxiEarning(t *testing.T) {
	n := 20
	rides := [][]int{{1, 6, 1}, {10, 12, 3}, {13, 18, 1}, {11, 12, 2},
		{12, 15, 2}, {3, 10, 2}}
	// rides := [][]int{}
	ans := maxTaxiEarnings(n, rides)
	t.Log(ans)
}
