package arrayofdoubledpairs

import (
	"sort"
	"testing"
)

func canReorderDoubled(arr []int) bool {
	n := len(arr)
	cnt := make(map[int]int, n)
	for _, v := range arr {
		cnt[v]++
	}
	vals := make([]int, 0, len(cnt))
	for i := range cnt {
		vals = append(vals, i)
	}
	sort.Slice(vals, func(i, j int) bool {
		return abs(vals[i]) < abs(vals[j])
	})

	for _, val := range vals {
		if cnt[2*val] < cnt[val] {
			return false
		}
		cnt[2*val] -= cnt[val]
	}
	return true
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func Test_array_of_doubled_pairs(t *testing.T) {
	arr := []int{3, 1, 3, 6}
	t.Log(canReorderDoubled(arr))
	arr = []int{2, 1, 2, 6}
	t.Log(canReorderDoubled(arr))
	arr = []int{4, -2, 2, -4}
	t.Log(canReorderDoubled(arr))
	arr = []int{2, 4, 0, 0, 8, 1}
	t.Log(canReorderDoubled(arr))
	arr = []int{2, 1, 2, 1, 1, 1, 2, 2}
	t.Log(canReorderDoubled(arr))
}
