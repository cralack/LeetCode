package heightchecker

import (
	"sort"
	"testing"
)

func heightChecker(heights []int) int {
	arr := make([]int, len(heights))
	copy(arr, heights)
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	cnt := 0
	for i := range heights {
		if arr[i] != heights[i] {
			cnt++
		}
	}
	return cnt
}
func heightChecker_cntsort(heights []int) (ans int) {
	cnt := make([]int, 101)
	for _, v := range heights {
		cnt[v]++
	}
	idx := 0
	for i, c := range cnt {
		for ; c > 0; c-- {
			if heights[idx] != i {
				ans++
			}
			idx++
		}
	}
	return
}
func Test_height_checker(t *testing.T) {
	heights := []int{1, 1, 4, 2, 1, 3}
	t.Log(heightChecker_cntsort(heights))
	heights = []int{5, 1, 2, 3, 4}
	t.Log(heightChecker(heights))
	heights = []int{1, 2, 3, 4, 5}
	t.Log(heightChecker(heights))
}
