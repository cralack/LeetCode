package sentencesimilarityiii

import (
	"strings"
	"testing"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Fields(sentence1)
	words2 := strings.Fields(sentence2)
	//keep len(words1)>len(words2)
	if len(words1) < len(words2) {
		words1, words2 = words2, words1
	}
	m, n := len(words1), len(words2)
	i, j := 0, 0
	//从前往后比对
	for i < n && words1[i] == words2[i] {
		i++
	}
	//从后往前比对
	for j < n && words1[m-j-1] == words2[n-j-1] {
		j++
	}
	//双端指针能覆盖短句
	return i+j >= n
}
func Test_sentence_similarity_iii(t *testing.T) {
	sentence1 := "My name is Haley"
	sentence2 := "My Haley"
	t.Log(areSentencesSimilar(sentence1, sentence2))
	sentence1 = "of"
	sentence2 = "A lot of words"
	t.Log(areSentencesSimilar(sentence1, sentence2))
	sentence1 = "Eating right now"
	sentence2 = "Eating"
	t.Log(areSentencesSimilar(sentence1, sentence2))
	sentence1 = "Luky"
	sentence2 = "Lucccky"
	t.Log(areSentencesSimilar(sentence1, sentence2))
	sentence1 = "eTUny i b R UFKQJ EZx JBJ Q xXz"
	sentence2 = "eTUny i R EZx JBJ xXz"
	t.Log(areSentencesSimilar(sentence1, sentence2))
}
