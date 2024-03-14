package meetingroomsii

import (
	"sort"
	"testing"
)

func minMeetingRooms(meetings [][]int) int {
	n := len(meetings)
	if n == 0 {
		return 0
	}
	begin, end := make([]int, n), make([]int, n)
	for i, meeting := range meetings {
		begin[i] = meeting[0]
		end[i] = meeting[1]
	}
	sort.Ints(begin)
	sort.Ints(end)
	cnt, res := 0, 0
	i, j := 0, 0
	for i < n && j < n {
		if begin[i] < end[j] { // 扫描到一个红点
			cnt++
			i++
		} else { // 扫描到一个绿点
			cnt--
			j++
		}
		if res < cnt { // 记录扫描过程中的最大值
			res = cnt
		}
	}
	return res
}
func Test_meeting_rooms_ii(t *testing.T) {
	meetins := [][]int{{0, 30}, {5, 10}, {15, 20}}
	t.Log(minMeetingRooms(meetins))
}
