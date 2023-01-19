package findingtheusersactiveminutes

import (
	"testing"
)

func findingUsersActiveMinutes(logs [][]int, k int) (ans []int) {
	ans = make([]int, k+1)
	cnt := make(map[int]map[int]struct{}, 0)
	for _, log := range logs {
		id, time := log[0], log[1]
		if cnt[id] == nil {
			cnt[id] = map[int]struct{}{}
		}
		cnt[id][time] = struct{}{}
	}
	for _, c := range cnt {
		ans[len(c)]++
	}
	return ans[1:]
}

func Test_finding_the_users_active_minutes(t *testing.T) {
	logs, k := [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5
	t.Log(findingUsersActiveMinutes(logs, k))
	logs, k = [][]int{{1, 1}, {2, 2}, {2, 3}}, 4
	t.Log(findingUsersActiveMinutes(logs, k))
	logs, k = [][]int{{305589003, 4136}, {305589004, 4139}, {305589004, 4141},
		{305589004, 4137}, {305589001, 4139}, {305589001, 4139}}, 6
	t.Log(findingUsersActiveMinutes(logs, k))
}
