package find_the_sum_of_subsequence_powers_hard

import (
	"sort"
	"testing"
)

const mod int = 1e9 + 7
const inf = 0x3f3f3f3f

func sumOfPowers(nums []int, k int) (ans int) {
	n := len(nums)
	// dp[end][len][energy]
	dp := make([][]map[int]int, n)
	for i := range dp {
		dp[i] = make([]map[int]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make(map[int]int)
		}
	}
	sort.Ints(nums)

	// 枚举子序列中的最后一个元素
	for end := 0; end < n; end++ {
		dp[end][1][inf] = 1
		// 枚举子序列中end的前一个元素
		for pre := 0; pre < end; pre++ {
			diff := abs(nums[end] - nums[pre])
			// 枚举子序列的长度
			for length := 2; length <= k; length++ {
				for v, cnt := range dp[pre][length-1] {
					dp[end][length][min(diff, v)] = (dp[end][length][min(diff, v)] + cnt) % mod
				}
			}
		}
		for v, cnt := range dp[end][k] {
			ans = (ans + v*cnt%mod) % mod
		}
	}

	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Test_find_the_sum_of_subsequence_powers(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 3, 4}, k: 3, want: 4},
		{nums: []int{2, 2}, k: 2, want: 0},
		{nums: []int{4, 3, -1}, k: 2, want: 10},
	}
	for _, tt := range tests {
		t.Log(sumOfPowers(tt.nums, tt.k) == tt.want)
	}
}
