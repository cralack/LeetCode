package longestpalindromicsubsequence

import (
	"testing"
)

func longestPalindromeSubseq(s string) int {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res < v {
				res = v
			}
		}
		return res
	}
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	//i :end->start
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
func Test_longest_palindromic_subsequence(t *testing.T) {
	s := "bbbab"
	t.Log(longestPalindromeSubseq(s))
	s = "cbbd"
	t.Log(longestPalindromeSubseq(s))
	s = "aabaa"
	t.Log(longestPalindromeSubseq(s))
}
