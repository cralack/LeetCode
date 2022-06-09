package findandreplacepattern

import (
	"testing"
)

func findAndReplacePattern(words []string, pattern string) (ans []string) {
	for _, word := range words {
		if match(word, pattern) && match(pattern, word) {
			ans = append(ans, word)
		}
	}
	return
}
func match(word, pattern string) bool {
	set := map[rune]byte{}
	for i, x := range word {
		y := pattern[i]
		if set[x] == 0 {
			set[x] = y
		} else if set[x] != y {
			return false
		}
	}
	return true
}
func Test_find_and_replace_pattern(t *testing.T) {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	t.Log(findAndReplacePattern(words, pattern))
}
