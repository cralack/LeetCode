package asteroidcollisionmid

import (
	"testing"
)

func asteroidCollision(asteroids []int) (stack []int) {
	for _, cur := range asteroids {
		alive := true
		for alive && cur < 0 && 0 < len(stack) && 0 < stack[len(stack)-1] {
			alive = stack[len(stack)-1] < -cur
			if stack[len(stack)-1] <= -cur {
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, cur)
		}
	}
	return stack
}
func Test_asteroid_collision(t *testing.T) {
	asteroids := []int{5, 10, -5}
	t.Log(asteroidCollision(asteroids))
	asteroids = []int{8, -8}
	t.Log(asteroidCollision(asteroids))
	asteroids = []int{10, 2, -5}
	t.Log(asteroidCollision(asteroids))
	asteroids = []int{-2, -1, 1, 2}
	t.Log(asteroidCollision(asteroids))
}
