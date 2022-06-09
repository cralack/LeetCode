package kdiffpairsinanarray

import (
	"sort"
	"testing"
)

func findPairs(nums []int, k int) int {
	vis := make(map[int]bool, 0)
	set := make(map[int]bool, 0)
	for _, n := range nums {
		if _, has := vis[n+k]; has {
			set[n] = true
		}
		if _, has := vis[n-k]; has {
			set[n-k] = true
		}
		vis[n] = true
	}
	return len(set)
}
func findPairs_sort_slw(nums []int, k int) int {
	sort.Ints(nums)
	cnt, left, right := 0, 0, 1
	for right < len(nums) {
		for right < len(nums)-1 && nums[right] == nums[right+1] {
			right++
		}
		dif := nums[right] - nums[left]
		if dif == k && left != right {
			cnt++
			left++
			right++
		} else if dif > k {
			left++
		} else {
			right++
		}
	}
	return cnt
}
func Test_k_diff_pairs_in_an_array(t *testing.T) {
	nums, k := []int{3, 1, 4, 1, 5}, 2
	t.Log(findPairs_sort_slw(nums, k))
	nums, k = []int{1, 2, 3, 4, 5}, 1
	t.Log(findPairs(nums, k))
	nums, k = []int{1, 3, 1, 5, 4}, 0
	t.Log(findPairs(nums, k))
}
