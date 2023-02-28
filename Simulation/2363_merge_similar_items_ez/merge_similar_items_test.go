package mergesimilaritems

import (
	"sort"
	"testing"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) (ans [][]int) {
	kvMap := make(map[int]int)
	for _, it := range items1 {
		kvMap[it[0]] += it[1]
	}
	for _, it := range items2 {
		kvMap[it[0]] += it[1]
	}
	for k, v := range kvMap {
		ans = append(ans, []int{k, v})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return
}
func Test_merge_similar_items(t *testing.T) {
	items1 := [][]int{{1, 1}, {4, 5}, {3, 8}}
	items2 := [][]int{{3, 1}, {1, 5}}
	t.Log(mergeSimilarItems(items1, items2))
	items1 = [][]int{{1, 1}, {3, 2}, {2, 3}}
	items2 = [][]int{{2, 1}, {3, 2}, {1, 3}}
	t.Log(mergeSimilarItems(items1, items2))
	items1 = [][]int{{1, 3}, {2, 2}}
	items2 = [][]int{{7, 1}, {2, 2}, {1, 4}}
	t.Log(mergeSimilarItems(items1, items2))
}
