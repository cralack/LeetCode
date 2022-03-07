package fibonacci

import "testing"

func TestFib(t *testing.T) {
	for i := 0; i <= 10; i++ {
		t.Logf("%d : %d", i, fib(i))
		// fib(i)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	pre, cur := 1, 1
	for i := 2; i < n; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}
