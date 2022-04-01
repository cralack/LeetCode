package implementstrstr

import "testing"

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, 256)
	}
	dp[0][needle[0]] = 1
	x := 0
	for j := 1; j < m; j++ {
		for c := 0; c < 256; c++ {
			dp[j][c] = dp[x][c]
		}
		dp[j][needle[j]] = j + 1
		x = dp[x][needle[j]]
	}
	j := 0
	for i := 0; i < n; i++ {
		j = dp[j][haystack[i]]
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
func Test_implement_strstr(t *testing.T) {
	t.Log(strStr("hello", "ll"))
	t.Log(strStr("aaaaa", "bba"))
	t.Log(strStr("", ""))
}
