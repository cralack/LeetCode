package heaters

import (
	"sort"
	"testing"
)

func TestHeaters(t *testing.T) {
	houses := [][]int{{1, 2, 3}, {1, 3, 2, 4}, {1, 5}}
	heaters := [][]int{{2}, {1, 4}, {2}}
	ans := []int{}
	for idx, house := range houses {
		ans = append(ans, findRadius(house, heaters[idx]))
	}
	t.Log(ans)
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	ans, j := 0, 0
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for _, house := range houses {
		dis := abs(house - heaters[j])
		// j！=last heaters
		for j+1 < len(heaters) &&
			// house到左heater不小于到右heater距离
			abs(house-heaters[j]) >= abs(house-heaters[j+1]) {
			j++
			if abs(house-heaters[j]) < dis {
				dis = abs(house - heaters[j])
			}
		}
		if dis > ans {
			ans = dis
		}
	}
	return ans
}
