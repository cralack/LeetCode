package minimumnumberofrefuelingstopshard

import (
	"container/heap"
	"sort"
	"testing"
)

func minRefuelStops_DP(target int, startFuel int, stations [][]int) int {
	if startFuel >= target {
		return 0
	}
	n := len(stations)
	/*dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = startFuel
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i][j]: 前i个加油站，加油<=j次，能到达的最远距离
	for i := 1; i <= n; i++ {
		sta := stations[i-1]
		for j := 1; j <= n; j++ {
			if dp[i-1][j-1] < sta[0] {
				// canot make it
				dp[i][j] = dp[i-1][j]
			} else {
				// we made it
				dp[i][j] = max(dp[i-1][j-1]+sta[1], dp[i-1][j]) //refuel or not
			}
		}
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if dp[j][i] >= target {
				return i
			}
		}
	}*/

	//一维空间压缩
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i := 1; i <= n; i++ {
		sta := stations[i-1]
		for j := i; j >= 1; j-- {
			if dp[j-1] >= sta[0] && dp[j-1]+sta[1] > dp[j] {
				dp[j] = dp[j-1] + sta[1]
			}
		}
	}
	for i, v := range dp {
		if v >= target {
			return i
		}
	}

	return -1
}

//greedy algorithm + priority queue
func minRefuelStops_greedy(target int, startFuel int, stations [][]int) (ans int) {
	pq := &IntHeap{}
	for cur, idx, fuel := 0, 0, startFuel; cur < target; {
		cur += fuel
		for idx < len(stations) && stations[idx][0] <= cur {
			heap.Push(pq, stations[idx][1])
			idx++
		}
		if cur < target {
			if pq.Len() == 0 {
				return -1
			}
			ans++
			fuel = heap.Pop(pq).(int)
		}
	}
	return
}

//PriorityQueue实现
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	//按油量多少降序排列
	sort.Slice(stations, func(i, j int) bool {
		if stations[i][1] == stations[j][1] {
			return stations[i][0] < stations[j][0]
		}
		return stations[i][1] > stations[j][1]
	})
	n, res := len(stations), 0
	//遍历油站
	for i := 0; i < n && startFuel < target; i++ {
		//油站未被使用(标0)&&当前油量足够到达
		if 0 < stations[i][0] && stations[i][0] <= startFuel {
			startFuel += stations[i][1]
			stations[i][0] = 0
			res++
			i = -1
		}
	}
	if startFuel < target {
		//用光所有汽油后
		return -1
	}
	return res
}

func Test_minimum_number_of_refueling_stops(t *testing.T) {
	target, startFuel, stations := 100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
	t.Log(minRefuelStops(target, startFuel, stations))
	target, startFuel, stations = 100, 1, [][]int{{10, 100}}
	t.Log(minRefuelStops(target, startFuel, stations))
	target, startFuel, stations = 1, 1, [][]int{{}}
	t.Log(minRefuelStops(target, startFuel, stations))
}
