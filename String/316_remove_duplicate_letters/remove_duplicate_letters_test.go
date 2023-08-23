package removeduplicateletters

import (
	"testing"
)

func removeDuplicateLetters(s string) string {
	stack := []byte{}
	inStack := make(map[byte]bool, 0)
	cnt := make(map[byte]int, 0)
	for _, c := range s {
		cnt[byte(c)]++
	}
	for _, c := range s {
		// 每遍历过一个字符，都将对应的计数减一
		cnt[byte(c)]--
		if val, ok := inStack[byte(c)]; val && ok {
			continue
		}
		// 插入之前，和之前的元素比较一下大小
		// 如果字典序比前面的小，pop 前面的元素
		for len(stack) > 0 && stack[len(stack)-1] > byte(c) {
			if cnt[stack[len(stack)-1]] == 0 {
				break
			}
			inStack[stack[len(stack)-1]] = false
			// pop
			stack = stack[:len(stack)-1]
		}
		// push c
		stack = append(stack, byte(c))
		inStack[byte(c)] = true
	}

	return string(stack)
}
func Test_remove_duplicate_letters(t *testing.T) {
	s := "bcabc"
	t.Log(removeDuplicateLetters(s))
	s = "cbacdcbc"
	t.Log(removeDuplicateLetters(s))
}
