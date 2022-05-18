package minimummovestoequalarrayelementsii

import (
	"sort"
	"testing"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	lo, hi := 0, len(nums)-1
	cnt := 0
	for lo < hi {
		cnt += nums[hi] - nums[lo]
		lo++
		hi--
	}
	return cnt
}
func Test_minimum_moves_to_equal_array_elements_ii(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(minMoves2(nums))
	nums = []int{1, 10, 2, 9}
	t.Log(minMoves2(nums))
}
