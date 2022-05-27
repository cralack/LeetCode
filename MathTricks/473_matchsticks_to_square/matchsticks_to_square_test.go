package matchstickstosquare

import (
	"sort"
	"testing"
)

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, n := range matchsticks {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	edges := make([]int, 4)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return true
		}
		if matchsticks[idx] > sum/4 {
			return false
		}
		for i := range edges {
			edges[i] += matchsticks[idx]
			if edges[i] <= sum/4 && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}
	return dfs(0)
}
func Test_matchsticks_to_square(t *testing.T) {
	matchsticks := []int{1, 1, 2, 2, 2}
	t.Log(makesquare(matchsticks))
	matchsticks = []int{3, 3, 3, 3, 4}
	t.Log(makesquare(matchsticks))
	matchsticks = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 102}
	t.Log(makesquare(matchsticks))
}
