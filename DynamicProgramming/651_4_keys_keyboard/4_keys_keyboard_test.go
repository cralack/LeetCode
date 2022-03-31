package keyskeyboard

import "testing"

func maxA_rec(n int) int {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res < v {
				res = v
			}
		}
		return res
	}
	memo := make(map[int]map[int]map[int]int, 0)

	var dp func(n, a_num, copy int) int
	dp = func(n, a_num, copy int) int {
		//base case
		if n <= 0 {
			return a_num
		} //memo
		if val, ok := memo[n][a_num][copy]; ok {
			return val
		}
		//if memo need create
		if _, ok := memo[n]; !ok {
			memo[n] = make(map[int]map[int]int)
		}
		if _, ok := memo[n][a_num]; !ok {
			memo[n][a_num] = make(map[int]int)
		}
		memo[n][a_num][copy] = max(dp(n-1, a_num+1, copy), //A
			dp(n-1, a_num+copy, copy), //C-V
			dp(n-2, a_num, a_num))     //C-A C-C
		return memo[n][a_num][copy]
	}
	res := dp(n, 0, 0)
	return res
}

func maxA_ite(n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		//按a
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], dp[j-2]*(i-j+1))
		}
	}
	return dp[n]
}

func Test_4_keys_keyboard(t *testing.T) {
	t.Log(maxA_rec(3))
	t.Log(maxA_rec(7))
	t.Log(maxA_ite(10))
}
func Benchmark_4_keys_keyboard(b *testing.B) {
	b.Run("ite", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxA_ite(10)
		}
		b.StopTimer()
	})
	b.Run("rec", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxA_rec(10)
		}
		b.StopTimer()
	})
}
