package bulbswitcher

import (
	"math"
	"testing"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
func Test_bulb_switcher(t *testing.T) {
	t.Log(bulbSwitch(3))
	t.Log(bulbSwitch(0))
	t.Log(bulbSwitch(1))
}
