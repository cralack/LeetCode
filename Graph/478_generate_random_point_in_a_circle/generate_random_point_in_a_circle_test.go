package generaterandompointinacircle

import (
	"math/rand"
	"testing"
)

type Solution struct {
	R        float64
	X_center float64
	Y_center float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (p *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y < 1 {
			return []float64{p.X_center + p.R*x, p.Y_center + p.R*y}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
func Test_generate_random_point_in_a_circle(t *testing.T) {
	c1 := []string{"Solution", "randPoint", "randPoint", "randPoint"}
	arg := []float64{1, 0, 0}
	var sol Solution
	for _, c := range c1 {
		switch c {
		case "Solution":
			sol = Constructor(arg[0], arg[1], arg[2])
		case "randPoint":
			t.Log(sol.RandPoint())
		}
	}
}
