package range_frequency_queries_mid

import (
	"sort"
	"testing"
)

type RangeFreqQuery map[int][]int

func Constructor(arr []int) RangeFreqQuery {
	pos := make(map[int][]int)
	for i, x := range arr {
		pos[x] = append(pos[x], i)
	}
	return pos
}

func (p *RangeFreqQuery) Query(left int, right int, value int) int {
	arr := (*p)[value]

	return sort.SearchInts(arr, right+1) - sort.SearchInts(arr, left)
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */

func Test_range_frequency_queries(t *testing.T) {
	tests := []struct {
		cmds []string
		cmdi [][]int
	}{
		{
			cmds: []string{"RangeFreqQuery", "query", "query"},
			cmdi: [][]int{{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56},
				{1, 2, 4}, {0, 11, 33}},
		},
	}
	for _, tt := range tests {
		var sol RangeFreqQuery
		for i, cmd := range tt.cmds {
			switch cmd {
			case "RangeFreqQuery":
				sol = Constructor(tt.cmdi[i])
			case "query":
				if sol != nil {
					t.Log(sol.Query(tt.cmdi[i][0], tt.cmdi[i][1], tt.cmdi[i][2]))
				}
			}
		}
	}
}
