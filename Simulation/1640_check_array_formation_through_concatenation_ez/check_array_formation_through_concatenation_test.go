package checkarrayformationthroughconcatenationez

import "testing"

func canFormArray(arr []int, pieces [][]int) bool {
	idx := make(map[int]int, 0)
	for i, n := range arr {
		idx[n] = i
	}
	valid, n := 0, len(arr)
	for _, piece := range pieces {
		for i := 0; i < len(piece); i++ {
			//piece第一个元素有对应的idx
			if _, has := idx[piece[i]]; has && i == 0 ||
				//piece后续的元素能idx+1
				i > 0 && (idx[piece[i]] == idx[piece[i-1]]+1) {
				valid++
			} else {
				return false
			}
		}
	}
	return valid == n
}
func Test_check_array_formation_through_concatenation(t *testing.T) {
	arr := []int{15, 88}
	pieces := [][]int{{88}, {15}}
	t.Log(canFormArray(arr, pieces))
	arr = []int{49, 18, 16}
	pieces = [][]int{{16, 18, 49}}
	t.Log(canFormArray(arr, pieces))
	arr = []int{91, 4, 64, 78}
	pieces = [][]int{{78}, {4, 64}, {91}}
	t.Log(canFormArray(arr, pieces))
	arr = []int{12}
	pieces = [][]int{{1}}
	t.Log(canFormArray(arr, pieces))
}
