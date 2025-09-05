package minimumoperationstomaketheintegerzeromid

import (
	"math/bits"
	"testing"
)

func makeTheIntegerZero(num1, num2 int) int {
	for k := 1; k <= num1-num2*k; k++ {
		tar := num1 - num2*k
		if k >= bits.OnesCount(uint(tar)) {
			return k
		}
	}
	return -1
}

func Test_minimum_operations_to_make_the_integer_zero(t *testing.T) {
	tests := []struct {
		num1 int
		num2 int
	}{
		{num1: 3, num2: -2},
		{num1: 5, num2: 7},
	}
	for _, tt := range tests {
		t.Log(makeTheIntegerZero(tt.num1, tt.num2))
	}
}
