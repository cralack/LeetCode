package numberofmatchingsub

import (
	"testing"
)

func numMatchingSubseq_v1(s string, words []string) (ans int) {
	dic := make(map[byte][]int, 0)
	for i := range s {
		dic[s[i]] = append(dic[s[i]], i)
	}

	for _, word := range words {
		ok := true
		m, idx := len(word), -1
		for i := 0; i < m && ok; i++ {
			curDic := dic[word[i]]
			//binSearch
			left, right := 0, len(curDic)-1
			var mid int
			for left < right {
				mid = left + (right-left)>>1
				if curDic[mid] > idx {
					right = mid
				} else {
					left = mid + 1
				}
			}
			// 如果当前字符在原串不存在 || 当前下标满足递增要求，
			//e.g. 当前匹配 "abc", a在原串index = 4, b在原串的index = 2 ，不符合题意
			if right < 0 || curDic[right] <= idx {
				ok = false
			} else { // 匹配正确，下标递增
				idx = curDic[right]
			}
		}
		if ok {
			ans++
		}
	}
	return
}

// 分桶
func numMatchingSubseq_v2(s string, words []string) (ans int) {
	dic := [26][]string{}
	for _, word := range words {
		dic[word[0]-'a'] = append(dic[word[0]-'a'], word)
	}
	for _, c := range s {
		que := dic[c-'a']
		dic[c-'a'] = nil
		for _, t := range que {
			if len(t) == 1 {
				ans++
			} else {
				dic[t[1]-'a'] = append(dic[t[1]-'a'], t[1:])
			}
		}
	}
	return
}

var pat = "rwpddkvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvjubjgloeofnpjqlkdsqvruvabjrikfwronbrdyyjnakstqjac"
var txt = []string{"wpddkvbnn", "lnagtva",
	"kvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvju",
	"rwpddkvbnnugln", "gloeofnpjqlkdsqvruvabjrikfwronbrdyyj",
	"vbgeinupkvgmgxeaaiuiyojmoqkahwvbpwugdainxciedbdkos",
	"mspuhbykmmumtveoighlcgpcapzczomshiblnvhjzqjlfkpina",
	"rgmliajkiknongrofpugfgajedxicdhxinzjakwnifvxwlokip",
	"fhepktaipapyrbylskxddypwmuuxyoivcewzrdwwlrlhqwzikq",
	"qatithxifaaiwyszlkgoljzkkweqkjjzvymedvclfxwcezqebx"}

func Test_number_of_matching_subs(t *testing.T) {
	s := "abcde"
	words := []string{"a", "bb", "acd", "ace"}
	t.Log(numMatchingSubseq_v1(s, words))
	t.Log(numMatchingSubseq_v2(s, words))

	s = "dsahjpjauf"
	words = []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}
	t.Log(numMatchingSubseq_v1(s, words))
	t.Log(numMatchingSubseq_v2(s, words))

	t.Log(numMatchingSubseq_v1(pat, txt))
	t.Log(numMatchingSubseq_v2(pat, txt))
}

func Benchmark_matching_subs(b *testing.B) {
	b.Run("v1", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			numMatchingSubseq_v1(pat, txt)
		}
		b.StopTimer()
	})
	b.Run("v2", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			numMatchingSubseq_v2(pat, txt)
		}
		b.StopTimer()
	})
}
