package reduce_array_size_to_the_half_mid

import (
	"sort"
	"testing"
)

func minSetSize(arr []int) int {
	n := len(arr)
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	cnts := make([]int, 0)
	for _, c := range freq {
		cnts = append(cnts, c)
	}
	sort.Slice(cnts, func(i, j int) bool {
		return cnts[i] > cnts[j]
	})
	sum := 0
	for i, c := range cnts {
		sum += c
		if sum*2 >= n {
			return i + 1
		}
	}
	return -1
}

func Test_reduce_array_size_to_the_half(t *testing.T) {
	tests := []struct {
		arr []int
	}{
		{arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}},
		{arr: []int{7, 7, 7, 7, 7, 7}},
	}
	for _, tt := range tests {
		t.Log(minSetSize(tt.arr))
	}
}
