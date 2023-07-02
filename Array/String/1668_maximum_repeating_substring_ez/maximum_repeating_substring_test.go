package maximumrepeatingsubstringez

import "testing"

func maxRepeating(sequence string, word string) (ans int) {
	cnt, n := 0, len(word)
	i := 0
	for i <= len(sequence)-n {
		if sequence[i:i+n] == word { // 能重复就按步长n重复
			cnt++
			i += n
		} else { // 否则cnt清零
			i -= cnt * n
			i++
			cnt = 0
		}
		if ans < cnt {
			ans = cnt
		}
	}
	return
}

func Test_maximum_repeating_substring(t *testing.T) {
	sequence := "ababc"
	word := "ab"
	t.Log(maxRepeating(sequence, word))
	sequence = "ababc"
	word = "ba"
	t.Log(maxRepeating(sequence, word))
	sequence = "a"
	word = "a"
	t.Log(maxRepeating(sequence, word))
	sequence = "aaa"
	word = "a"
	t.Log(maxRepeating(sequence, word))
	sequence = "ababc"
	word = "ac"
	t.Log(maxRepeating(sequence, word))
	// "aaaba,aaaba,aaba,aaaba,aaaba,aaaba,aaaba"
	sequence = "aaabaaaabaaabaaaabaaaabaaaabaaaaba"
	word = "aaaba"
	t.Log(maxRepeating(sequence, word))
}
