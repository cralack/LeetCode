package nimgame

import "testing"

func canWinNim(n int) bool {
	// 如果上来就踩到 4 的倍数，那就认输吧
	// 否则，可以把对方控制在 4 的倍数，必胜
	return n%4 != 0
}
func Test_nim_game(t *testing.T) {
	t.Log(canWinNim(4))
	t.Log(canWinNim(1))
	t.Log(canWinNim(2))
}
