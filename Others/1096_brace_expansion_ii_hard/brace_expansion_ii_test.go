package braceexpansionii

import (
	"sort"
	"strings"
	"testing"
)

func braceExpansionII(expression string) (ans []string) {
	check := make(map[string]struct{})
	var dfs func(string)
	dfs = func(exp string) {
		//先找到第一个}
		j := strings.Index(exp, "}")
		//base case
		if j == -1 {
			//找不到则exp为单一元素
			check[exp] = struct{}{}
			return
		}
		//j对应的{
		i := strings.LastIndex(exp[:j], "{")
		//前缀和后缀
		a, c := exp[:i], exp[j+1:]
		for _, b := range strings.Split(exp[i+1:j], ",") {
			dfs(a + b + c)
		}
	}

	dfs(expression)
	ans = make([]string, 0, len(check))
	for k := range check {
		ans = append(ans, k)
	}
	sort.Strings(ans)
	return
}
func Test_brace_expansion_ii(t *testing.T) {
	expression := "{a,b}{c,{d,e}}"
	t.Log(braceExpansionII(expression))
	expression = "{{a,z},a{b,c},{ab,z}}"
	t.Log(braceExpansionII(expression))
}
