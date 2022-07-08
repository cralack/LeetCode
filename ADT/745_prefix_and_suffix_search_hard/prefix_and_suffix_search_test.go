package prefixandsuffixsearchhard

import "testing"

type WordFilter map[string]int

func Constructor(words []string) WordFilter {
	wf := WordFilter{}
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			for k := 0; k < len(word); k++ {
				wf[word[:j]+"#"+word[k:]] = i
			}
		}
	}
	return wf
}

func (this *WordFilter) F(pref string, suff string) int {
	dic := *this
	if idx, has := dic[pref+"#"+suff]; has {
		return idx
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
func Test_prefix_and_suffix_search(t *testing.T) {
	c1 := []string{"WordFilter", "f"}
	c2 := [][]string{{"apple"}, {"a", "e"}}
	var sol WordFilter
	for i, c := range c1 {
		switch c {
		case "WordFilter":
			sol = Constructor(c2[0])
		case "f":
			t.Log(sol.F(c2[i][0], c2[i][1]))
		}
	}
	c2 = [][]string{{"abbba", "abba", "sad"}, {"ab", "ba"}}
	for i, c := range c1 {
		switch c {
		case "WordFilter":
			sol = Constructor(c2[0])
		case "f":
			t.Log(sol.F(c2[i][0], c2[i][1]))
		}
	}
}
