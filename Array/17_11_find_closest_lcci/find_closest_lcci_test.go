package findclosestlcci

import "testing"

func findClosest(words []string, word1 string, word2 string) int {
	dis := len(words)
	pw1, pw2 := -1, -1
	for i, word := range words {
		if word == word1 {
			pw1 = i
		}
		if word == word2 {
			pw2 = i
		}
		if pw1 != -1 && pw2 != -1 && dis > abs(pw1-pw2) {
			dis = abs(pw1 - pw2)
		}
	}
	return dis
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func Test_find_closest_lcci(t *testing.T) {
	words := []string{"I", "am", "a", "student", "from", "a", "university",
		"in", "a", "city"}
	word1 := "a"
	word2 := "student"
	t.Log(findClosest(words, word1, word2))
}
