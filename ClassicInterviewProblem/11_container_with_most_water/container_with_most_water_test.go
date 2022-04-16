package containerwithmostwater

import "testing"

func maxArea(height []int) int {
	n, max_area, area := len(height), 0, 0
	left, right := 0, n-1
	for left < right {
		area = min(height[left], height[right]) * (right - left)
		if max_area < area {
			max_area = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max_area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_container_with_most_water(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(height))
	height = []int{1, 1}
	t.Log(maxArea(height))
}
