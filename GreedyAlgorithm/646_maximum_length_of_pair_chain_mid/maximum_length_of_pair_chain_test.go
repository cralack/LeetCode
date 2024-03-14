package maximumlengthofpairchainmid

import (
	"sort"
	"testing"
)

func findLongestChain(pairs [][]int) (ans int) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	n := len(pairs)
	ans = 1
	right := pairs[0][1]
	for i := 1; i < n; i++ {
		curLeft, curRight := pairs[i][0], pairs[i][1]
		if right < curLeft {
			ans++
			right = curRight
		}
	}
	return
}
func Test_maximum_length_of_pair_chain(t *testing.T) {
	pairs := [][]int{{1, 2}, {2, 3}, {3, 4}}
	t.Log(findLongestChain(pairs))
	pairs = [][]int{{-6, 9}, {1, 6}, {8, 10}, {-1, 4}, {-6, -2}, {-9, 8}, {-5, 3}, {0, 3}}
	t.Log(findLongestChain(pairs))
}
