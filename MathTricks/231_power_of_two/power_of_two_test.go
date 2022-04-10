package poweroftwo

import "testing"

func isPowerOfTwo(n int) bool {
	cnt := 0
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return cnt == 1
}
func Test_power_of_two(t *testing.T) {
	t.Log(isPowerOfTwo(1))
	t.Log(isPowerOfTwo(16))
	t.Log(isPowerOfTwo(3))
	t.Log(isPowerOfTwo(4))
	t.Log(isPowerOfTwo(5))
}
