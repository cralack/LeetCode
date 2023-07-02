package largestvaluesfromlabels

import (
	"sort"
	"testing"
)

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) (ans int) {
	// zip
	type pair struct {
		val, lab int
	}
	pairs := make([]pair, len(values))
	for i := range values {
		pairs[i] = pair{values[i], labels[i]}
	}
	// sort↘
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val > pairs[j].val
	})
	// greed
	dic := make(map[int]int)
	for i := 0; i < len(values); i++ {
		if c := dic[pairs[i].lab]; c < useLimit && numWanted > 0 {
			ans += pairs[i].val
			dic[pairs[i].lab]++
			numWanted--
		}
	}
	return
}

func Test_largest_values_from_labels(t *testing.T) {
	tests := []struct {
		values    []int
		labels    []int
		numWanted int
		useLimit  int
	}{
		{values: []int{5, 4, 3, 2, 1}, labels: []int{1, 1, 2, 2, 3}, numWanted: 3, useLimit: 1},
		{values: []int{5, 4, 3, 2, 1}, labels: []int{1, 3, 3, 3, 2}, numWanted: 3, useLimit: 2},
		{values: []int{9, 8, 8, 7, 6}, labels: []int{0, 0, 0, 1, 1}, numWanted: 3, useLimit: 1},
	}

	for _, tt := range tests {
		t.Log(largestValsFromLabels(tt.values, tt.labels, tt.numWanted, tt.useLimit))
	}
}
