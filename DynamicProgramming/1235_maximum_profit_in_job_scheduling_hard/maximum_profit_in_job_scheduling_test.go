package maximumprofitinjobscheduling

import (
	"sort"
	"testing"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]job, 0, n)
	for i := 0; i < n; i++ { // 对job进行封装
		jobs = append(jobs, job{startTime[i], endTime[i], profit[i]})
	} // 按end升序排列jobs
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})
	dp := make([]int, n+1)
	for i, curJob := range jobs {
		// 	二分求j,使 jobs[j].end <= curJob.start
		j := sort.Search(i, func(j int) bool {
			return jobs[j].end > curJob.start
		})
		dp[i+1] = max(dp[i], dp[j]+curJob.profit)
	}
	return dp[n]
}

type job struct {
	start, end int
	profit     int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maximum_profit_in_job_scheduling(t *testing.T) {
	startTime := []int{1, 2, 3, 3}
	endTime := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}
	t.Log(jobScheduling(startTime, endTime, profit))
	startTime = []int{1, 2, 3, 4, 6}
	endTime = []int{3, 5, 10, 6, 9}
	profit = []int{20, 20, 100, 70, 60}
	t.Log(jobScheduling(startTime, endTime, profit))
	startTime = []int{1, 1, 1}
	endTime = []int{2, 3, 4}
	profit = []int{5, 6, 4}
	t.Log(jobScheduling(startTime, endTime, profit))
}
