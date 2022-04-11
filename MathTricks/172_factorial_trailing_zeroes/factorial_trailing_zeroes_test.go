package factorialtrailingzeroes

import "testing"

func trailingZeroes(n int) (res int) {
	divisor := 5
	for divisor <= n {
		res += n / divisor
		divisor *= 5
	}
	// for d := n; d/5 > 0; d = d / 5 {
	// 	res += d / 5
	// }
	return res
}

func Test_factorial_trailing_zeroes(t *testing.T) {
	t.Log(trailingZeroes(3))
	t.Log(trailingZeroes(5))
	t.Log(trailingZeroes(125))
	t.Log(trailingZeroes(0))
}
