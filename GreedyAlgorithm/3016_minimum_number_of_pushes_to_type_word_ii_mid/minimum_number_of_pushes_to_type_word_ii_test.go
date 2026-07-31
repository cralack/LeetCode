package minimum_number_of_pushes_to_type_word_ii_mid

import (
	"slices"
	"testing"
)

func minimumPushes(word string) (ans int) {
	cnt := [26]int{}
	for _, ch := range word {
		cnt[ch-'a']++
	}
	slices.SortFunc(cnt[:], func(a, b int) int {
		return b - a
	})
	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return
}

func Test_minimum_number_of_pushes_to_type_word_ii(t *testing.T) {
	tests := []struct {
		word string
	}{
		{"abcde"},
		{"xyzxyzxyzxyz"},
		{"aabbccddeeffgghhiiiiii"},
	}
	for _, test := range tests {
		t.Log(minimumPushes(test.word))
	}
}
