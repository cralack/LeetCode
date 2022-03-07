package houserobber

import "testing"

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dpi1, dpi2 := 0, 0
	dpi := 0
	for _, house := range nums {
		dpi = max(dpi1, dpi2+house)
		dpi2 = dpi1
		dpi1 = dpi
	}
	return dpi
}
func Test_house_robber(t *testing.T) {
	streets := [][]int{{4, 3}, {1, 2, 3, 1}, {2, 7, 9, 3, 1}}
	for _, houses := range streets {
		res := rob(houses)
		t.Log(res)
	}

}
