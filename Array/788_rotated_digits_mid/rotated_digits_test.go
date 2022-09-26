package rotateddigits

import "testing"

func rotatedDigits(n int) (ans int) {
	check := []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	for i := 1; i <= n; i++ {
		valid, diff := true, false
		num := i
		for num > 0 {
			tmp := num % 10
			if check[tmp] == -1 {
				valid = false
			} else if check[tmp] == 1 {
				diff = true
			}
			num /= 10
		}
		if valid && diff {
			ans++
		}
	}
	return
}
func Test_rotated_digits(t *testing.T) {
	n := 10
	t.Log(rotatedDigits(n))
	n = 857
	t.Log(rotatedDigits(n))
}
