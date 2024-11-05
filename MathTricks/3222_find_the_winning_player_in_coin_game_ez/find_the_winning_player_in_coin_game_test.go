package find_the_winning_player_in_coin_game

import (
	"testing"
)

var player = []string{"Bob", "Alice"}

func losingPlayer(x int, y int) string {
	maxTurn := min(x, y/4)
	return player[maxTurn%2]
}

func Test_find_the_winning_player_in_coin_game(t *testing.T) {
	tests := []struct {
		x, y int
	}{
		{2, 7},
		{4, 11},
	}
	for _, test := range tests {
		t.Log(losingPlayer(test.x, test.y))
	}
}
