package expressivewords

import "testing"

func expressiveWords(s string, words []string) (ans int) {
	for _, word := range words {
		if check(word, s) {
			ans++
		}
	}
	return
}

// e.g  txt:"hello" pat:"heeellooo"
func check(txt, pat string) bool {
	if len(txt) > len(pat) {
		return false
	}
	pi, pj := 0, 0
	for pi < len(txt) && pj < len(pat) {
		if txt[pi] != pat[pj] {
			return false
		}
		// 相等字符连续长度
		char := txt[pi]
		cntTxt, cntPat := 0, 0
		for pi < len(txt) && txt[pi] == char {
			cntTxt++
			pi++
		}
		for pj < len(pat) && pat[pj] == char {
			cntPat++
			pj++
		}
		// 单次扩张长度为3
		if cntPat < cntTxt || (cntPat != cntTxt) && cntPat < 3 {
			return false
		}
	}
	return pi == len(txt) && pj == len(pat)
}

func Test_expressive_words(t *testing.T) {
	s := "heeellooo"
	words := []string{"hello", "hi", "helo"}
	t.Log(expressiveWords(s, words))
	s = "abcd"
	words = []string{"abc"}
	t.Log(expressiveWords(s, words))
}
