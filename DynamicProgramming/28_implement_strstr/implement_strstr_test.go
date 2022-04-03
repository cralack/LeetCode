package implementstrstr

import "testing"

func strStr_2d(txt string, pat string) int {
	n, m := len(txt), len(pat)
	if m == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, 256)
	}
	dp[0][pat[0]] = 1
	x := 0
	for j := 1; j < m; j++ {
		for c := 0; c < 256; c++ {
			dp[j][c] = dp[x][c]
		}
		dp[j][pat[j]] = j + 1
		x = dp[x][pat[j]]
	}
	j := 0
	for i := 0; i < n; i++ {
		j = dp[j][txt[i]]
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
func strStr_Next(txt, pat string) int {
	n, m := len(txt), len(pat)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && pat[i] != pat[j] {
			j = pi[j-1]
		}
		if pat[i] == pat[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && txt[i] != pat[j] {
			j = pi[j-1]
		}
		if txt[i] == pat[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
func Test_implement_strstr(t *testing.T) {
	t.Log(strStr_Next("hello", "ll"))
	t.Log(strStr_Next("aaaaa", "bba"))
	t.Log(strStr_Next("000100001", "00001"))
	t.Log(strStr_Next("mississippi", "issipi"))
	t.Log(strStr_2d("", ""))
}
