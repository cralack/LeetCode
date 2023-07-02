package maximumscorefromremovingstones

import (
	"sort"
	"testing"
)

func maximumScore(a int, b int, c int) (ans int) {
	arr := []int{a, b, c}
	sort.Ints(arr)
	for arr[1] > 0 {
		ans++
		arr[1]--
		arr[2]--
		sort.Ints(arr)
	}
	return
}
func maximumScore_math(a int, b int, c int) (ans int) {
	arr := []int{a, b, c}
	sort.Ints(arr)
	// triangle?
	if arr[0]+arr[1] <= arr[2] {
		return arr[0] + arr[1]
	}
	return (a + b + c) >> 1
}
func Test_maximum_score_from_removing_stones(t *testing.T) {
	a, b, c := 2, 4, 6
	t.Log(maximumScore(a, b, c))
	a, b, c = 4, 4, 6
	t.Log(maximumScore(a, b, c))
	a, b, c = 1, 8, 8
	t.Log(maximumScore(a, b, c))
}
