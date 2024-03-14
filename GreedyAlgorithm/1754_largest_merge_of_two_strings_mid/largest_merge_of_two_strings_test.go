package largestmergeoftwostrings

import "testing"

func largestMerge(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	ans := make([]byte, 0, m+n)
	i, j := 0, 0
	for i < n && j < m {
		if word1[i:] > word2[j:] {
			ans = append(ans, word1[i])
			i++
		} else {
			ans = append(ans, word2[j])
			j++
		}
	}
	ans = append(ans, word1[i:]...)
	ans = append(ans, word2[j:]...)
	return string(ans)
}
func Test_largest_merge_of_two_strings(t *testing.T) {
	word1 := "cabaa"
	word2 := "bcaaa"
	t.Log(largestMerge(word1, word2))
	word1 = "abcabc"
	word2 = "abdcaba"
	t.Log(largestMerge(word1, word2))
}
