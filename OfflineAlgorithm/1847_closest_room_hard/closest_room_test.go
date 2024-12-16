package closest_room_hard

import (
	"math"
	"slices"
	"testing"

	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

func closestRoom(rooms [][]int, queries [][]int) []int {
	slices.SortFunc(rooms, func(a, b []int) int {
		return b[1] - a[1]
	})

	queryIds := make([]int, len(queries))
	for i := range queryIds {
		queryIds[i] = i
	}
	// sort.Slice不稳定
	slices.SortFunc(queryIds, func(i, j int) int {
		return queries[j][1] - queries[i][1]
	})

	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}

	roomIds := redblacktree.New[int, struct{}]()
	j := 0
	for _, i := range queryIds {
		preferredId, minSize := queries[i][0], queries[i][1]
		for j < len(rooms) && rooms[j][1] >= minSize {
			roomIds.Put(rooms[j][0], struct{}{})
			j++
		}

		diff := math.MaxInt
		// 左边的差
		if node, ok := roomIds.Floor(preferredId); ok {
			diff = preferredId - node.Key
			ans[i] = node.Key
		}
		// 右边的差
		if node, ok := roomIds.Ceiling(preferredId); ok && node.Key-preferredId < diff {
			ans[i] = node.Key
		}
	}
	return ans
}

func Test_closest_room(t *testing.T) {
	tests := []struct {
		rooms   [][]int
		queries [][]int
		wanna   []int
	}{
		{rooms: [][]int{{2, 2}, {1, 2}, {3, 2}}, queries: [][]int{{3, 1}, {3, 3}, {5, 2}}, wanna: []int{3, -1, 3}},
		{rooms: [][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, queries: [][]int{{2, 3}, {2, 4}, {2, 5}}, wanna: []int{2, 1, 3}},
		{rooms: [][]int{{23, 22}, {6, 20}, {15, 6}, {22, 19}, {2, 10}, {21, 4}, {10, 18}, {16, 1}, {12, 7}, {5, 22}},
			queries: [][]int{{12, 5}, {15, 15}, {21, 6}, {15, 1}, {23, 4}, {15, 11}, {1, 24}, {3, 19}, {25, 8}, {18, 6}},
			wanna:   []int{12, 10, 22, 15, 23, 10, -1, 5, 23, 15}},
	}
	for _, tt := range tests {
		t.Log(closestRoom(tt.rooms, tt.queries))
		t.Log(tt.wanna)
		t.Log()
	}
}
