package count_the_number_of_vowel_strings_in_range

import (
	"testing"
)

func vowelStrings(words []string, left int, right int) (ans int) {
	valid := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	for _, word := range words[left : right+1] {
		_, va1 := valid[word[0]]
		_, va2 := valid[word[len(word)-1]]
		if va1 && va2 {
			ans++
		}

	}

	return
}

func Test_123(t *testing.T) {
	tests := []struct {
		words []string
		left  int
		right int
	}{
		{words: []string{"are", "amy", "u"}, left: 0, right: 2},
		{words: []string{"hey", "aeo", "mu", "ooo", "artro"}, left: 1, right: 4},
	}
	for _, tt := range tests {
		t.Log(vowelStrings(tt.words, tt.left, tt.right))
	}
}
