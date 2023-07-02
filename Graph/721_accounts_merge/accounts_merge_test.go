package accountsmerge

import (
	"sort"
	"testing"
)

func accountsMerge(accounts [][]string) (ans [][]string) {
	mail2index := make(map[string]int, 0)
	mail2name := make(map[string]string, 0)
	for _, acaccount := range accounts {
		name := acaccount[0]
		for _, mail := range acaccount[1:] {
			if _, has := mail2index[mail]; !has { // 避免mail重复登记
				mail2index[mail] = len(mail2index)
				mail2name[mail] = name
			}
		}
	}
	parent := make([]int, len(mail2index))
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}
	for _, acaccount := range accounts {
		firstIndex := mail2index[acaccount[1]]
		for _, mail := range acaccount[2:] {
			union(mail2index[mail], firstIndex)
		}
	}

	index2mail := map[int][]string{}
	for mail, index := range mail2index {
		index = find(index)
		index2mail[index] = append(index2mail[index], mail)
	}

	for _, mails := range index2mail {
		sort.Strings(mails)
		account := append([]string{mail2name[mails[0]]}, mails...)
		ans = append(ans, account)
	}

	return
}
func Test_accounts_merge(t *testing.T) {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"}}
	t.Log(accountsMerge(accounts))
	accounts = [][]string{
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
		{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
		{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
		{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
		{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"}}
	t.Log(accountsMerge(accounts))
}
