package formarraybyconcatenatingsubarraysofanotherarray

import "testing"

func canChoose(groups [][]int, nums []int) bool {
	check := func(a, b []int, j int) bool {
		m, n := len(a), len(b)
		i := 0
		for ; i < m && j < n; i, j = i+1, j+1 {
			if a[i] != b[j] {
				return false
			}
		}
		return i == m
	}
	n, m := len(groups), len(nums)
	i := 0
	for j := 0; i < n && j < m; {
		if check(groups[i], nums, j) {
			j += len(groups[i])
			i++
		} else {
			j++
		}
	}
	return i == n
}
func Test_form_array_by_concatenating_subarrays(t *testing.T) {
	groups := [][]int{{1, -1, -1}, {3, -2, 0}}
	nums := []int{1, -1, 0, 1, -1, -1, 3, -2, 0}
	t.Log(canChoose(groups, nums))
	groups = [][]int{{10, -2}, {1, 2, 3, 4}}
	nums = []int{1, 2, 3, 4, 10, -2}
	t.Log(canChoose(groups, nums))
	groups = [][]int{{1, 2, 3}, {3, 4}}
	nums = []int{7, 7, 1, 2, 3, 4, 7, 7}
	t.Log(canChoose(groups, nums))
}
