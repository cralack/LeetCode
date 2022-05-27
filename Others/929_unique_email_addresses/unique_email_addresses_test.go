package uniqueemailaddresses

import (
	"strings"
	"testing"
)

func numUniqueEmails(emails []string) int {
	addrCnt := make(map[string]int, 0)
	for _, email := range emails {
		i := strings.IndexByte(email, '@')
		local := strings.SplitN(email[:i], "+", 2)[0] // 去掉本地名第一个加号之后的部分
		local = strings.ReplaceAll(local, ".", "")    // 去掉本地名中所有的句点
		addrCnt[local+email[i:]]++
	}
	return len(addrCnt)
}
func Test_unique_email_addresses(t *testing.T) {
	emails := []string{"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com"}
	t.Log(numUniqueEmails(emails))
	emails = []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	t.Log(numUniqueEmails(emails))
}
