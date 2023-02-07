package alertusingsamekeycard

import (
	"sort"
	"testing"
)

func alertNames(keyName []string, keyTime []string) (ans []string) {
	dic := make(map[string][]int)
	for i, name := range keyName {
		t := keyTime[i]
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		dic[name] = append(dic[name], hour*60+minute)
	}
	for name, times := range dic {
		if len(times) > 2 {
			sort.Ints(times)
			for i := 0; i < len(times)-2; i++ {
				if times[i+2]-times[i] <= 60 {
					ans = append(ans, name)
					break
				}
			}
		}
	}
	sort.Strings(ans)
	return
}
func Test_alert_using_same_key_card(t *testing.T) {
	keyName := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	keyTime := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
	t.Log(alertNames(keyName, keyTime))
	keyName = []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"}
	keyTime = []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}
	t.Log(alertNames(keyName, keyTime))
	keyName = []string{"john", "john", "john"}
	keyTime = []string{"23:58", "23:59", "00:01"}
	t.Log(alertNames(keyName, keyTime))
	keyName = []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"}
	keyTime = []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"}
	t.Log(alertNames(keyName, keyTime))
}
