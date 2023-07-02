package statisticsfromalargesample

import "testing"

func sampleStats(count []int) []float64 {
	max, min := -1, -1
	flag := false
	sum, cnt := 0, 0
	mode := 0
	for i, c := range count {
		if c > 0 {
			cnt += c
			sum += c * i
			max = i
			if !flag {
				min = i
				flag = true
			}
			if c > count[mode] {
				mode = i
			}
		}
	}

	// find median
	find := func(idx int) int {
		for i, c := 0, 0; ; i++ {
			c += count[i]
			if c >= idx {
				return i
			}
		}
	}
	var median float64
	if cnt&1 == 1 {
		median = float64(find(cnt/2 + 1))
	} else {
		median = float64(find(cnt/2+1)+find(cnt/2)) / 2
	}

	// [minimum, maximum, mean, median, mode]
	return []float64{float64(min), float64(max), float64(sum) / float64(cnt),
		median, float64(mode)}
}
func Test_statistics_from_a_large_sample(t *testing.T) {
	tests := []struct {
		count []int
	}{
		{count: []int{0, 1, 3, 4, 0, 0, 0}},
		{count: []int{0, 4, 3, 2, 2, 0, 0}},
	}
	for _, tt := range tests {
		ans := sampleStats(tt.count)
		t.Log(ans)
	}
}
