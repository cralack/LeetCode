package weekly_contest

import (
	"testing"
)

/************ 1st test************/
func differenceOfSum(nums []int) (ans int) {
	eleSum, numSum := 0, 0
	for _, n := range nums {
		eleSum += n
		for n > 0 {
			numSum += n % 10
			n /= 10
		}
	}
	return eleSum - numSum
}

func Test_1st(t *testing.T) {
	nums := []int{1, 15, 6, 3}
	t.Log(differenceOfSum(nums))
	nums = []int{1, 2, 3, 4}
	t.Log(differenceOfSum(nums))
}

/************ 2nd test************/
func rangeAddQueries(n int, queries [][]int) (ans [][]int) {
	// 二维差分解法
	m := n

	// 二维差分模板
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, m+1)
	}
	// 本题中x为1
	update := func(r1, c1, r2, c2, x int) {
		diff[r1][c1] += x
		diff[r1][c2+1] -= x
		diff[r2+1][c1] -= x
		diff[r2+1][c2+1] += x // 前面一共减了两次x,这里再补一个x
	}
	for _, q := range queries {
		update(q[0], q[1], q[2], q[3], 1)
	}

	// 用二维前缀和复原
	ans = make([][]int, n+1)
	ans[0] = make([]int, m+1)
	for i := 0; i < n; i++ {
		ans[i+1] = make([]int, m+1)
		for j := 0; j < m; j++ {
			ans[i+1][j+1] = ans[i+1][j] + ans[i][j+1] - ans[i][j] +
				diff[i][j]
		}
	}

	// shrink
	ans = ans[1:]
	for i, row := range ans {
		ans[i] = row[1:]
	}
	return
}
func Test_2nd(t *testing.T) {
	queries, n := [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}, 3
	t.Log(rangeAddQueries(n, queries))
	queries, n = [][]int{{0, 0, 1, 1}}, 2
	t.Log(rangeAddQueries(n, queries))
}

/************ 3rd test************/
func countGood(nums []int, k int) (ans int64) {
	cnt := make(map[int]int, 0)
	pair := 0
	n := len(nums)
	left := 0
	for right, num := range nums {
		pair += cnt[num]
		cnt[num]++
		for pair >= k {
			ans += int64(n - right)
			cnt[nums[left]]--
			pair -= cnt[nums[left]]
			left++
		}
	}
	return
}

func Test_3rd(t *testing.T) {
	nums, k := []int{1, 1, 1, 1, 1}, 10
	t.Log(countGood(nums, k))
	nums, k = []int{3, 1, 4, 3, 2, 2, 4}, 2
	t.Log(countGood(nums, k))
	nums, k = []int{2, 1, 3, 1, 2, 2, 3, 3, 2, 2, 1, 1, 1, 3, 1}, 11
	t.Log(countGood(nums, k))

}

/************ 4th test************/

func Test_4th(t *testing.T) {
}
