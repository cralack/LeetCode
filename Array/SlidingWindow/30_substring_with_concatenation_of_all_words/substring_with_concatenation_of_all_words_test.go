package substringwithconcatenationofallwords

import (
	"testing"
)

func findSubstring(s string, words []string) (ans []int) {
	n, m, w := len(s), len(words), len(words[0])
	// 统计 words 中「每个目标单词」的出现次数
	totalCnt := 0
	totalMap := make(map[string]int, 0)
	for _, word := range words {
		if _, has := totalMap[word]; !has {
			totalCnt++
		}
		totalMap[word]++
	}
	for i := 0; i < w; i++ {
		//copy(curMap,totalMap)
		cnt := 0
		curMap := make(map[string]int, 0)
		// 滑动窗口的大小固定是 m * w，
		// 每次将下一个单词添加进 curMap，上一个单词移出 curMap
		for j := i; j+w <= n; j += w {
			cur := s[j : j+w]
			curMap[cur]++
			if _, has := totalMap[cur]; has && curMap[cur] == totalMap[cur] {
				cnt++
			}
			if j >= i+m*w {
				idx := j - m*w
				prev := s[idx : idx+w]
				curMap[prev]--
				if curMap[prev]+1 == totalMap[prev] {
					cnt--
				}
				if curMap[prev] == 0 {
					delete(curMap, prev)
				}
				if curMap[prev] != totalMap[prev] {
					continue
				}
			}
			if curMap[cur] != totalMap[cur] {
				continue
			}
			if len(curMap) == len(totalMap) && cnt == totalCnt {
				ans = append(ans, j-(m-1)*w)
			}
		}
	}
	return
}
func Test_substring_with_concatenation_of_all_words(t *testing.T) {
	s, words := "barfoothefoobarman", []string{"foo", "bar"}
	t.Log(findSubstring(s, words))
	s, words = "wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}
	t.Log(findSubstring(s, words))
	s, words = "barfoofoobarthefoobarman", []string{"foo", "bar", "the"}
	t.Log(findSubstring(s, words))
	s, words = "wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}
	t.Log(findSubstring(s, words))
	s, words = "bcabbcaabbccacacbabccacaababcbb", []string{"c", "b", "a", "c", "a", "a", "a", "b", "c"}
	t.Log(findSubstring(s, words))
}
