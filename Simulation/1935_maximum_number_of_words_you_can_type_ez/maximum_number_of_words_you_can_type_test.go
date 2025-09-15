package maximum_number_of_words_you_can_type_ez

import (
	"strings"
	"testing"
)

func canBeTypedWords(text string, brokenLetters string) (ans int) {
	mask := 0
	for _, b := range brokenLetters {
		mask |= 1 << int(b&31)
	}
	for _, s := range strings.Split(text, " ") {
		flag := 1
		for _, ch := range s {
			if 1<<(ch&31)&mask > 0 {
				flag = 0
				break
			}
		}
		ans += flag
	}
	return
}

func Test_maximum_number_of_words_you_can_type(t *testing.T) {
	tests := []struct {
		text          string
		brokenLetters string
		want          int
	}{
		{text: "hello world", brokenLetters: "ad", want: 1},
		{text: "leet code", brokenLetters: "lt", want: 1},
		{text: "leet code", brokenLetters: "e", want: 0},
	}
	for _, tt := range tests {
		t.Log(canBeTypedWords(tt.text, tt.brokenLetters) == tt.want)
	}
}
