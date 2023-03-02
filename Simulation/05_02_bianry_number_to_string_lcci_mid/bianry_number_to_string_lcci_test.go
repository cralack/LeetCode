package bianrynumbertostringlcci

import (
	"strings"
	"testing"
)

func printBin(num float64) string {
	bin := &strings.Builder{}
	bin.WriteString("0.")
	// 题目注明至多循环 6 次
	for i := 0; i < 6; i++ {
		num *= 2
		if num < 1 {
			bin.WriteByte('0')
		} else {
			bin.WriteByte('1')
			num--
			if num == 0 {
				return bin.String()
			}
		}
	}
	return "ERROR"
}
func Test_bianry_number_to_string_lcci(t *testing.T) {
	num := 0.625
	t.Log(printBin(num))
	num = 0.1
	t.Log(printBin(num))
}
