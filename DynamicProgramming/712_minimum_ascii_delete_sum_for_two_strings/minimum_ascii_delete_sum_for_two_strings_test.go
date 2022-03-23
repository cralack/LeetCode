package minimumasciideletesumfortwostrings

import "testing"

func minimumDeleteSum(s1 string, s2 string) int {
	min := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res > v {
				res = v
			}
		}
		return res
	}
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	sum1, sum2 := 0, 0
	for i := 1; i <= m; i++ {
		sum1 += int(s1[i-1])
		dp[i][0] = sum1
	}
	for j := 1; j <= n; j++ {
		sum2 += int(s2[j-1])
		dp[0][j] = sum2
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]),
					dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[m][n]
}
func Test_minimum_ascii_delete_sum_for_two_strings(t *testing.T) {
	s1 := "delete"
	s2 := "leet"
	t.Log(minimumDeleteSum(s1, s2))
}
