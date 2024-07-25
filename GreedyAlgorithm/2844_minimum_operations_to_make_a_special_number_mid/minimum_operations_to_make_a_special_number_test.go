package minimum_operations_to_make_a_special_number_mid

import (
	"testing"
)

func minimumOperations(num string) int {
	n := len(num)
	var found0, found5 bool
	for i := n - 1; i >= 0; i-- {
		c := num[i]
		if found0 && (c == '0' || c == '5') ||
			found5 && (c == '2' || c == '7') {
			return n - i - 2
		}
		if c == '0' {
			found0 = true
		} else if c == '5' {
			found5 = true
		}
	}
	if found0 {
		return n - 1
	}
	return n
}

func Test_minimum_operations_to_make_a_special_number(t *testing.T) {
	tests := []struct {
		num   string
		wannt int
	}{
		{num: "2245047", wannt: 2},
		{num: "2908305", wannt: 3},
		{num: "10", wannt: 1},
	}
	for _, test := range tests {
		t.Log(minimumOperations(test.num) == test.wannt)
	}

}
