package pancakesorting

import "testing"

var res []int

func pancakeSort(arr []int) []int {
	res = []int{}
	sort(arr, len(arr))
	return res
}
func sort(cakes []int, n int) {
	if n == 1 { // base case
		return
	}
	maxCake, maxCakeIdx := 0, 0
	for i := 0; i < n; i++ {
		if cakes[i] > maxCake {
			maxCakeIdx = i
			maxCake = cakes[i]
		}
	}
	reverse(cakes, 0, maxCakeIdx)
	res = append(res, maxCakeIdx+1)

	reverse(cakes, 0, n-1)
	res = append(res, n)

	sort(cakes, n-1)
}
func reverse(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
func Test_pancake_sorting(t *testing.T) {
	arr := []int{3, 2, 4, 1}
	t.Log(pancakeSort(arr))
	arr = []int{1, 2, 3}
	t.Log(pancakeSort(arr))
}
