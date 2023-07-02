package perfectrectangle

import "testing"

func isRectangleCover(rectangles [][]int) bool {
	type point struct{ x, y int }
	points := make(map[point]bool, 0)
	xMin, yMin := 999999, 999999 // 完美矩形左下角坐标
	xMax, yMax := 0, 0           // 完美矩形右上角坐标
	sumArea := 0                 // 矩形数组面积和
	for _, rect := range rectangles {
		x, y := rect[0], rect[1]
		a, b := rect[2], rect[3]
		xMin, xMax = min(xMin, x), max(xMax, a) // 完美矩形的理论顶点坐标
		yMin, yMax = min(yMin, y), max(yMax, b)
		sumArea += (a - x) * (b - y)
		// 单个小矩形的4顶点
		p1, p2 := point{x, y}, point{a, b}
		p3, p4 := point{a, y}, point{x, b}
		// 矩形顶点 成对相消
		for _, p := range []point{p1, p2, p3, p4} {
			if _, ok := points[p]; ok {
				delete(points, p)
			} else {
				points[p] = true
			}
		}
	}
	if sumArea != (xMax-xMin)*(yMax-yMin) { // 面积不等肯定不对
		return false
	}
	if len(points) != 4 { // 判断最终留下的顶点个数不是4也不对
		return false
	}
	// 剩下的4个顶点和完美矩形4个顶点不能一一对应了也不对
	p1, p2 := point{xMin, yMin}, point{xMax, yMax}
	p3, p4 := point{xMin, yMax}, point{xMax, yMin}
	for _, p := range []point{p1, p2, p3, p4} {
		if _, ok := points[p]; !ok {
			return false
		}
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_perfect_rectangle(t *testing.T) {
	rectangles := [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}
	t.Log(isRectangleCover(rectangles))
	rectangles = [][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}
	t.Log(isRectangleCover(rectangles))
	rectangles = [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}}
	t.Log(isRectangleCover(rectangles))
}
