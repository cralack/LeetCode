package find_x_sum_of_all_k_long_subarrays_i_ez

import (
	"cmp"
	"testing"

	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

type pair struct{ cnt, key int } // 出现次数，元素值

func less(p, q pair) int {
	return cmp.Or(p.cnt-q.cnt, p.key-q.key) // 频率降序 → 同频数值降序
}

func findXSum(nums []int, k, x int) []int64 {
	L := redblacktree.NewWith[pair, struct{}](less) // 维护当前 Top-X 大的 (freq, val)
	R := redblacktree.NewWith[pair, struct{}](less) // 维护其余的

	sumL := 0 // L 的元素和
	cnt := map[int]int{}
	add := func(x int) {
		p := pair{cnt[x], x}
		if p.cnt == 0 {
			return
		}
		if !L.Empty() && less(p, L.Left().Key) > 0 { // p 比 L 中最小的还大
			sumL += p.cnt * p.key
			L.Put(p, struct{}{})
		} else {
			R.Put(p, struct{}{})
		}
	}
	del := func(x int) {
		p := pair{cnt[x], x}
		if p.cnt == 0 {
			return
		}
		if _, ok := L.Get(p); ok {
			sumL -= p.cnt * p.key
			L.Remove(p)
		} else {
			R.Remove(p)
		}
	}
	l2r := func() {
		p := L.Left().Key
		sumL -= p.cnt * p.key
		L.Remove(p)
		R.Put(p, struct{}{})
	}
	r2l := func() {
		p := R.Right().Key
		sumL += p.cnt * p.key
		R.Remove(p)
		L.Put(p, struct{}{})
	}

	ans := make([]int64, len(nums)-k+1)
	for right, in := range nums {
		// 添加 in
		del(in)
		cnt[in]++
		add(in)

		left := right + 1 - k
		if left < 0 {
			continue
		}

		// 维护大小
		for !R.Empty() && L.Size() < x {
			r2l()
		}
		for L.Size() > x {
			l2r()
		}
		ans[left] = int64(sumL)

		// 移除 out
		out := nums[left]
		del(out)
		cnt[out]--
		add(out)
	}
	return ans
}

func Test_find_x_sum_of_all_k_long_subarrays_i(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		x    int
	}{
		{nums: []int{1, 1, 2, 2, 3, 4, 2, 3}, k: 6, x: 2},
		{nums: []int{3, 8, 7, 8, 7, 5}, k: 2, x: 2},
	}
	for _, tt := range tests {
		t.Log(findXSum(tt.nums, tt.k, tt.x))
	}
}
