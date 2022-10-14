package distinctsubsequencesii

import "testing"

func distinctSubseqII(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	preCnt := make([]int, 26)
	curAns := 1
	str := []byte(s)
	for i := 0; i < n; i++ {
		newCnt := curAns
		curAns = ((curAns+newCnt)%mod - preCnt[str[i]-'a']%mod + mod) % mod
		preCnt[str[i]-'a'] = newCnt
	}
	return curAns - 1
}

func Test_distinct_subsequences_ii(t *testing.T) {
	s := "abc"
	t.Log(distinctSubseqII(s))
	s = "aba"
	t.Log(distinctSubseqII(s))
	s = "aaa"
	t.Log(distinctSubseqII(s))
}
