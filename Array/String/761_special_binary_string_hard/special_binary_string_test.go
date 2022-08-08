package specialbinarystringhard

import (
	"sort"
	"strings"
	"testing"
)

func makeLargestSpecial(s string) string {
	//递归到的字符串长度小于等于 22 时，说明字符串要么为空，
	//要么为 10，此时字符串就是字典序最大的结果，可以直接返回
	if len(s) <= 2 {
		return s
	}
	subs := sort.StringSlice{}
	cnt, left := 0, 0
	for i, ch := range s {
		if ch == '1' { //首位 1 和末位 0 直接移除
			cnt++
		} else if cnt--; cnt == 0 { //拆分成多个特殊序列
			subs = append(subs, "1"+makeLargestSpecial(s[left+1:i])+"0")
			left = i + 1
		}
	} //降序排序后拼接,得到字典序最大的字符串
	sort.Sort(sort.Reverse(subs))
	return strings.Join(subs, "")
}

func Test_special_binary_string(t *testing.T) {
	str := "11011000"
	t.Log(makeLargestSpecial(str))
}
