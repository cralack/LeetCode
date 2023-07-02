package weekly_contest

import (
	"sort"
	"testing"
)

/************ 1st test************/
func sortPeople(names []string, heights []int) []string {
	hei2na := make(map[int]int, 0)
	for i, hei := range heights {
		hei2na[hei] = i
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	ans := make([]string, len(names))
	for i, hei := range heights {
		ans[i] = names[hei2na[hei]]
	}
	return ans
}
func Test_1st(t *testing.T) {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	t.Log(sortPeople(names, heights))
	names = []string{"Alice", "Bob", "Bob"}
	heights = []int{155, 185, 150}
	t.Log(sortPeople(names, heights))
}

/************ 2nd test************/
func longestSubarray(nums []int) (cnt int) {
	max := 0
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	tmp := 0
	for _, n := range nums {
		if n == max {
			tmp++
			if cnt < tmp {
				cnt = tmp
			}
		} else {
			tmp = 0
		}
	}
	return
}
func Test_2nd(t *testing.T) {
	nums := []int{1, 2, 3, 3, 2, 2}
	t.Log(longestSubarray(nums))
	nums = []int{1, 2, 3, 4}
	t.Log(longestSubarray(nums))
}

/************ 3rd test************/
func goodIndices(nums []int, k int) (ans []int) {
	n := len(nums)
	down := make([]int, n)
	up := make([]int, n)
	cnt := 0

	for i := 1; i < n; i++ {
		if nums[i-1] >= nums[i] {
			cnt++
		} else {
			cnt = 0
		}
		down[i] = cnt
	}

	cnt = 0
	for i := n - 2; i >= 0; i-- {
		if nums[i+1] >= nums[i] {
			cnt++
		} else {
			cnt = 0
		}
		up[i] = cnt
	}

	for i := k; i < n-k; i++ {
		if down[i-1] >= k-1 && up[i+1] >= k-1 {
			ans = append(ans, i)
		}
	}

	return
}
func Test_3rd(t *testing.T) {
	nums := []int{2, 1, 1, 1, 3, 4, 1}
	k := 2
	t.Log(goodIndices(nums, k))
	nums = []int{2, 1, 1, 2}
	k = 2
	t.Log(goodIndices(nums, k))
	nums = []int{878724, 201541, 179099, 98437, 35765, 327555, 475851, 598885, 849470, 943442}
	k = 4
	t.Log(goodIndices(nums, k))
}

/************ 4th test************/
func numberOfGoodPaths(vals []int, edges [][]int) (ans int) {
	// todo

	return
}
func Test_4th(t *testing.T) {
	vals := []int{1, 3, 2, 1, 3}
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}
	t.Log(numberOfGoodPaths(vals, edges))
	vals = []int{1, 1, 2, 2, 3}
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}
	t.Log(numberOfGoodPaths(vals, edges))
	vals = []int{1}
	edges = [][]int{}
	t.Log(numberOfGoodPaths(vals, edges))
}
