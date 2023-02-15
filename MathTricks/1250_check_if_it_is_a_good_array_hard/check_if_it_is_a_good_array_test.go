package checkifitisagoodarray

import "testing"

func isGoodArray(nums []int) bool {
	g := 0
	for _, x := range nums {
		g = gcd(x, g)
	}
	return g == 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Test_check_if_it_is_a_good_array(t *testing.T) {
	nums := []int{12, 5, 7, 23}
	t.Log(isGoodArray(nums))
	nums = []int{29, 6, 10}
	t.Log(isGoodArray(nums))
	nums = []int{3, 6}
	t.Log(isGoodArray(nums))
}
