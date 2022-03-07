package uniquebinarysearchtrees

import "testing"

func numTreesWithMemo(n int) int {
	var count func(lo, hi int) int
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	count = func(lo, hi int) int {
		if lo > hi {
			return 1
		}
		if memo[lo][hi] != 0 {
			return memo[lo][hi]
		}
		res := 0
		for mid := lo; mid <= hi; mid++ {
			left := count(lo, mid-1)
			right := count(mid+1, hi)
			res += left * right
		}
		memo[lo][hi] = res
		return res
	}
	return count(1, n)
}
func numTrees(n int) int {
	var count func(lo, hi int) int
	count = func(lo, hi int) int {
		if lo > hi {
			return 1
		}
		res := 0
		for i := lo; i <= hi; i++ {
			left := count(lo, i-1)
			right := count(i+1, hi)
			res += left * right
		}
		return res
	}
	return count(1, n)
}

func Test_unique_binary_search_trees(t *testing.T) {
	n := 19
	k := numTrees(n)
	t.Log(k)
}

func Benchmark_unique_bst(b *testing.B) {
	b.Run("without memo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			numTrees(19)
		}
		b.StopTimer()
	})
	b.Run("with memo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			numTreesWithMemo(19)
		}
		b.StopTimer()
	})
}
