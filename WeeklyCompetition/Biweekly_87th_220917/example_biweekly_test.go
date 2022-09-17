package biweeklycompetition

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

/************ 1st test************/
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	s1, s2 := day2idx(arriveAlice), day2idx(arriveBob)
	e1, e2 := day2idx(leaveAlice), day2idx(leaveBob)
	if e1 < s2 || e2 < s1 {
		return 0
	}
	return min(e1, e2) - maxint(s1, s2) + 1
}
func maxint(a, b int) int {
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

func day2idx(date string) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	md := strings.Split(date, "-")
	month, _ := strconv.Atoi(md[0])
	day, _ := strconv.Atoi(md[1])
	idx := 0
	for i := 0; i < month-1; i++ {
		idx += days[i]
	}
	idx += day
	return idx
}
func Test_1st(t *testing.T) {
	arriveAlice := "08-15"
	leaveAlice := "08-18"
	arriveBob := "08-16"
	leaveBob := "08-19"
	t.Log(countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob))
	arriveAlice = "10-01"
	leaveAlice = "10-31"
	arriveBob = "11-01"
	leaveBob = "12-31"
	t.Log(countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob))
	arriveAlice = "9-01"
	leaveAlice = "10-19"
	arriveBob = "06-19"
	leaveBob = "10-20"
	t.Log(countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob))
}

/************ 2nd test************/
func matchPlayersAndTrainers(players []int, trainers []int) (ans int) {
	sort.Ints(players)
	sort.Ints(trainers)
	i, j := 0, 0
	ans = 0
	for ; j < len(trainers) && i < len(players); j++ {
		if players[i] <= trainers[j] {
			ans++
			i++
		}
	}
	return ans
}
func Test_2nd(t *testing.T) {
	players := []int{4, 7, 9}
	trainers := []int{8, 2, 5, 8}
	t.Log(matchPlayersAndTrainers(players, trainers))
	players = []int{1, 1, 1}
	trainers = []int{10}
	t.Log(matchPlayersAndTrainers(players, trainers))
	players = []int{4, 2}
	trainers = []int{4, 4, 3}
	t.Log(matchPlayersAndTrainers(players, trainers))
}

/************ 3rd test************/
func smallestSubarrays_enum(nums []int) []int {
	const maxp = 30
	n := len(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, maxp+1)
	}
	for i := maxp; i >= 0; i-- {
		if nums[n-1]>>i&1 == 1 {
			f[n-1][i] = n - 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		for j := maxp; j >= 0; j-- {
			f[i][j] = f[i+1][j]
			if nums[i]>>j&1 == 1 {
				f[i][j] = i
			}
		}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		mx := i
		for j := 0; j <= maxp; j++ {
			mx = maxint(mx, f[i][j])
			ans[i] = mx - i + 1
		}
	}
	return ans
}
func Test_3rd(t *testing.T) {
	nums := []int{1, 0, 2, 1, 3}
	t.Log(smallestSubarrays_enum(nums))
	nums = []int{1, 2}
	t.Log(smallestSubarrays_enum(nums))
	nums = []int{1, 15, 2, 18, 4, 6, 16}
	t.Log(smallestSubarrays_enum(nums))
	nums = []int{8, 10, 8}
	t.Log(smallestSubarrays_enum(nums))
}

/************ 4th test************/
func minimumMoney(transactions [][]int) int64 {
	var ans, neg int64 = 0, 0
	var det int64
	for _, vec := range transactions {
		det = int64(vec[1] - vec[0])
		if det < 0 {
			neg += det
		}
	}
	for _, vec := range transactions {
		det = int64(vec[1] - vec[0])
		if det < 0 {
			ans = max(ans, int64(vec[0])-(neg-det))
		} else {
			ans = max(ans, int64(vec[0])-neg)
		}
	}
	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Test_4th(t *testing.T) {
	transactions := [][]int{{2, 1}, {5, 0}, {4, 2}}
	t.Log(minimumMoney(transactions))
	transactions = [][]int{{3, 0}, {0, 3}}
	t.Log(minimumMoney(transactions))
}
