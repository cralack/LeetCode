package replacewords

import (
	"strings"
	"testing"
)

type Trie struct {
	IsEnd    bool
	Children map[byte]*Trie
}

func (_this *Trie) Add(word string) {
	cur := _this
	for _, ch := range word {
		if cur.Children == nil {
			cur.Children = make(map[byte]*Trie)
		}
		if _, ok := cur.Children[byte(ch)]; !ok {
			cur.Children[byte(ch)] = &Trie{}
		}
		cur = cur.Children[byte(ch)]
	}
	cur.IsEnd = true
}

func (_this *Trie) SearchPrefix(word string) (string, bool) {
	cur := _this
	sub := &strings.Builder{}
	for _, ch := range word {
		if cur.Children == nil {
			return "", false
		}
		if _, ok := cur.Children[byte(ch)]; !ok {
			return "", false
		}
		cur = cur.Children[byte(ch)]
		sub.WriteByte(byte(ch))
		if cur.IsEnd {
			return sub.String(), true
		}
	}
	return "", false
}

func replaceWords(dictionary []string, sentence string) string {
	root := &Trie{}
	for _, dic := range dictionary {
		root.Add(dic)
	}
	res := &strings.Builder{}
	spl := " "
	des := strings.Split(sentence, spl)
	for _, str := range des {
		if sub, ok := root.SearchPrefix(str); ok {
			res.WriteString(sub)
		} else {
			res.WriteString(str)
		}
		res.WriteByte(' ')
	}

	return res.String()[:len(res.String())-1]
}
func Test_replace_words(t *testing.T) {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	str := replaceWords(dictionary, sentence)
	t.Log(str)
	dic := []string{"a", "b", "c"}
	sen := "aadsfasf absbs bbab cadsfafs"
	t.Log(replaceWords(dic, sen))
}
