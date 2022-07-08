package minimumcosttomovechipstothesamepositioneasy

import "testing"

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, pos := range position {
		if pos%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd < even {
		return odd
	}
	return even
}
func Test_minimum_cost_to_move_chips_to_the_same_position(t *testing.T) {
	position := []int{1, 2, 3}
	t.Log(minCostToMoveChips(position))
	position = []int{2, 2, 2, 3, 3}
	t.Log(minCostToMoveChips(position))
	position = []int{1, 1000000000}
	t.Log(minCostToMoveChips(position))
}
