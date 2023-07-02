package weekly_contest

import (
	"sort"
	"testing"
)

/************ 1st test************/
func distinctIntegers(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}
func Test_1st(t *testing.T) {
	n := 5
	t.Log(distinctIntegers(n))
	n = 3
	t.Log(distinctIntegers(n))
}

/************ 2nd test************/
func monkeyMove(n int) int {
	const mod int64 = 1e9 + 7
	ans := fastPow(2, n, mod)
	return int((ans - 2 + mod) % mod)
}

func fastPow(a, b int, mod int64) (ans int64) {
	ans = 1
	for b > 0 {
		if b&1 == 1 {
			ans = (ans * int64(a)) % mod
		}
		a = (a * a) % int(mod)
		b >>= 1
	}
	return
}

func Test_2nd(t *testing.T) {
	n := 3
	t.Log(monkeyMove(n))
	n = 4
	t.Log(monkeyMove(n))
	n = 1e9
	t.Log(monkeyMove(n))
	n = 500000003
	t.Log(monkeyMove(n)) // 1,000,000,006
}

/************ 3rd test************/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func putMarbles(wt []int, k int) (ans int64) {
	// 将序列分成k段，也就是说要找到(k−1)个分割点
	// 每个分割点的价格是分割点左右元素之和
	for i, w := range wt[1:] {
		wt[i] += w
	}
	wt = wt[:len(wt)-1]
	sort.Ints(wt)
	// 后k个大数
	for _, w := range wt[len(wt)-k+1:] {
		ans += int64(w)
	}
	// 前k-1个小数
	for _, w := range wt[:k-1] {
		ans -= int64(w)
	}
	return

}
func Test_3rd(t *testing.T) {
	weights, k := []int{1, 3, 5, 1}, 2
	t.Log(putMarbles(weights, k))
	weights, k = []int{1, 3}, 2
	t.Log(putMarbles(weights, k))
}

/************ 4th test************/
func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	// less[i][j] 表示 [0, j] 范围内比 nums[i] 小的数有多少个，其中 j < i
	less := make([][]int, n)
	for i := 1; i < n; i++ {
		less[i] = make([]int, n)
		for j := 0; j < i; j++ {
			if j > 0 {
				less[i][j] = less[i][j-1]
			}
			if nums[j] < nums[i] {
				less[i][j]++
			}
		}
	}
	// great[i][j] 表示 [j, n) 范围内比 nums[i] 大的数有多少个，其中 i < j
	great := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		great[i] = make([]int, n)
		for j := n - 1; j > i; j-- {
			if j < n-1 {
				great[i][j] = great[i][j+1]
			}
			if nums[i] < nums[j] {
				great[i][j]++
			}
		}
	}

	// 枚举中间两个数, 保证 j < k 并且 nums[k] < nums[j]
	for j := 1; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			if nums[k] < nums[j] {
				ans += int64(less[k][j-1] * great[j][k+1])
			}
		}
	}
	return
}

func Test_4th(t *testing.T) {
	nums := []int{1, 3, 2, 4, 5}
	t.Log(countQuadruplets(nums))
	nums = []int{1, 2, 3, 4}
	t.Log(countQuadruplets(nums))
	nums = []int{1, 10, 2, 11, 12}
	t.Log(countQuadruplets(nums))
}
