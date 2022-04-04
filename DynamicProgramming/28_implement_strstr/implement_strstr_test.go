package implementstrstr

import "testing"

func strStr_2D_DP(txt string, pat string) int {
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
func strStr_kmp(txt, pat string) int {
	next := buildNext(pat) //构造next表
	n, i := len(txt), 0    //文本串指针
	m, j := len(pat), 0    //模式串指针
	for j < m && i < n {   //自左向右，逐个比对
		if 0 > j || txt[i] == pat[j] { //若匹配
			i++ //则携手共进
			j++
		} else { //否则,pat右移,txt不回退
			j = next[j]
		}
	}
	if j == m {
		return i - j
	}
	return -1
}
func buildNext(pat string) []int {
	m, j := len(pat), 0    //主串指针
	next := make([]int, m) //next[]表
	next[0] = -1
	t := -1 //模式串指针(p[-1])通配符
	for j < m-1 {
		if 0 > t || pat[j] == pat[t] {
			j++
			t++
			next[j] = t
		} else {
			t = next[t]
		}
	}
	return next
}
func Test_implement_strstr(t *testing.T) {
	t.Log(strStr_kmp("hello", "ll"))
	t.Log(strStr_kmp("aaaaa", "bba"))
	t.Log(strStr_kmp("000100001", "00001"))
	t.Log(strStr_kmp("mississippi", "issipi"))
	t.Log(strStr_2D_DP("", ""))
}
