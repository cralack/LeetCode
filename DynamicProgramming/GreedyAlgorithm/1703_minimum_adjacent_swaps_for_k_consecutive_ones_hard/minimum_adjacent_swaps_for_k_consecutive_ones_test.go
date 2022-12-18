package minimumadjacentswapsforkconsecutiveones

import (
	"math"
	"testing"
)

func minMoves(nums []int, k int) int {
	p := []int{}
	for i, v := range nums {
		if v != 0 {
			p = append(p, i-len(p))
		}
	}
	m := len(p)
	s := make([]int, m+1) // p 的前缀和
	for i, v := range p {
		s[i+1] = s[i] + v
	}
	ans := math.MaxInt
	for i, v := range s[:m-k+1] { // p[i] 到 p[i+k-1] 中所有数到 p[i+k/2] 的距离之和，取最小值
		ans = min(ans, v+s[i+k]-s[i+k/2]*2-p[i+k/2]*(k%2))
	}
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func Test_minimum_adjacent_swaps_for_k_consecutive_ones(t *testing.T) {
	nums, k := []int{1, 0, 0, 1, 0, 1}, 2
	t.Log(minMoves(nums, k))
	nums, k = []int{1, 0, 0, 0, 0, 0, 1, 1}, 3
	t.Log(minMoves(nums, k))
	nums, k = []int{1, 1, 0, 1}, 2
	t.Log(minMoves(nums, k))
}
