package fraction_to_recurring_decimal_mid

import (
	"strconv"
	"testing"
)

func fractionToDecimal(numerator int, denominator int) string {
	sign := ""
	if numerator*denominator < 0 {
		sign = "-"
	}
	numerator, denominator = abs(numerator), abs(denominator)

	q, r := numerator/denominator, numerator%denominator
	if r == 0 {
		return sign + strconv.Itoa(q)
	}
	ans := []byte(sign + strconv.Itoa(q) + ".")
	rToPos := map[int]int{r: len(ans)}
	for r != 0 {
		r *= 10
		q = r / denominator
		r %= denominator
		ans = append(ans, byte(q)+'0')
		if pos, ok := rToPos[r]; ok {
			return string(ans[:pos]) + "(" + string(ans)[pos:] + ")"
		}
		rToPos[r] = len(ans)
	}
	return string(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Test_fraction_to_recurring_decimal(t *testing.T) {
	tests := []struct {
		numerator   int
		denominator int
	}{
		{numerator: 1, denominator: 2},
		{numerator: 2, denominator: 1},
		{numerator: 4, denominator: 333},
		{numerator: -50, denominator: 8},
	}
	for _, tt := range tests {
		t.Log(fractionToDecimal(tt.numerator, tt.denominator))
	}
}
