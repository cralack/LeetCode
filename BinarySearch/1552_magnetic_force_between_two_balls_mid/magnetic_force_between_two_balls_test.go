package magnetic_force_between_two_balls_mid

import (
	"sort"
	"testing"
)

func maxDistance(position []int, m int) int {
	sort.Ints(position)

	// 返回使cnt<m的最小磁力
	return sort.Search(position[len(position)-1], func(mid int) bool {
		prev := position[0]
		cnt := 1
		for _, cur := range position {
			if cur-prev >= mid {
				cnt++
				prev = cur
			}
		}
		return cnt < m
	}) - 1
}

func Test_(t *testing.T) {
	tests := []struct {
		position []int
		m        int
		want     int
	}{
		{position: []int{1, 2, 3, 4, 7}, m: 3, want: 3},
		{position: []int{5, 4, 3, 2, 1, 1000000000}, m: 2, want: 999999999},
	}
	for _, tt := range tests {
		t.Log(maxDistance(tt.position, tt.m) == tt.want)
	}
}
