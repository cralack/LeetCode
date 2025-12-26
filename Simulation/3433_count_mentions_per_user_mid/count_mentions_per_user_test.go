package count_mentions_per_user_mid

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	// 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
	slices.SortFunc(events, func(a, b []string) int {
		ta, _ := strconv.Atoi(a[1])
		tb, _ := strconv.Atoi(b[1])
		return cmp.Or(ta-tb, int(b[0][0])-int(a[0][0]))
	})

	ans := make([]int, numberOfUsers)
	onlineT := make([]int, numberOfUsers)
	for _, e := range events {
		curT, _ := strconv.Atoi(e[1]) // 当前时间
		mention := e[2]
		if e[0][0] == 'O' { // 离线
			i, _ := strconv.Atoi(mention)
			onlineT[i] = curT + 60 // 下次在线时间
		} else if mention[0] == 'A' { // @所有人
			for i := range ans {
				ans[i]++
			}
		} else if mention[0] == 'H' { // @所有在线用户
			for i, t := range onlineT {
				if t <= curT { // 在线
					ans[i]++
				}
			}
		} else { // @id
			for _, s := range strings.Split(mention, " ") {
				i, _ := strconv.Atoi(s[2:])
				ans[i]++
			}
		}
	}
	return ans
}

func Test_count_mentions_per_user(t *testing.T) {
	tests := []struct {
		numberOfUsers int
		events        [][]string
	}{
		{numberOfUsers: 2, events: [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "71", "HERE"}}},
		{numberOfUsers: 2, events: [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "12", "ALL"}}},
		{numberOfUsers: 2, events: [][]string{{"OFFLINE", "10", "0"}, {"MESSAGE", "12", "HERE"}}},
	}

	for _, test := range tests {
		t.Log(countMentions(test.numberOfUsers, test.events))
	}
}
