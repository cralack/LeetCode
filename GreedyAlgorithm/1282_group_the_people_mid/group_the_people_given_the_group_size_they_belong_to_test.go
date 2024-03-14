package groupthepeoplegiventhegroupsizetheybelongtomid

import "testing"

func groupThePeople(groupSizes []int) (ans [][]int) {
	mp := make(map[int][]int, 0)
	for idx, cnt := range groupSizes {
		mp[cnt] = append(mp[cnt], idx)
	}
	for cnt, idxs := range mp {
		for j := 0; j < len(idxs); j += cnt {
			ans = append(ans, idxs[j:j+cnt])
		}
	}
	return
}
func Test_group_the_people_given_the_group_size_they_belong_to(t *testing.T) {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	t.Log(groupThePeople(groupSizes))
	groupSizes = []int{2, 1, 3, 3, 3, 2}
	t.Log(groupThePeople(groupSizes))
}
