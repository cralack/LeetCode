package find_maximum_value_in_a_constrained_sequence_mid

import (
	"slices"
	"testing"
)

func findMaxVal(n int, restrictions [][]int, diff []int) int {
	maxVal := make([]int, n)
	for i := range maxVal {
		maxVal[i] = 1 << 31
	}
	for _, r := range restrictions {
		maxVal[r[0]] = r[1]
	}
	ans := make([]int, n)
	for i, d := range diff {
		ans[i+1] = min(ans[i]+d, maxVal[i+1])
	}
	for i := n - 2; i > 0; i-- {
		ans[i] = min(ans[i], ans[i+1]+diff[i])
	}
	return slices.Max(ans)
}

func Test_find_maximum_value_in_a_constrained_sequence(t *testing.T) {
	tests := []struct {
		n            int
		restrictions [][]int
		diff         []int
	}{
		{n: 10, restrictions: [][]int{{3, 1}, {8, 1}}, diff: []int{2, 2, 3, 1, 4, 5, 1, 1, 2}},
		{n: 8, restrictions: [][]int{{3, 2}}, diff: []int{3, 5, 2, 4, 2, 3, 1}},
	}
	for _, test := range tests {
		t.Log(findMaxVal(test.n, test.restrictions, test.diff))
	}
}
