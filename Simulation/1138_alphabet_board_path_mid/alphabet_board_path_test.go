package alphabetboardpath

import "testing"

func alphabetBoardPath(target string) string {
	ans := []byte{}
	var i, j int
	for _, c := range target {
		v := int(c - 'a')
		x, y := v/5, v%5
		for j > y {
			j--
			ans = append(ans, 'L')
		}
		for i > x {
			i--
			ans = append(ans, 'U')
		}
		for j < y {
			j++
			ans = append(ans, 'R')
		}
		for i < x {
			i++
			ans = append(ans, 'D')
		}
		ans = append(ans, '!')
	}
	return string(ans)
}

func Test_alphabet_board_path(t *testing.T) {
	// board := []string{
	// 	"abcde",
	// 	"fghij",
	// 	"klmno",
	// 	"pqrst",
	// 	"uvwxy",
	// 	"z"}

	target := "leet"
	t.Log(alphabetBoardPath(target))
	target = "code"
	t.Log(alphabetBoardPath(target))
}
