package buildanarraywithstackoperations

import "testing"

func buildArray(target []int, n int) (res []string) {
	size := len(target)
	res = make([]string, 0, size)
	p := 0
	for i := 1; i <= n && p < size; i++ {
		res = append(res, "Push")
		if target[p] > i {
			res = append(res, "Pop")
			p--
		}
		p++
	}
	return
}
func Test_build_an_array_with_stack_operations(t *testing.T) {
	target, n := []int{1, 3}, 3
	t.Log(buildArray(target, n))
	target, n = []int{1, 2, 3}, 3
	t.Log(buildArray(target, n))
	target, n = []int{1, 2}, 4
	t.Log(buildArray(target, n))
}
