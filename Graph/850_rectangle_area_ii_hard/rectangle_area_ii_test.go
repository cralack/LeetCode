package rectangleareaiihard

import (
	"sort"
	"testing"
)

const MOD = int64(1e9 + 7)

func rectangleArea(rectangles [][]int) int {
	list := []int{}
	for _, info := range rectangles {
		list = append(list, info[0])
		list = append(list, info[2])
	}
	sort.Ints(list)
	ans := int64(0)
	for i := 1; i < len(list); i++ {
		a, b, length := list[i-1], list[i], list[i]-list[i-1]
		if length == 0 {
			continue
		}
		lines := [][]int{}
		for _, info := range rectangles {
			if info[0] <= a && b <= info[2] {
				lines = append(lines, []int{info[1], info[3]})
			}
		}
		sort.Slice(lines, func(i, j int) bool {
			if lines[i][0] != lines[j][0] {
				return lines[i][0]-lines[j][0] < 0
			}
			return lines[i][1]-lines[j][1] < 0
		})
		total, l, r := int64(0), -1, -1
		for _, cur := range lines {
			if cur[0] > r {
				total += int64(r - l)
				l, r = cur[0], cur[1]
			} else if cur[1] > r {
				r = cur[1]
			}
		}
		total += int64(r - l)
		ans += total * int64(length)
		ans %= MOD
	}
	return int(ans)
}

func Test_rectangle_area_ii(t *testing.T) {
	rectangles := [][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}
	t.Log(rectangleArea(rectangles))
	rectangles = [][]int{{0, 0, 1000000000, 1000000000}}
	t.Log(rectangleArea(rectangles))
}
