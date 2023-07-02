package freedomtrail

import (
	"math"
	"testing"
)

func findRotateSteps_ite(ring string, key string) int {
	min := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res > v {
				res = v
			}
		}
		return res
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	m, n := len(ring), len(key)
	char2idx := make(map[byte][]int, 0)
	// dp=[m+1][n+1]int
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 记录圆环上字符到索引的映射
	for i := 0; i < m; i++ {
		c := ring[i]
		char2idx[c] = append(char2idx[c], i)
	}
	for j := n - 1; j >= 0; j-- { // 前面的结果要依赖后面的结果
		for i := 0; i < m; i++ {
			c := key[j]
			result := math.MaxInt32
			for _, idx := range char2idx[c] {
				dist1 := abs(idx - i)
				dist2 := m - dist1 // 逆时针转动
				step := min(dist1, dist2)
				// 结果是 从 之前的指针位置转动到现在的位置 的最小值
				result = min(result, step+1+dp[idx][j+1])
			}
			dp[i][j] = result
		}
	}

	return dp[0][0]
}
func findRotateSteps_rec(ring string, key string) int {
	min := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res > v {
				res = v
			}
		}
		return res
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	m, n := len(ring), len(key)
	// 字符 -> 索引列表
	char2idx := make(map[byte][]int, 0)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	// 记录圆环上字符到索引的映射
	for i, c := range ring {
		char2idx[byte(c)] = append(char2idx[byte(c)], i)
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case 完成输入
		if j == len(key) {
			return 0
		}
		// 查找备忘录，避免重叠子问题
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		res := math.MaxInt32
		// ring 上可能有多个字符 key[j]
		for _, k := range char2idx[key[j]] {
			// 拨动指针的次数
			delta := abs(k - i)
			// 选择顺时针还是逆时针
			delta = min(delta, m-delta)
			// 将指针拨到 ring[k]，继续输入 key[j+1..]
			subProblem := dp(k, j+1)
			// 选择「整体」操作次数最少的
			// 加一是因为按动按钮也是一次操作
			res = min(res, 1+delta+subProblem)
		}
		memo[i][j] = res
		return res
	}
	return dp(0, 0)
}
func Test_freedom_trail(t *testing.T) {
	ring := "godding"
	key := "gd"
	t.Log(findRotateSteps_ite(ring, key))
	t.Log(findRotateSteps_rec(ring, key))
}

func Benchmark_freedom_trail(b *testing.B) {
	ring := "goddingasdply"
	b.Run("ite", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findRotateSteps_ite(ring, "gd")
			findRotateSteps_ite(ring, "gdonidg")
		}
		b.StopTimer()
	})
	b.Run("rec", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findRotateSteps_rec(ring, "gd")
			findRotateSteps_rec(ring, "gdonidg")
		}
		b.StopTimer()
	})
}
