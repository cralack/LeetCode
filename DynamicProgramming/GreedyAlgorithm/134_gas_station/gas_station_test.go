package gasstation

import "testing"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum, minsum := 0, 0
	start := 0
	for i := range gas {
		// 假设累加值最低点为起点
		sum += gas[i] - cost[i]
		if sum < minsum {
			// 经过第 i 个站点后，使 sum 到达新低
			// 所以站点 i + 1 就是最低点（起点）
			minsum = sum
			start = i + 1
		}
	}
	if sum < 0 {
		// 总油量小于总的消耗，无解
		return -1
	}
	// 环形数组特性
	if start == n {
		return 0
	}
	return start
}
func Test_gas_station(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	t.Log(canCompleteCircuit(gas, cost))
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	t.Log(canCompleteCircuit(gas, cost))
	gas = []int{4, 3, 1, 2, 7, 4}
	cost = []int{1, 2, 7, 3, 2, 5}
	t.Log(canCompleteCircuit(gas, cost))

}
