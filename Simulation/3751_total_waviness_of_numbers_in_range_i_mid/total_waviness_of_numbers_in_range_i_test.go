package total_waviness_of_numbers_in_range_i_mid

import (
	"strconv"
	"testing"
)

func totalWaviness(num1 int, num2 int) (ans int) {
	for x := num1; x <= num2; x++ {
		if x < 100 {
			continue
		}
		str := strconv.Itoa(x)
		for i := 1; i < len(str)-1; i++ {
			if str[i-1] < str[i] && str[i] > str[i+1] {
				ans++
			} else if str[i-1] > str[i] && str[i] < str[i+1] {
				ans++
			}
		}
	}
	return
}

func Test_total_waviness_of_numbers_in_range_i(t *testing.T) {
	tests := []struct {
		num1 int
		num2 int
	}{
		{120, 130},
		{198, 202},
		{4848, 4848},
	}
	for _, tt := range tests {
		t.Log(totalWaviness(tt.num1, tt.num2))
	}
}
