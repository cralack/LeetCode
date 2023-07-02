package verifyinganaliendictionary

import (
	"testing"
)

func isAlienSorted(words []string, order string) bool {
	dic := make(map[string]int, 0)
	for i, c := range order {
		dic[string(c)] = i + 1
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j == len(words[i+1]) { // len(w1)<len(w2)
				return false
			}
			if diff := dic[string(words[i][j])] - dic[string(words[i+1][j])]; diff > 0 {
				return false
			} else if diff < 0 {
				break
			}
		}
	}
	return true
}
func Test_verifying_an_alien_dictionary(t *testing.T) {
	words, order := []string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"
	t.Log(isAlienSorted(words, order))
	words, order = []string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"
	t.Log(isAlienSorted(words, order))
	words, order = []string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"
	t.Log(isAlienSorted(words, order))
}
