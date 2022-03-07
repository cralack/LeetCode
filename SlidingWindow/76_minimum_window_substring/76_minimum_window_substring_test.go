package minimumwindowsubstring

import "testing"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, char := range t {
		need[byte(char)]++
	}

	left, right := 0, 0
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	start, lenS := 0, int(^uint(0)>>1)
	size := len(s)
	for right < size {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < lenS {
				start = left
				lenS = right - left
			}
			// d 是将移出窗口的字符//d
			d := s[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 返回最小覆盖子串
	if lenS == int(^uint(0)>>1) {
		return ""
	}
	return s[start : start+lenS]
}
func TestMinWindowSubstring(t *testing.T) {
	source := []string{"ADOBECODEBANC", "a", "a", "aa"}
	tar := []string{"ABC", "a", "aa", "aa"}
	for i, val := range source {
		res := minWindow(val, tar[i])
		t.Log(res)
	}
	s := source[0]
	s1 := s[9:13]
	s2 := s[9:12]
	t.Log(s1, s2)
}
