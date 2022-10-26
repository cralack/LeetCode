package shortestsubarraywithsumatleastk

import (
	"testing"
)

func shortestSubarray(nums []int, k int) (ans int) {
	n := len(nums)
	ans = n + 10
	preSum := make([]int, n+1)
	for i, n := range nums {
		preSum[i+1] = preSum[i] + n
	}
	que := make([]int, 0) //存放idx
	for i, curSum := range preSum {
		//满足curSum-preSum[head]>=k的情况下收缩队首
		for len(que) > 0 && curSum-preSum[que[0]] >= k {
			ans = min(ans, i-que[0])
			que = que[1:]
		}
		//保持单调递增
		for len(que) > 0 && preSum[que[len(que)-1]] >= curSum {
			que = que[:len(que)-1]
		}
		que = append(que, i)
	}
	if ans > n {
		return -1
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_shortest_subarray_with_sum_at_least_k(t *testing.T) {
	nums, k := []int{1}, 1
	t.Log(shortestSubarray(nums, k))
	nums, k = []int{1, 2}, 4
	t.Log(shortestSubarray(nums, k))
	nums, k = []int{2, -1, 2}, 3
	t.Log(shortestSubarray(nums, k))
	nums, k = []int{17, 85, 93, -45, -21}, 150
	t.Log(shortestSubarray(nums, k))
	nums, k = []int{84, -37, 32, 40, 95}, 167
	t.Log(shortestSubarray(nums, k))
}
