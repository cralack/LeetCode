package minimum_number_of_people_to_teach_mid

import (
	"slices"
	"testing"
)

func minimumTeachings(n int, languages [][]int, friendships [][]int) (ans int) {
	m := len(languages)
	dic := make(map[int][]bool, m)
	for i, l := range languages {
		dic[i] = make([]bool, n)
		for _, x := range l {
			dic[i][x-1] = true
		}
	}
	// check if u and v can speak
	check := func(u, v int) bool {
		for i := 0; i < n; i++ {
			if dic[u-1][i] && dic[v-1][i] == true {
				return true
			}
		}
		return false
	}
	// how many user need language education
	s := map[int]bool{}
	for _, e := range friendships {
		u, v := e[0], e[1]
		if !check(u, v) {
			s[u], s[v] = true, true
		}
	}

	// which language needed users most
	cnt := make([]int, n)
	for u := range s {
		for _, l := range languages[u-1] {
			cnt[l-1]++
		}
	}
	return len(s) - slices.Max(cnt)
}

func Test_minimum_number_of_people_to_teach(t *testing.T) {
	tests := []struct {
		n           int
		languages   [][]int
		friendships [][]int
	}{
		{n: 2, languages: [][]int{{1}, {2}, {1, 2}}, friendships: [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{n: 3, languages: [][]int{{2}, {1, 3}, {1, 2}, {3}}, friendships: [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}}},
	}
	for _, tt := range tests {
		t.Log(minimumTeachings(tt.n, tt.languages, tt.friendships))
	}
}
