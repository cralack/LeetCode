package designauthenticationmanager

import "testing"

type AuthenticationManager struct {
	timeLimit int
	cache     map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeLimit: timeToLive,
		cache:     make(map[string]int),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.cache[tokenId] = currentTime + this.timeLimit
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, has := this.cache[tokenId]; !has || v <= currentTime {
		return
	}
	this.Generate(tokenId, currentTime)
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	ans := 0
	for _, exp := range this.cache {
		if exp > currentTime {
			ans++
		}
	}

	return ans
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */

func Test_design_authentication_manager(t *testing.T) {
	cm1 := []string{"AuthenticationManager", "renew", "generate", "countUnexpiredTokens",
		"generate", "renew", "renew", "countUnexpiredTokens"}
	cm2 := []string{"", "aaa", "aaa", "", "bbb", "aaa", "bbb", ""}
	cm3 := []int{5, 1, 2, 6, 7, 8, 10, 15}

	var sol AuthenticationManager
	for i, c := range cm1 {
		switch c {
		case "AuthenticationManager":
			sol = Constructor(cm3[i])
		case "renew":
			sol.Renew(cm2[i], cm3[i])
		case "generate":
			sol.Generate(cm2[i], cm3[i])
		case "countUnexpiredTokens":
			t.Log(sol.CountUnexpiredTokens(cm3[i]))
		}
	}

}
