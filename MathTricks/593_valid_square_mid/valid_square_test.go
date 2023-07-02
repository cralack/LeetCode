package validsquaremid

import (
	"testing"
)

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return calc(p1, p2, p3) && calc(p1, p2, p4) &&
		calc(p1, p3, p4) && calc(p2, p3, p4)
}
func calc(a, b, c []int) bool {
	// 三点围成三角形的三边长
	l1 := (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
	l2 := (a[0]-c[0])*(a[0]-c[0]) + (a[1]-c[1])*(a[1]-c[1])
	l3 := (b[0]-c[0])*(b[0]-c[0]) + (b[1]-c[1])*(b[1]-c[1])
	// 能形成等腰&&直角的三角形
	ok := (l1 == l2 && l1+l2 == l3) ||
		(l1 == l3 && l1+l3 == l2) ||
		(l2 == l3 && l2+l3 == l1)
	if l1 == 0 || l2 == 0 || l3 == 0 {
		return false
	}
	return ok
}
func Test_valid_square(t *testing.T) {
	p1 := []int{0, 0}
	p2 := []int{1, 1}
	p3 := []int{1, 0}
	p4 := []int{0, 1}
	t.Log(validSquare(p1, p2, p3, p4))
	p1 = []int{0, 0}
	p2 = []int{1, 1}
	p3 = []int{1, 0}
	p4 = []int{0, 12}
	t.Log(validSquare(p1, p2, p3, p4))
	p1 = []int{1, 0}
	p2 = []int{-1, 0}
	p3 = []int{0, 1}
	p4 = []int{0, -1}
	t.Log(validSquare(p1, p2, p3, p4))
}
