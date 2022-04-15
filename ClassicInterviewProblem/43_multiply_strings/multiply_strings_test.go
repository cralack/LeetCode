package multiplystrings

import (
	"strings"
	"testing"
)

func multiply(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	sb := &strings.Builder{}
	res := make([]int, n+m)
	for i := n - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := m - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			mul := n1 * n2
			p1, p2 := i+j, i+j+1
			sum := mul + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}
	for ; i < len(res); i++ {
		sb.WriteByte(byte(res[i] + '0'))
	}
	if len(sb.String()) == 0 {
		return "0"
	}
	return sb.String()
}
func Test_multiply_strings(t *testing.T) {
	num1, num2 := "123", "456"
	t.Log(multiply(num1, num2))
}
