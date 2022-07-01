package minimumabsolutedifferenceeasy

import (
	"sort"
	"testing"
)

func minimumAbsDifference(arr []int) (ans [][]int) {
	sort.Ints(arr)
	minDif := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if minDif > diff {
			minDif = diff
			ans = [][]int{{arr[i], arr[i+1]}}
		} else if minDif == diff {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return
}
func Test_minimum_absolute_difference(t *testing.T) {
	arr := []int{4, 2, 1, 3}
	t.Log(minimumAbsDifference(arr))
	arr = []int{1, 3, 6, 10, 15}
	t.Log(minimumAbsDifference(arr))
	arr = []int{3, 8, -10, 23, 19, -4, -14, 27}
	t.Log(minimumAbsDifference(arr))
}
