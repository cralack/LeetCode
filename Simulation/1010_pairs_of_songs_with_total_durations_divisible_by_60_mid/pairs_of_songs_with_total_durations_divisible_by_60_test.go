package pairsofsongswithtotaldurationsdivisibleby60

import "testing"

func numPairsDivisibleBy60(time []int) (ans int) {
	cnt := make([]int, 60)
	for _, dura := range time {
		ans += cnt[(60-dura%60)%60]
		cnt[dura%60]++
	}
	return
}
func Test_pairs_of_songs_with_total_durations_divisible_by_60(t *testing.T) {
	tests := []struct {
		time  []int
		wanna int
	}{
		{[]int{30, 20, 150, 100, 40}, 3},
		{[]int{60, 60, 60}, 3},
	}
	for _, tt := range tests {
		t.Log(numPairsDivisibleBy60(tt.time) == tt.wanna)
	}
}
