package weekly_contest

import (
	"sort"
	"testing"
)

/************ 1st test************/
func checkDistances(s string, distance []int) bool {
	idx := make(map[rune]int, 0)
	for i, char := range s {
		if _, has := idx[char]; !has {
			idx[char] = i
		} else {
			dis := i - idx[char] - 1
			if distance[char-'a'] != dis {
				return false
			}
		}
	}
	return true
}
func Test_1st(t *testing.T) {
	s := "abaccb"
	dis := []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	t.Log(checkDistances(s, dis))
	s = "aa"
	dis = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	t.Log(checkDistances(s, dis))
}

/************ 2nd test************/
func numberOfWays(startPos int, endPos int, k int) int {
	const MOD int = 1e9 + 7
	dis := abs(startPos - endPos)
	if (dis+k)%2 == 1 || dis > k {
		return 0
	}
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 0; i <= k; i++ {
		dp[i][0] = 1
		for j := 1; j <= i; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i-1][j-1]) % MOD
		}
	}
	return dp[k][(dis+k)/2]
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func Test_2nd(t *testing.T) {
	startPos := 1
	endPos := 2
	k := 3
	t.Log(numberOfWays(startPos, endPos, k))
	startPos = 2
	endPos = 5
	k = 10
	t.Log(numberOfWays(startPos, endPos, k))
	startPos = 1
	endPos = 1000
	k = 999
	t.Log(numberOfWays(startPos, endPos, k))
	startPos = 989
	endPos = 1000
	k = 99
	t.Log(numberOfWays(startPos, endPos, k))
}

/************ 3rd test************/
func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 {
			or ^= nums[left]
			left += 1
		}
		or |= x
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func isNice(slow, fast int, nums []int) bool {
	for i := slow; i < fast; i++ {
		for j := i + 1; j <= fast; j++ {
			if nums[i]&nums[j] != 0 {
				return false
			}
		}
	}
	return true
}
func Test_3rd(t *testing.T) {
	//t.Log(isNice(1, 3, []int{1, 3, 8, 48, 10}))
	nums := []int{1, 3, 8, 48, 10}
	t.Log(longestNiceSubarray(nums))
	nums = []int{3, 1, 5, 11, 13}
	t.Log(longestNiceSubarray(nums))
}

/************ 4th test************/
func mostBooked(n int, meetings [][]int) (ans int) {
	// 将会议按开始时间排序
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[i][j]
		}
		return meetings[i][0] < meetings[j][0]
	})
	// 每个会议室被分配了多少会议
	cnt := make([]int, n)
	// 每个会议室的最早可用时间
	tim := make([]int, n)
	for _, meet := range meetings {
		// best 表示当前不可用，但可用时间最早的会议室编号
		best := 0
		for i := 0; i < n; i++ {
			if tim[i] < meet[0] { // 当前会议室可用，直接分配
				cnt[i]++
				tim[i] = meet[1]
				goto OK
			} else if tim[i] < tim[best] { // 当前会议室不可用，更新 best
				best = i
			}
		}
		cnt[best]++
		tim[best] += meet[1] - meet[0]
	OK:
		continue
	}
	// 统计答案
	for i := 0; i < n; i++ {
		if cnt[i] > cnt[ans] {
			ans = i
		}
	}
	return
}

func Test_4th(t *testing.T) {
	n := 2
	meetings := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	t.Log(mostBooked(n, meetings))
	n = 3
	meetings = [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	t.Log(mostBooked(n, meetings))
	n = 4
	meetings = [][]int{{48, 49}, {22, 30}, {13, 31}, {31, 46}, {37, 46}, {32, 36}, {25, 36}, {49, 50}, {24, 34}, {6, 41}}
	t.Log(mostBooked(n, meetings))
}
