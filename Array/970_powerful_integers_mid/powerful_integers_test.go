package powerfulintegers

import "testing"

func powerfulIntegers(x int, y int, bound int) (ans []int) {
	dic := make(map[int]bool)
	for a := 1; a <= bound; a *= x {
		for b := 1; a+b <= bound; b *= y {
			dic[a+b] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	for key := range dic {
		ans = append(ans, key)
	}
	return
}
func Test_powerful_integers(t *testing.T) {
	tests := []struct {
		x     int
		y     int
		bound int
	}{
		{2, 3, 10},
		{3, 5, 15},
	}

	for _, tt := range tests {
		t.Log(powerfulIntegers(tt.x, tt.y, tt.bound))
	}
}
