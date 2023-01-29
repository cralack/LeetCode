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
	t.Log(monkeyMove(n)) //1,000,000,006
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
	//将序列分成k段，也就是说要找到(k−1)个分割点
	//每个分割点的价格是分割点左右元素之和
	for i, w := range wt[1:] {
		wt[i] += w
	}
	wt = wt[:len(wt)-1]
	sort.Ints(wt)
	//后k个大数
	for _, w := range wt[len(wt)-k+1:] {
		ans += int64(w)
	}
	//前k-1个小数
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
	//不限制1<=nums[i]<=n的话可能需要单调栈来进行预处理
	//需要知道k右边有多少个比nums[j]大的数
	//		j左边有多少个比nums[k]小的数
	n := len(nums)
	//预处理great[k][x]定义成在k右边的比x大的数的个数
	great := make([][]int, n)
	great[n-1] = make([]int, n+1)
	//倒序枚举k
	for k := n - 2; k >= 2; k-- {
		great[k] = append([]int(nil), great[k+1]...)
		for x := nums[k+1] - 1; x > 0; x-- {
			great[k][x]++ // x < nums[k+1]，对于 x，大于它的数的个数 +1
		}
	}
	//预处理less[x]定义成在j左边的比x小的数的个数
	less := make([]int, n+1)
	for j := 1; j < n-2; j++ {
		for x := nums[j-1] + 1; x <= n; x++ {
			less[x]++ // x > nums[j-1]，对于 x，小于它的数的个数 +1
		}
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				ans += int64(less[nums[k]] * great[k][nums[j]])
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
