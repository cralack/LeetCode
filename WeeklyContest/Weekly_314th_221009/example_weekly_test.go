package weekly_contest

import (
	"testing"
)

/************ 1st test************/
func hardestWorker(n int, logs [][]int) (ans int) {
	cur := 0
	maxDu := 0
	for _, log := range logs {
		id := log[0]
		dura := log[1] - cur
		cur = log[1]
		if maxDu < dura {
			maxDu = dura
			ans = id
		}
		if maxDu == dura && id < ans {
			ans = id
		}
	}
	return
}
func Test_1st(t *testing.T) {
	n := 10
	logs := [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}
	t.Log(hardestWorker(n, logs))
	n = 26
	logs = [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}
	t.Log(hardestWorker(n, logs))
	n = 2
	logs = [][]int{{0, 10}, {1, 20}}
	t.Log(hardestWorker(n, logs))
	n = 70
	logs = [][]int{{36, 3}, {1, 5}, {12, 8}, {25, 9}, {53, 11}, {29, 12}, {52, 14}}
	t.Log(hardestWorker(n, logs))
}

/************ 2nd test************/
func findArray(pref []int) []int {
	n := len(pref)
	ans := make([]int, n)
	ans[0] = pref[0]
	for i := 1; i < n; i++ {
		ans[i] = pref[i-1] ^ pref[i]
	}
	return ans
}
func Test_2nd(t *testing.T) {
	pref := []int{5, 2, 0, 3, 1}
	t.Log(findArray(pref))
	pref = []int{13}
	t.Log(findArray(pref))
}

/************ 3rd test************/
func robotWithString(s string) string {
	n := len(s)
	f := make([]byte, n+1)
	f[n] = 'z' + 1
	for i := n - 1; i >= 0; i-- {
		f[i] = min(f[i+1], s[i])
	}
	ans := []byte{}
	stk := []byte{}
	for i := 0; i < n; i++ {
		stk = append(stk, s[i])
		for len(stk) != 0 && stk[len(stk)-1] <= f[i+1] {
			ans = append(ans, stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}
	}
	return string(ans)
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}
func Test_3rd(t *testing.T) {
	s := "zza"
	t.Log(robotWithString(s))
	s = "bac"
	t.Log(robotWithString(s))
	s = "bdda"
	t.Log(robotWithString(s))
}

/************ 4th test************/
func numberOfPaths(grid [][]int, K int) (ans int) {
	const mod int = 1e9 + 7
	n, m := len(grid), len(grid[0])
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, K)
		}
	}
	dp[0][0][grid[0][0]%K] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < K; k++ {
				if i > 0 { // 从上方走过来
					dp[i][j][k] = (dp[i][j][k] +
						dp[i-1][j][(k-grid[i][j]%K+K)%K]) % mod
				}
				if j > 0 { // 从左方走过来
					dp[i][j][k] = (dp[i][j][k] +
						dp[i][j-1][(k-grid[i][j]%K+K)%K]) % mod
				}
			}
		}
	}
	return dp[n-1][m-1][0]
}
func Test_4th(t *testing.T) {
	grid := [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}
	k := 3
	t.Log(numberOfPaths(grid, k))
	grid = [][]int{{0, 0}}
	k = 2
	t.Log(numberOfPaths(grid, k))
	grid = [][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}}
	k = 1
	t.Log(numberOfPaths(grid, k))
}
