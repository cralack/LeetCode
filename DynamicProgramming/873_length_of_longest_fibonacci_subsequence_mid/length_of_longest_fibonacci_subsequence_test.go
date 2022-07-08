package lengthoflongestfibonaccisubsequencemid

import "testing"

func lenLongestFibSubseq_DP_2pointer(arr []int) int {
	l, max := len(arr), 0
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
	}
	left, right := 0, 0
	for i := 2; i < l; i++ {
		left, right = 0, i-1
		for left < right {
			// sum = arr[left] + arr[right]
			if arr[left]+arr[right] > arr[i] {
				right--
			} else if arr[left]+arr[right] < arr[i] {
				left++
			} else {
				dp[right][i] = dp[left][right] + 1
				if dp[right][i] > max {
					max = dp[right][i]
				}
				right--
				left++
			}
		}
	}
	if max > 0 {
		max += 2
	}
	return max
}

func lenLongestFibSubseq_DPonly(arr []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(arr)
	val2idx := make(map[int]int, n)
	dp := make([][]int, n)
	for idx, val := range arr {
		val2idx[val] = idx
		dp[idx] = make([]int, n)
	}
	for i, x := range arr {
		for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
			if k, ok := val2idx[x-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
			}
		}
	}
	return
}

func Test_length_of_longest_fibonacci_subsequence(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t.Log(lenLongestFibSubseq_DP_2pointer(arr))
	arr = []int{1, 3, 7, 11, 12, 14, 18}
	t.Log(lenLongestFibSubseq_DPonly(arr))
}

func Benchmark_lolfibseq(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b.Run("DPonly", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lenLongestFibSubseq_DPonly(arr)
		}
		b.StopTimer()
	})
	b.Run("DP_2pointer", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lenLongestFibSubseq_DP_2pointer(arr)
		}
		b.StopTimer()
	})
}
