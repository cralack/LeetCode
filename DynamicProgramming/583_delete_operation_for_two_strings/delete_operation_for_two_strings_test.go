package deleteoperationfortwostrings

import "testing"

func maxDistance(word1 string, word2 string) int {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res < v {
				res = v
			}
		}
		return res
	}
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// len(lcs)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return m + n - 2*dp[n][m]
}
func Test_delete_operation_for_two_strings(t *testing.T) {
	word1 := "etco"
	word2 := "leetcode"
	t.Log(maxDistance(word1, word2))
	word1 = "a"
	word2 = "b"
	t.Log(maxDistance(word1, word2))
}
