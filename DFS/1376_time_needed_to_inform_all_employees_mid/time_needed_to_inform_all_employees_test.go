package timeneededtoinformallemployees

import "testing"

func numOfMinutes(n int, headID int, manager []int, informTime []int) (ans int) {
	graph := make([][]int, n)
	for to, from := range manager {
		if from >= 0 {
			graph[from] = append(graph[from], to)
		}
	}

	var dfs func(cur int) int
	dfs = func(cur int) (maxSum int) {
		for _, e := range graph[cur] {
			maxSum = max(maxSum, dfs(e))
		}
		return maxSum + informTime[cur]
	}

	return dfs(headID)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_time_needed_to_inform_all_employees(t *testing.T) {
	tests := []struct {
		n          int
		headID     int
		manager    []int
		informTime []int
		wanna      int
	}{
		{1, 0, []int{-1}, []int{0}, 0},
		{6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}, 1},
	}
	for _, tt := range tests {
		t.Log(numOfMinutes(tt.n, tt.headID, tt.manager, tt.informTime) == tt.wanna)
	}
}
