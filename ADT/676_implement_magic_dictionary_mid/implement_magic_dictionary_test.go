package implementmagicdictionarymid

import (
	"testing"
)

type MagicDictionary struct {
	Dic []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{Dic: make([]string, 0)}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	m.Dic = dictionary
}

func (m *MagicDictionary) Search(searchWord string) bool {
	for _, dic := range m.Dic {
		if len(dic) != len(searchWord) || searchWord == dic {
			continue
		} else {
			cnt := 0
			for i := 0; i < len(dic); i++ {
				if dic[i] != searchWord[i] {
					cnt++
				}
			}
			if cnt == 1 {
				return true
			}
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
func Test_implement_magic_dictionary(t *testing.T) {
	c1 := []string{"MagicDictionary", "buildDict",
		"search", "search", "search", "search"}
	c2 := [][]string{{}, {"hello", "leetcode"},
		{"hello"}, {"hhllo"}, {"hell"}, {"leetcoded"}}
	var sol MagicDictionary
	for i, c := range c1 {
		switch c {
		case "MagicDictionary":
			sol = Constructor()
		case "buildDict":
			sol.BuildDict(c2[i])
		case "search":
			t.Log(sol.Search(c2[i][0]))
		}
	}
}
