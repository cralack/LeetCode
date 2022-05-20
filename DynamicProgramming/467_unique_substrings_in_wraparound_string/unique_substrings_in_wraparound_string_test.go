package uniquesubstringsinwraparoundstring

import "testing"

func findSubstringInWraproundString(p string) (ans int) {
	dp := [26]int{}
	k := 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { // 字符之差为 1 或 -25
			k++
		} else {
			k = 1
		}
		dp[ch-'a'] = max(dp[ch-'a'], k)
	}
	for _, v := range dp {
		ans += v
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_unique_substrings_in_wraparound_string(t *testing.T) {
	p := "a"
	t.Log(findSubstringInWraproundString(p))
	p = "cac"
	t.Log(findSubstringInWraproundString(p))
	p = "zab"
	t.Log(findSubstringInWraproundString(p))
}
