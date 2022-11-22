package nthmagicalnumber

import "testing"

func nthMagicalNumber(n int, a int, b int) int {
	const mod int = 1e9 + 7
	//最小公倍数= (a*b)/gcd(a,b)
	lcm := a * b / gcd(a, b)
	left, right := min(a, b), min(a, b)*n
	for left < right {
		mid := left + (right-left)>>1
		curCnt := mid/a + mid/b - mid/lcm
		if curCnt >= n {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func Test_nth_magical_number(t *testing.T) {
	n, a, b := 1, 2, 3
	t.Log(nthMagicalNumber(n, a, b))
	n, a, b = 4, 2, 3
	t.Log(nthMagicalNumber(n, a, b))
	n, a, b = 5, 2, 4
	t.Log(nthMagicalNumber(n, a, b))
	n, a, b = 8, 4, 6
	t.Log(nthMagicalNumber(n, a, b))
}
