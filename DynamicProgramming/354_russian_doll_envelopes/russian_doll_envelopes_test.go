package russiandollenvelopes

import (
	"sort"
	"testing"
)

func maxEnvelopes_DP(envelopes [][]int) int {
	// 先对宽度 w 进行升序排序
	// 如果遇到 w 相同的情况
	// 则按照高度 h 降序排序
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] ==
			envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}
	n := len(envelopes)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	// 之后把所有的 h 作为一个数组
	// 在这个数组上计算 LIS 的长度就是答案
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	// 首先，对宽度 w 从小到大排序，确保了 w 这个维度可以互相嵌套，
	// 所以我们只需要专注高度 h 这个维度能够互相嵌套即可。
	// 其次，两个 w 相同的信封不能相互包含，所以对于宽度 w 相同的信封，
	// 对高度 h 进行降序排序，保证 LIS 中不存在多个 w 相同的信封。
	return max(dp...)
}

func maxEnvelopes_bSearch(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] ==
			envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
func Test_russian_doll_envelopes(t *testing.T) {
	envelopes := [][]int{{5, 4}, {6, 4}, {1, 8},
		{6, 7}, {2, 3}, {5, 2}}
	t.Log(maxEnvelopes_bSearch(envelopes))
	envelopess := [][]int{}
	for i := 1; i <= 100000; i++ {
		envelopess = append(envelopess, []int{i, i})
	}
	t.Log(maxEnvelopes_DP(envelopess))
	t.Log(maxEnvelopes_bSearch(envelopess))
}

func Benchmark_russian_doll_envelopes(b *testing.B) {
	envelopes := [][]int{}
	for i := 1; i <= 1000; i++ {
		envelopes = append(envelopes, []int{i, i})
	}
	b.Run("DP", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxEnvelopes_DP(envelopes)
		}
		b.StopTimer()
	})
	b.Run("bSearch", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxEnvelopes_bSearch(envelopes)
		}
		b.StopTimer()
	})
}
