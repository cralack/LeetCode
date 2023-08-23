package checkifawordoccursasaprefixofanywordinasentenceez

import (
	"strings"
	"testing"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	strs := strings.Split(sentence, " ")
	for idx, str := range strs {
		if len(str) >= len(searchWord) &&
			str[0:len(searchWord)] == searchWord {
			return idx + 1
		}
	}
	return -1
}
func Test_check_if_prefix(t *testing.T) {
	sentence := "i love eating burger"
	searchWord := "burg"
	t.Log(isPrefixOfWord(sentence, searchWord))
	sentence = "this problem is an easy problem"
	searchWord = "pro"
	t.Log(isPrefixOfWord(sentence, searchWord))
	sentence = "i am tired"
	searchWord = "you"
	t.Log(isPrefixOfWord(sentence, searchWord))
	sentence = "hello from the other side"
	searchWord = "they"
	t.Log(isPrefixOfWord(sentence, searchWord))
	sentence = "love errichto jonathan dumb"
	searchWord = "dumb"
	t.Log(isPrefixOfWord(sentence, searchWord))
}
