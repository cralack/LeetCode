package aliendictionary

import "testing"

func alienOrder(words []string) string {
	// 只有一个单词，任意排序
	if len(words) == 1 {
		return words[0]
	}
	// 某字母对应的入度（没有前置条件）
	inCount := make(map[string]int, 0)
	// 有向图
	graph := make(map[string][]string, 0)
	// 记录出现的字母
	canTake := make(map[string]bool, 0)
	index := 0
	for index < len(words)-1 {
		next := index + 1
		// 从左往右比较index跟next的首个不同字母，可以确定这两个字母的顺序的大小
		found := false
		charIndex := 0
		for charIndex < len(words[index]) || charIndex < len(words[next]) {
			if charIndex < len(words[index]) && charIndex < len(words[next]) {
				from := string(words[index][charIndex])
				to := string(words[next][charIndex])
				// 记录出现的字母
				canTake[from] = true
				canTake[to] = true
				if words[index][charIndex] != words[next][charIndex] && !found {
					// 发现一个能确定大小顺序的字符对：from->to
					graph[from] = append(graph[from], to)
					// 记录入度
					inCount[to]++
					// 记录确认过（两个单词只能记录一对字母的大小-首次不同的那一对）
					found = true
				}
			} else if charIndex < len(words[index]) {
				// 记录出现的字母
				canTake[string(words[index][charIndex])] = true
			} else {
				// 记录出现的字母
				canTake[string(words[next][charIndex])] = true
			}
			charIndex++
		}
		// 如果都一样，但是前面的字母比后面的字母要长，不合理，无法确定顺序
		if !found && len(words[index]) > len(words[next]) {
			return ""
		}
		index = next
	}

	// 记录单词数量：用于记录是否全部字母排序完成
	charCount := len(canTake)
	// fmt.Println(inCount)
	// fmt.Println(canTake)
	// fmt.Println(charCount)
	// fmt.Println(graph)
	result := []byte{}
	stop := false
	for {
		// 没有发现入度为0的字符，无法继续去边
		if stop {
			break
		}
		stop = true
		for char := range canTake {
			// 首次发现入度为0的字符，去边，并减少后面排序的字母入度，方便下一次循环
			if inCount[char] == 0 && canTake[char] {
				result = append(result, byte(char[0]))
				canTake[char] = false
				for _, to := range graph[char] {
					inCount[to]--
				}
				stop = false
			}
			// fmt.Println(string(result))
		}
	}

	// 是否全部取完
	if len(result) == charCount {
		return string(result)
	}
	return ""
}

func Test_alien_dictionary(t *testing.T) {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	t.Log(alienOrder(words))
	words = []string{"z", "x"}
	t.Log(alienOrder(words))
	words = []string{"z", "x", "z"}
	t.Log(alienOrder(words))
}
