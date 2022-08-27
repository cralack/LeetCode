package weeklycompetition

import (
	"sort"
	"testing"
)

/************ 1st test************/
func answerQueries(nums []int, queries []int) (ans []int) {
	n := len(nums)
	sort.Ints(nums)
	ans = make([]int, 0, len(queries))
	for _, q := range queries {
		sum := 0
		for i := 0; i < n; i++ {
			sum += nums[i]
			if sum > q {
				ans = append(ans, i)
				break
			}
		}
		if sum <= q {
			ans = append(ans, n)
		}
	}
	return
}
func Test_1st(t *testing.T) {
	nums := []int{4, 5, 2, 1}
	queries := []int{3, 10, 21}
	t.Log(answerQueries(nums, queries))
	nums = []int{2, 3, 4, 5}
	queries = []int{1}
	t.Log(answerQueries(nums, queries))
	nums = []int{469781, 45635, 628818, 324948, 343772, 713803, 452081}
	queries = []int{816646, 929491}
	t.Log(answerQueries(nums, queries))
}

/************ 2nd test************/
func removeStars(s string) string {
	ans := []byte{}
	for _, c := range s {
		if c == '*' {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, byte(c))
		}
	}

	return string(ans)
}
func Test_2nd(t *testing.T) {
	s := "leet**cod*e"
	t.Log(removeStars(s))
	s = "erase*****"
	t.Log(removeStars(s))
}

/************ 3rd test************/
func garbageCollection(garbage []string, travel []int) (ans int) {
	rightBound := make([]int, 3) //M,P,G  last idx
	for i, gar := range garbage {
		ans += len(gar)
		for _, c := range gar {
			switch c {
			case 'M':
				rightBound[0] = i
			case 'P':
				rightBound[1] = i
			case 'G':
				rightBound[2] = i
			}
		}
	}
	for i := range travel {
		for j := 0; j < 3; j++ {
			if rightBound[j] > 0 && rightBound[j] > i {
				ans += travel[i]
			}
		}
	}
	return
}
func Test_3rd(t *testing.T) {
	garbage := []string{"G", "P", "GP", "GG"}
	travel := []int{2, 4, 3}
	t.Log(garbageCollection(garbage, travel))
	garbage = []string{"MMM", "PGM", "GP"}
	travel = []int{3, 10}
	t.Log(garbageCollection(garbage, travel))
}

/************ 4th test************/
