package longestcommonsubsequence

import (
	"testing"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res < v {
				res = v
			}
		}
		return res
	}
	//init
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	//dp
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
func Test_longest_common_subsequence(t *testing.T) {
	text1 := "abcde" //ace
	text2 := "ace"
	t.Log(longestCommonSubsequence(text1, text2))
	text1 = "zabcde" //ace
	text2 = "acez"
	t.Log(longestCommonSubsequence(text1, text2))
	text1 = "psnw" //s
	text2 = "vozsh"
	t.Log(longestCommonSubsequence(text1, text2))

}
