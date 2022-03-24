package longestpalindromicsubsequence

import (
	"testing"
)

func longestPalindromeSubseq_V1(s string) int {
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

func Benchmark_longest_palindromic_subsequence(b *testing.B) {
	s1 := "bbbab"
	s2 := "cbbd"
	s3 := "aabaa"
	b.Run("2v", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			longestPalindromeSubseq_V1(s1)
			longestPalindromeSubseq_V1(s2)
			longestPalindromeSubseq_V1(s3)
		}
		b.StopTimer()
	})
	b.Run("1v", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			longestPalindromeSubseq(s1)
			longestPalindromeSubseq(s2)
			longestPalindromeSubseq(s3)
		}
		b.StopTimer()
	})
}
