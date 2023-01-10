package crackingthesafe

import (
	"math"
	"strings"
	"testing"
)

func crackSafe(n int, k int) string {
	mod := int(math.Pow(10, float64(n-1)))
	vis := map[int]bool{}
	ans := &strings.Builder{}
	var dfs func(int)
	dfs = func(node int) {
		for x := 0; x < k; x++ {
			nei := node*10 + x
			if !vis[nei] {
				vis[nei] = true
				v := nei % mod
				dfs(v)
				ans.WriteByte(byte('0' + x))
			}
		}
	}
	dfs(0)
	ans.WriteString(strings.Repeat("0", n-1))
	return ans.String()
}
func Test_cracking_the_safe(t *testing.T) {
	n, k := 1, 2
	t.Log(crackSafe(n, k))
	n, k = 2, 2
	t.Log(crackSafe(n, k))
	n, k = 2, 5
	t.Log(crackSafe(n, k))
	n, k = 4, 2
	t.Log(crackSafe(n, k))
}
