package longestpalindromicsubsequence

import (
	"testing"
)

func longestPalindromeSubseq(s string) int {
	dp := make([]int, len(s))
	for i := range s {
		rightOver := dp[i]
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			tmp := dp[j]
			if s[i] == s[j] {
				dp[j] = rightOver + 2
			} else if dp[j+1] > dp[j] {
				dp[j] = dp[j+1]
			}
			rightOver = tmp
		}
	}
	return dp[0]
}
func Test_longest_palindromic_subsequence(t *testing.T) {
	s := "bbbab"
	t.Log(longestPalindromeSubseq(s))
	s = "cbbd"
	t.Log(longestPalindromeSubseq(s))
	s = "aabaa"
	t.Log(longestPalindromeSubseq(s))
}
