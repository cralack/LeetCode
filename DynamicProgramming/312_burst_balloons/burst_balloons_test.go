package burstballoons

import (
	"testing"
)

func maxCoins(nums []int) int {
	n := len(nums)
	points := make([]int, n+2)    //为气球两边增加边界
	points[0], points[n+1] = 1, 1 //即points={1,气球数组nums,1}
	for i, n := range nums {
		points[i+1] = n
	}
	//dp[i][j] = x表示，戳破气球i和气球j之间
	//（开区间，不包括i和j）的所有气球，可以获得的最高分数为x
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for i := n; i >= 0; i-- { //左边界
		for j := i + 1; j < n+2; j++ { //右边界
			for k := i + 1; k < j; k++ { //最后一个戳破气球k
				//要最后戳破气球k的话得先把开区间(i, k)的气球都戳破
				//再把开区间(k, j)的气球都戳破；最后剩下的气球k
				//相邻的就是气球i和气球j，这时候戳破k的话得到的分数
				//就是points[i]*points[k]*points[j]
				//戳破开区间(i, k)和开区间(k, j)的气球最多能得到的分数就是
				//就是dp[i][k]和dp[k][j]
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+
					points[i]*points[j]*points[k])
			}
		}
	}
	return dp[0][n+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Test_burst_balloons(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	t.Log(maxCoins(nums))
}
