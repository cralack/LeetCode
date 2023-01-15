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
func rangeAddQueries(n int, queries [][]int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for _, query := range queries {
		row1, col1, row2, col2 := query[0], query[1], query[2], query[3]
		for i := row1; i <= row2; i++ {
			for j := col1; j <= col2; j++ {
				matrix[i][j]++
			}
		}
	}
	return matrix
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
	left, right := 0, 0
	for ; right < n; right++ {
		pair += cnt[nums[right]]
		cnt[nums[right]]++
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
