package videostitching

import (
	"sort"
	"testing"
)

func videoStitching(clips [][]int, time int) int {
	if time == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] != clips[j][0] {
			return clips[i][0] < clips[j][0]
		}
		return clips[i][1] > clips[j][1]
	})
	cnt, i := 0, 0
	curEnd, nextEnd := 0, 0
	for i < len(clips) && clips[i][0] <= curEnd {
		for i < len(clips) && clips[i][0] <= curEnd {
			nextEnd = max(clips[i][1], nextEnd)
			i++
		}
		curEnd = nextEnd
		cnt++
		if curEnd >= time {
			return cnt
		}
	}
	return -1
}
func Test_video_stitching(t *testing.T) {
	clips, time := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10
	t.Log(videoStitching(clips, time))

	clips, time = [][]int{{0, 1}, {1, 2}}, 5
	t.Log(videoStitching(clips, time))

	clips, time = [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3},
		{6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5},
		{5, 7}, {6, 9}}, 9
	t.Log(videoStitching(clips, time))
}
