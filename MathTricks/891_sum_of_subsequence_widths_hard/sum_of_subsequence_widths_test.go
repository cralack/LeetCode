package sumofsubsequencewidths

import (
	"sort"
	"testing"
)

func sumSubseqWidths(nums []int) (ans int) {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	n := len(nums)
	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		// 预处理 2 的幂次
		pow2[i] = pow2[i-1] * 2 % mod
	}
	for i, x := range nums {
		// 在题目的数据范围下，这不会溢出
		ans += (pow2[i] - pow2[n-1-i]) * x
	}
	// 注意上面有减法，ans 可能为负数
	return (ans%mod + mod) % mod
}

func Test_sum_of_subsequence_widths(t *testing.T) {
	nums := []int{2, 1, 3}
	t.Log(sumSubseqWidths(nums))
	nums = []int{2}
	t.Log(sumSubseqWidths(nums))
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Log(sumSubseqWidths(nums))
}
