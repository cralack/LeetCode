package bestteamwithnoconflicts

import (
	"sort"
	"testing"
)

func bestTeamScore(scores []int, ages []int) (ans int) {
	type person struct {
		score, age int
	}
	n := len(scores)
	persons := make([]person, n)
	for i, v := range scores {
		persons[i] = person{v, ages[i]}
	}
	sort.Slice(persons, func(i, j int) bool {
		a, b := persons[i], persons[j]
		return a.score < b.score || a.score == b.score && a.age < b.age
	})

	dp := make([]int, n)
	for i, p := range persons {
		for j, q := range persons[:i] {
			if q.age <= p.age {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += p.score
		ans = max(ans, dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_best_team_with_no_conflicts(t *testing.T) {
	tests := []struct {
		scores, ages []int
	}{
		{scores: []int{1, 3, 5, 10, 15}, ages: []int{1, 2, 3, 4, 5}},
		{scores: []int{4, 5, 6, 5}, ages: []int{2, 1, 2, 1}},
		{scores: []int{1, 2, 3, 5}, ages: []int{8, 9, 10, 1}},
	}

	for _, tt := range tests {
		t.Log(bestTeamScore(tt.scores, tt.ages))
	}
}
