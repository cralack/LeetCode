package maketwoarraysequalbyreversingsubarraysez

import "testing"

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	cnt := make(map[int]int, 0)
	for i, x := range target {
		cnt[x]++
		cnt[arr[i]]--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}
func Test_make_two_arrays_equal_by_reversing_sub_arrays(t *testing.T) {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	t.Log(canBeEqual(target, arr))
	target = []int{7}
	arr = []int{7}
	t.Log(canBeEqual(target, arr))
	target = []int{3, 7, 9}
	arr = []int{7, 3, 11}
	t.Log(canBeEqual(target, arr))
}
