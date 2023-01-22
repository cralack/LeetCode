package weekly_contest

import (
	"sort"
	"strconv"
	"testing"
)

/************ 1st test************/
func alternateDigitSum(x int) (ans int) {
	str := strconv.Itoa(x)
	op := 1
	for i := 0; i < len(str); i++ {
		cur := int(str[i]-'0') * op
		ans += cur
		op *= -1
	}
	return
}

func Test_1st(t *testing.T) {
	n := 521
	t.Log(alternateDigitSum(n))
	n = 111
	t.Log(alternateDigitSum(n))
	n = 886996
	t.Log(alternateDigitSum(n))
}

/************ 2nd test************/
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}
func Test_2nd(t *testing.T) {
	score, k := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2
	t.Log(sortTheStudents(score, k))
	score, k = [][]int{{3, 4}, {5, 6}}, 0
	t.Log(sortTheStudents(score, k))
}

/************ 3rd test************/
func makeStringsEqual(s string, target string) bool {
	if s == target {
		return true
	}
	sCnt, tCnt := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			sCnt++
		}
		if target[i] == '1' {
			tCnt++
		}
	}

	return sCnt > 0 && tCnt != 0
}

func Test_3rd(t *testing.T) {
	s := "1010"
	target := "0110"
	t.Log(makeStringsEqual(s, target))
	s = "11"
	target = "00"
	t.Log(makeStringsEqual(s, target))
	s = "101110100"
	target = "110011000"
	t.Log(makeStringsEqual(s, target))
}

/************ 4th test************/
func minCost(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := range dp[1:] {
		dp[i+1] = 1 << 30
	}

	for i := 1; i <= n; i++ {
		cnt := make(map[int]int)
		trim := 0
		for j := i - 1; j >= 0; j-- {
			// 计算 trimmed 的值
			cnt[nums[j]]++
			// 出现次数超过一次才对 trimmed 有贡献
			if cnt[nums[j]] == 2 {
				trim += 2
			} else if cnt[nums[j]] > 2 {
				trim++
			}
			dp[i] = min(dp[i], dp[j]+trim+k)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_4th(t *testing.T) {
	nums, k := []int{1, 2, 1, 2, 1, 3, 3}, 2
	t.Log(minCost(nums, k))
	nums, k = []int{1, 2, 1, 2, 1}, 2
	t.Log(minCost(nums, k))
	nums, k = []int{1, 2, 1, 2, 1}, 5
	t.Log(minCost(nums, k))
}
