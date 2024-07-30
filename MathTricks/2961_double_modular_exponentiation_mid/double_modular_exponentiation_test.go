package double_modular_exponentiation_mid

import (
	"testing"
)

func getGoodIndices(variables [][]int, target int) (ans []int) {
	for i, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return
}

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func Test_double_modular_exponentiation(t *testing.T) {
	tests := []struct {
		variables [][]int
		target    int
	}{
		{variables: [][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}, target: 2},
		{variables: [][]int{{39, 3, 1000, 1000}}, target: 17},
	}
	for _, test := range tests {
		t.Log(getGoodIndices(test.variables, test.target))
	}
}
