package reorderdatainlogfiles

import (
	"sort"
	"strings"
	"testing"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s, t := logs[i], logs[j]
		s1 := strings.SplitN(s, " ", 2)[1]
		s2 := strings.SplitN(t, " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		if isDigit1 && isDigit2 {
			return false
		}
		if !isDigit1 && !isDigit2 {
			return s1 < s2 || s1 == s2 && s < t
		}
		return !isDigit1
	})
	return logs
}
func Test_reorder_data_in_log_files(t *testing.T) {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6",
		"let2 own kit dig", "let3 art zero"}
	t.Log(reorderLogFiles(logs))
	logs = []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}
	t.Log(reorderLogFiles(logs))
}
