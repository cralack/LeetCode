package findthewinnerofthecirculargame

import "testing"

func findTheWinner(n int, k int) int {
	winner := 1
	for i := 2; i <= n; i++ {
		winner = (k+winner-1)%i + 1
	}
	return winner
}
func Test_find_the_winner_of_the_circular_game(t *testing.T) {
	n, k := 5, 2
	t.Log(findTheWinner(n, k))
	n, k = 6, 5
	t.Log(findTheWinner(n, k))
}
