package customsortstring

import (
	"sort"
	"testing"
)

func customSortString(order string, s string) string {
	idx := make(map[byte]int, 0)
	for i, char := range order {
		idx[byte(char)] = i
	}
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return idx[str[i]] < idx[str[j]]
	})
	return string(str)
}

func Test_custom_sort_string(t *testing.T) {
	order := "cba"
	s := "abcd"
	t.Log(customSortString(order, s))
	order = "cbafg"
	s = "abcd"
	t.Log(customSortString(order, s))
}
