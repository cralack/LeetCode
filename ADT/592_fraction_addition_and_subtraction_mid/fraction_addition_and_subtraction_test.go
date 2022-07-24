package fractionadditionandsubtractionmid

import (
	"strconv"
	"strings"
	"testing"
)

type fraction struct {
	x, y int
}

func parse(str string) fraction {
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	}
	nums := strings.Split(str, "/")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return fraction{sign * x, y}
}
func cal(a, b fraction) fraction {
	x := a.x*b.y + b.x*a.y
	y := a.y * b.y
	gcd := gcd(x, y)
	ans := fraction{x: x / gcd, y: y / gcd}
	if ans.y < 0 {
		ans.y *= -1
		ans.x *= -1
	}
	return ans
}
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
func fra2str(a fraction) string {
	sb := &strings.Builder{}
	sb.WriteString(strconv.Itoa(a.x))
	sb.WriteString("/")
	sb.WriteString(strconv.Itoa(a.y))
	return sb.String()
}

func fractionAddition(expression string) string {
	n := len(expression)
	ans := fraction{x: 0, y: 0}
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n && expression[j] != '+' && expression[j] != '-' {
			j++
		}
		num := parse(expression[i:j])
		if ans.y != 0 {
			ans = cal(ans, num)
		} else {
			ans = num
		}
		i = j - 1
	}
	return fra2str(ans)
}
func Test_fraction_addition_and_subtraction(t *testing.T) {
	//t.Log(gcd(12, 52))
	expression := "-1/2+1/2"
	t.Log(fractionAddition(expression))
	expression = "-1/2+1/2+1/3"
	t.Log(fractionAddition(expression))
	expression = "1/3-1/2"
	t.Log(fractionAddition(expression))
}
