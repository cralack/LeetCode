package regularexpressionmatching

import (
	"strconv"
	"testing"
)

func isMatch_rec(txt string, pat string) bool {
	memo := make(map[string]bool)
	m, n := len(txt), len(pat)
	var DP func(i, j int) bool
	DP = func(i, j int) bool {
		//base case
		if j == n {
			return i == m
		}
		if i == m {
			// 如果能匹配空串，一定是字符和 * 成对儿出现
			if (n-j)%2 == 1 {
				return false
			} // 检查是否为 x*y*z* 这种形式
			for ; j+1 < n; j += 2 {
				if pat[j+1] != '*' {
					return false
				}
			}
			return true
		}
		// 记录状态 (i, j)，消除重叠子问题
		key := strconv.Itoa(i) + "," + strconv.Itoa(j)
		if res, ok := memo[key]; ok {
			return res
		}
		res := false

		if txt[i] == pat[j] || pat[j] == '.' { // 匹配
			if j < len(pat)-1 && pat[j+1] == '*' {
				// 1.1 通配符匹配 0 次或多次
				res = DP(i, j+2) ||
					DP(i+1, j)
			} else {
				// 1.2 常规匹配 1 次
				res = DP(i+1, j+1)
			}
		} else {
			// 不匹配
			if j < len(pat)-1 && pat[j+1] == '*' {
				// 2.1 通配符匹配 0 次
				res = DP(i, j+2)
			} else {
				// 2.2 无法继续匹配
				res = false
			}
		}
		memo[key] = res
		return res
	}
	return DP(0, 0)
}

func isMatch_ite(txt, pat string) bool {
	m, n := len(txt), len(pat)
	match := func(i, j int) bool {
		if i == 0 { //base case
			return false
		}
		if pat[j-1] == '.' { // 通配
			return true
		}
		return txt[i-1] == pat[j-1] //字符匹配
	}
	//init
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	//base case
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if pat[j-1] == '*' { // 1.1 通配符匹配 0 次或多次
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(i, j-1) { //通配状态延续一次
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(i, j) { // 1.2 常规匹配 1 次
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
func Test_regular_expression_matching(t *testing.T) {
	t.Log(isMatch_rec("aa", "a"))
	t.Log(isMatch_rec("aa", "a*"))
	t.Log(isMatch_rec("ab", ".*"))
	t.Log(isMatch_ite("zaaab", ".a*b"))
}

func Benchmark_ismatch(b *testing.B) {
	txt, pat := "abaaac", "ab.a*c"
	b.Run("ite", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			isMatch_ite(txt, pat)
		}
		b.StopTimer()
	})
	b.Run("rec", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			isMatch_rec(txt, pat)
		}
		b.StopTimer()
	})
}
