package shortestcommonsupersequence

import (
	"testing"
)

// 能通过的测试数据更多，但超内存了，还需要进一步优化
func shortestCommonSupersequence_memo(s, t string) string {
	n, m := len(s), len(t)
	memo := make([][]string, n)
	for i := range memo {
		memo[i] = make([]string, m)
	}
	// dfs(i,j) 返回 s 的前 i 个字母和 t 的前 j 个字母的最短公共超序列
	var dfs func(int, int) string
	dfs = func(i, j int) string {
		if i < 0 { // s 是空串，返回剩余的 t
			return t[:j+1]
		}
		if j < 0 { // t 是空串，返回剩余的 s
			return s[:i+1]
		}
		p := &memo[i][j]
		if *p != "" {
			return *p // 避免重复计算 dfs 的结果
		}
		if s[i] == t[j] { // 最短公共超序列一定包含 s[i]
			*p = dfs(i-1, j-1) + string(s[i])
		} else {
			ans1 := dfs(i-1, j)
			ans2 := dfs(i, j-1)
			if len(ans1) < len(ans2) { // 取 ans1 和 ans2 中更短的组成答案
				*p = ans1 + string(s[i])
			} else {
				*p = ans2 + string(t[j])
			}
		}
		return *p
	}
	return dfs(n-1, m-1)
}
func shortestCommonSupersequence_memo_plus(s, t string) string {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
	}
	// dfs(i,j) 返回 s 的前 i 个字母和 t 的前 j 个字母的最短公共超序列的长度
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 { // s 是空串，返回剩余的 t 的长度
			return j + 1
		}
		if j < 0 { // t 是空串，返回剩余的 s 的长度
			return i + 1
		}
		p := &memo[i][j]
		if *p > 0 {
			return *p // 避免重复计算 dfs 的结果
		}
		if s[i] == t[j] { // 最短公共超序列一定包含 s[i]
			*p = dfs(i-1, j-1) + 1
		} else {
			*p = min(dfs(i-1, j), dfs(i, j-1)) + 1
		}
		return *p
	}

	// makeAns(i,j) 返回 s 的前 i 个字母和 t 的前 j 个字母的最短公共超序列
	// 看上去和 dfs 没啥区别，但是末尾的递归是 if-else
	// makeAns(i-1,j) 和 makeAns(i,j-1) 不会都调用
	// 所以 makeAns 的递归树仅仅是一条链
	var makeAns func(int, int) string
	makeAns = func(i, j int) string {
		if i < 0 { // s 是空串，返回剩余的 t
			return t[:j+1]
		}
		if j < 0 { // t 是空串，返回剩余的 s
			return s[:i+1]
		}
		if s[i] == t[j] { // 最短公共超序列一定包含 s[i]
			return makeAns(i-1, j-1) + string(s[i])
		}
		// 如果下面 if 成立，说明上面 dfs 中的 min 取的是 dfs(i - 1, j)
		// 说明 dfs(i - 1, j) 对应的公共超序列更短
		// 那么就在 makeAns(i - 1, j) 的结果后面加上 s[i]
		// 否则说明 dfs(i, j - 1) 对应的公共超序列更短
		// 那么就在 makeAns(i, j - 1) 的结果后面加上 t[j]
		if dfs(i, j) == dfs(i-1, j)+1 {
			return makeAns(i-1, j) + string(s[i])
		}
		return makeAns(i, j-1) + string(t[j])
	}

	return makeAns(n-1, m-1)
}

func shortestCommonSupersequence_dp(s, t string) string {
	n, m := len(s), len(t)
	// f[i+1][j+1] 表示 s 的前 i 个字母和 t 的前 j 个字母的最短公共超序列的长度
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j < m; j++ {
		f[0][j] = j // 递归边界
	}
	for i := 1; i < n; i++ {
		f[i][0] = i // 递归边界
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i] == t[j] { // 最短公共超序列一定包含 s[i]
				f[i+1][j+1] = f[i][j] + 1
			} else { // 取更短的组成答案
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
			}
		}
	}

	ans := []byte{}
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if s[i] == t[j] { // 公共超序列一定包含 s[i]
			ans = append(ans, s[i])
			i--
			j-- // 相当于继续递归 makeAns(i - 1, j - 1)
		} else if f[i+1][j+1] == f[i][j+1]+1 {
			ans = append(ans, s[i])
			i-- // 相当于继续递归 makeAns(i - 1, j)
		} else {
			ans = append(ans, t[j])
			j-- // 相当于继续递归 makeAns(i, j - 1)
		}
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	// 补上前面的递归边界
	return s[:i+1] + t[:j+1] + string(ans)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func Test_shortest_common_supersequence(t *testing.T) {
	tests := []struct {
		str1, str2 string
	}{
		{"abac", "cab"},
	}
	for _, tt := range tests {
		t.Log(shortestCommonSupersequence_dp(tt.str1, tt.str2))
	}
}
