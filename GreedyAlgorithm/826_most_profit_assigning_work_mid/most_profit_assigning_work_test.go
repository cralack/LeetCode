package most_profit_assigning_work_mid

import (
	"sort"
	"testing"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) (ans int) {
	type job struct {
		difficulty int
		profit     int
	}
	jobs := make([]*job, len(difficulty))
	for i := range jobs {
		jobs[i] = &job{difficulty[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].profit > jobs[j].profit
	})
	sort.Slice(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})
	j := 0
	for _, w := range worker {
		for j < len(jobs) && w < jobs[j].difficulty {
			j++
		}
		if j == len(jobs) {
			break
		}
		ans += jobs[j].profit
	}

	return
}

func Test_most_profit_assigning_work(t *testing.T) {
	tests := []struct {
		difficulty []int
		profit     []int
		worker     []int
	}{
		{difficulty: []int{2, 4, 6, 8, 10}, profit: []int{10, 20, 30, 40, 50}, worker: []int{4, 5, 6, 7}},
		{difficulty: []int{85, 47, 57}, profit: []int{24, 66, 99}, worker: []int{40, 25, 25}},
	}
	for _, tt := range tests {
		t.Log(maxProfitAssignment(tt.difficulty, tt.profit, tt.worker))
	}
}
