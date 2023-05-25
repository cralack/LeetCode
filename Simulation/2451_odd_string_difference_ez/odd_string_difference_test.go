package oddstringdifference

import "testing"

func oddString(words []string) string {
	dic := make(map[string][]string)
	for _, word := range words {
		diff := make([]byte, len(word)-1)
		for j := 1; j < len(word); j++ {
			diff[j-1] = word[j] - word[j-1]
		}
		dic[string(diff)] = append(dic[string(diff)], word)
	}

	for _, word := range dic {
		if len(word) == 1 {
			return word[0]
		}
	}
	return ""
}

func Test_odd_string_difference(t *testing.T) {
	tests := []struct {
		words []string
	}{
		{[]string{"adc", "wzy", "abc"}},
		{[]string{"aaa", "bob", "ccc", "ddd"}},
		{[]string{"aaaba", "sssts", "vvvwv", "sssts", "ooopo", "rrrsr",
			"iiiji", "pppqp", "aabbb", "xxxyx", "nnnon", "bbbcb", "hhhih",
			"jjjkj", "hhhih", "kkklk", "yyyzy", "jjjkj", "nnnon", "eeefe",
			"eeefe", "ggghg", "sssts", "cccdc", "rrrsr"}},
	}
	for _, tt := range tests {
		t.Log(oddString(tt.words))
	}
}
