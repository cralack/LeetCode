package flipstringtomonotoneincreasing

import "testing"

func minFlipsMonoIncr(s string) int {
	n := len(s)
	dp := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		if s[i-1] == '1' { // s[i-1]=='1'的话 dp[0]+1
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		} else { // s[i-1]=='0'的话 min(dp)+1
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		}
	}
	return min(dp[n][0], dp[n][1])
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_flip_string_to_monotone_increasing(t *testing.T) {
	s := "00110"
	t.Log(minFlipsMonoIncr(s))
	s = "010110"
	t.Log(minFlipsMonoIncr(s))
	s = "110110"
	t.Log(minFlipsMonoIncr(s))
	s = "00011000"
	t.Log(minFlipsMonoIncr(s))
}
