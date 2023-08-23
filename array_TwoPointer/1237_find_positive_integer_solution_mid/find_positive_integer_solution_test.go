package findpositiveintegersolution

import "testing"

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) (ans [][]int) {
	x, y := 1, 1000
	for x <= 1000 && y > 0 {
		t := customFunction(x, y)
		if t < z {
			x++
		} else if t > z {
			y--
		} else {
			ans = append(ans, []int{x, y})
			x, y = x+1, y-1
		}
	}
	return
}

func cusFunc(id int) func(int, int) int {
	switch id {
	case 1:
		return func(x, y int) int {
			return x + y
		}
	case 2:
		return func(x, y int) int {
			return x * y
		}
	case 3:
		return func(x, y int) int {
			return x*x + y
		}
	case 4:
		return func(x, y int) int {
			return x + y*y
		}
	case 5:
		return func(x, y int) int {
			return x*x + y*y
		}
	case 6:
		return func(x, y int) int {
			return (x + y) * (x + y)
		}
	case 7:
		return func(x, y int) int {
			return x*x*x + y*y*y
		}
	case 8:
		return func(x, y int) int {
			return x * x * y
		}
	case 9:
		return func(x, y int) int {
			return x * y * y
		}
	default:
		return nil
	}

}
func Test_find_positive_integer_solution(t *testing.T) {
	function_id, z := 1, 5
	t.Log(findSolution(cusFunc(function_id), z))
}
