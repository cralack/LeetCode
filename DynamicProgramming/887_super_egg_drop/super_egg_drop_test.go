package supereggdrop

import (
	"math"
	"testing"
)

func superEggDrop_rec(k int, n int) int {
	memo := make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -777
		}
	} //K个鸡蛋,N层楼最差需要试错几次
	var dp func(k, n int) int
	dp = func(k, n int) int {
		//base case
		if n == 0 {
			return 0
		}
		if k == 1 {
			return n
		}
		if memo[k][n] != -777 {
			return memo[k][n]
		}
		res := math.MaxInt32
		for i := 1; i <= n; i++ {
			res = min(res,
				max(dp(k, n-i),
					dp(k-1, i-1))+1)
		}
		memo[k][n] = res
		return res
	}
	res := dp(k, n)
	return res
}
func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if res > v {
			res = v
		}
	}
	return res
}
func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if res < v {
			res = v
		}
	}
	return res
}
func superEggDrop_ite(k int, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[j][i] = dp[j][i-1] + dp[j-1][i-1] + 1
			if dp[j][i] >= n {
				return i
			}
		}
	}
	return n
}
func Test_super_egg_drop(t *testing.T) {
	k, n := 1, 2
	t.Log(superEggDrop_rec(k, n))
	t.Log(superEggDrop_ite(2, 6))
	t.Log(superEggDrop_ite(3, 14))
}

func Benchmark_super_egg_drop(b *testing.B) {
	k, n := 5, 1000
	b.Run("ite", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			superEggDrop_ite(k, n)
		}
		b.StopTimer()
	})
	b.Run("rec", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			superEggDrop_rec(k, n)
		}
		b.StopTimer()
	})
}
