package weekly_contest

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

/************ 1st test************/
// func haveConflict(event1 []string, event2 []string) bool {
// 	s1, e1 := time2idx(event1[0]), time2idx(event1[1])
// 	s2, e2 := time2idx(event2[0]), time2idx(event2[1])
// 	// fmt.Println(s1, e1, s2, e2)
// 	return s1 <= e2 && s2 <= e1
// }

func haveConflict(event1 []string, event2 []string) bool {
	return event1[0] <= event2[1] && event2[0] <= event1[1]
}

func time2idx(time string) (ans int) {
	times := strings.Split(time, ":")
	hour, _ := strconv.Atoi(times[0])
	ans, _ = strconv.Atoi(times[1])
	for i := 0; i < hour; i++ {
		ans += 60
	}
	return
}

func Test_1st(t *testing.T) {
	event1 := []string{"01:15", "02:00"}
	event2 := []string{"02:00", "03:00"}
	t.Log(haveConflict(event1, event2)) //t

	event1 = []string{"01:00", "02:00"}
	event2 = []string{"01:20", "03:00"}
	t.Log(haveConflict(event1, event2)) //t

	event1 = []string{"10:00", "11:00"}
	event2 = []string{"14:00", "15:00"}
	t.Log(haveConflict(event1, event2)) //f
}

/************ 2nd test************/
func subarrayGCD(nums []int, k int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i; j < n; j++ {
			g = gcd(g, nums[j])
			if g == k {
				ans++
			} else if g < k {
				break
			}
		}
	}
	return
}
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func Test_2nd(t *testing.T) {
	nums, k := []int{9, 3, 1, 2, 6, 3}, 3
	t.Log(subarrayGCD(nums, k))
	nums, k = []int{4}, 7
	t.Log(subarrayGCD(nums, k))
	nums, k = []int{3, 12, 9, 6}, 3
	t.Log(subarrayGCD(nums, k))
}

/************ 3rd test************/

func minCost(nums []int, cost []int) (ans int64) {
	type pair struct {
		val, cos int
	}
	n := len(nums)
	tot, note := 0, 0
	pairs := make([]*pair, 0, n)
	for i := 0; i < n; i++ {
		cur := &pair{nums[i], cost[i]}
		pairs = append(pairs, cur)
		tot += cost[i]
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	chosen := -1
	for _, cur := range pairs {
		num, c := cur.val, cur.cos
		note += c
		if note >= tot/2 {
			chosen = num
			break
		}
	}
	for _, cur := range pairs {
		num, c := cur.val, cur.cos
		ans += int64(abs(num-chosen) * c)
	}
	return
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func Test_3rd(t *testing.T) {
	nums := []int{1, 3, 5, 2}
	cost := []int{2, 3, 1, 14}
	t.Log(minCost(nums, cost))
	nums = []int{2, 2, 2, 2, 2}
	cost = []int{4, 2, 8, 1, 3}
	t.Log(minCost(nums, cost))
}

/************ 4th test************/
func makeSimilar(nums []int, target []int) (ans int64) {
	sort.Ints(nums)
	sort.Ints(target)
	js := [2]int{} //奇偶组指针
	for _, x := range nums {
		p := x % 2 //x奇偶性
		for target[js[p]]%2 != p {
			js[p]++
		}
		ans += int64(abs(x - target[js[p]])) //当前匹配组差值
		js[p]++
	}
	return ans / 4 //每组操作+2,-2能获得4差值
}
func Test_4th(t *testing.T) {
	nums := []int{8, 12, 6}
	target := []int{2, 14, 10}
	t.Log(makeSimilar(nums, target))
	nums = []int{1, 2, 5}
	target = []int{4, 1, 3}
	t.Log(makeSimilar(nums, target))
}
