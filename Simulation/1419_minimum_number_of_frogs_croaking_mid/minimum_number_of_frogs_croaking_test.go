package minimumnumberoffrogscroaking

import "testing"

func minNumberOfFrogs(croakOfFrogs string) int {
	// 每个字母在 "croak" 中的上一个字母
	previous := map[rune]rune{
		'c': 'k',
		'r': 'c',
		'o': 'r',
		'a': 'o',
		'k': 'a',
	}

	cnt := make(map[rune]int)
	for _, ch := range croakOfFrogs {
		pre := previous[ch] // pre 为 ch 在 "croak" 中的上一个字母
		if cnt[pre] > 0 {   // 如果有青蛙发出了 pre 的声音
			cnt[pre]-- // 复用一只
		} else if ch != 'c' { // 否则青蛙必须从 'c' 开始蛙鸣
			return -1 // 不符合要求
		}
		cnt[ch]++ // 发出了 ch 的声音
	}
	if cnt['c'] > 0 || cnt['r'] > 0 || cnt['o'] > 0 || cnt['a'] > 0 {
		return -1 // 有发出其它声音的青蛙，不符合要求
	}
	return cnt['k'] // 最后青蛙们都发出了 'k' 的声音
}

func Test_minimum_number_of_frogs_croaking(t *testing.T) {
	tests := []struct {
		croakOfFrogs string
	}{
		{"croakcroak"},
		{"crcoakroak"},
		{"croakcrook"},
	}
	for _, tt := range tests {
		t.Log(minNumberOfFrogs(tt.croakOfFrogs))
	}
}
