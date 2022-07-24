package ranktransformofanarrayez

import (
	"sort"
	"testing"
)

func arrayRankTransform(arr []int) []int {
	index := make([]int, len(arr))
	for i := range arr {
		index[i] = arr[i]
	}
	sort.Slice(index, func(i, j int) bool { return index[i] < index[j] })
	idx := make(map[int]int, 0)
	for _, val := range index {
		if _, has := idx[val]; !has {
			idx[val] = len(idx) + 1
		}

	}
	for i, val := range arr {
		arr[i] = idx[val]
	}
	return arr
}

func Test_rank_transform_of_an_array(t *testing.T) {
	arr := []int{40, 10, 20, 30}
	t.Log(arrayRankTransform(arr))
	arr = []int{100, 100, 100}
	t.Log(arrayRankTransform(arr))
	arr = []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	t.Log(arrayRankTransform(arr))
}
