package minimumfuelcosttoreporttothecapital

import (
	"testing"
)

func minimumFuelCost(roads [][]int, seats int) (ans int64) {
	graph := make([][]int, len(roads)+1)
	for _, r := range roads {
		from, to := r[0], r[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	var dfs func(cur int, from int) int
	dfs = func(cur int, from int) int {
		size := 1
		for _, next := range graph[cur] {
			if next != from {
				size += dfs(next, cur)
			}
		}
		if cur > 0 {
			ans += int64((size + seats - 1) / seats)
		}
		return size
	}
	dfs(0, -1)
	return
}

func Test_minimum_fuel_cost_to_report_to_the_capital(t *testing.T) {
	tests := []struct {
		roads [][]int
		seats int
	}{
		{roads: [][]int{{0, 1}, {0, 2}, {0, 3}}, seats: 5},
		{roads: [][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, seats: 2},
		{roads: [][]int{}, seats: 1},
	}
	for _, tt := range tests {
		t.Log(minimumFuelCost(tt.roads, tt.seats))
	}
}
