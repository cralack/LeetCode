package largestsumofaverages

import (
	"testing"
)

func largestSumOfAverages(nums []int, m int) float64 {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, m+1)
	}
	//将前 i 个元素
	for i := 1; i <= n; i++ {
		//分成 j 份
		for j := 1; j <= min(i, m); j++ {
			if j == 1 {
				//只分一组则sum/i
				dp[i][1] = float64(preSum[i]) / float64(i)
			} else {
				//枚举最后一个子数组的起点k
				for k := 2; k <= i; k++ {
					dp[i][j] = max(dp[i][j], dp[k-1][j-1]+
						//Avg(k~i)
						float64(preSum[i]-preSum[k-1])/float64(i-k+1))
				}
			}
		}
	}
	return dp[n][m]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
func Test_largest_sum_of_averages(t *testing.T) {
	nums, k := []int{9, 1, 2, 3, 9}, 3
	t.Log(largestSumOfAverages(nums, k))
	nums, k = []int{1, 2, 3, 4, 5, 6, 7}, 4
	t.Log(largestSumOfAverages(nums, k))
	nums, k = []int{4, 1, 7, 5, 6, 2, 3}, 4
	t.Log(largestSumOfAverages(nums, k))
}
