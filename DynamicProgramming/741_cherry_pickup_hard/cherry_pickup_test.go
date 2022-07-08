package cherrypickuphard

import (
	"math"
	"testing"
)

func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}

	}
	dp[0][0] = grid[0][0]
	for sum := 1; sum < 2*n-1; sum++ { //x1+y1=x2+y2=sum
		for x1 := min(sum, n-1); x1 >= max(sum-n+1, 0); x1-- {
			for x2 := min(sum, n-1); x2 >= x1; x2-- {
				y1, y2 := sum-x1, sum-x2 //枚举x1,x2, x2,y2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					dp[x1][x2] = math.MinInt32
					continue
				}

				res := dp[x1][x2] // 都往右
				if x1 > 0 {
					res = max(res, dp[x1-1][x2]) // 往下，往右
				}
				if x2 > 0 {
					res = max(res, dp[x1][x2-1]) // 往右，往下
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[x1-1][x2-1]) // 都往下
				}
				res += grid[x1][y1]
				if x2 != x1 {
					res += grid[x2][y2]
				}
				//走到第n步时,x1、x2状态下最大能摘果数
				dp[x1][x2] = res
			}
		}
	}
	return max(dp[n-1][n-1], 0)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func Test_cherry_pickup(t *testing.T) {
	grid := [][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1}}
	t.Log(cherryPickup(grid))
}
