package throne_inheritance_mid

import (
	"testing"
)

type ThroneInheritance struct {
	king     string
	children map[string][]string
	dead     map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		king:     kingName,
		children: make(map[string][]string),
		dead:     make(map[string]bool),
	}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.children[parentName] = append(t.children[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() (ans []string) {
	var dfs func(string)
	dfs = func(cur string) {
		if !t.dead[cur] {
			ans = append(ans, cur)
		}
		for _, child := range t.children[cur] {
			dfs(child)
		}
	}
	dfs(t.king)
	return ans
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
func Test_throne_inheritance(t *testing.T) {
	tests := []struct {
		command []string
		flag    [][]string
	}{
		{
			[]string{"ThroneInheritance", "birth", "birth", "birth", "birth", "birth", "birth", "getInheritanceOrder", "death", "getInheritanceOrder"},
			[][]string{{"king"}, {"king", "andy"}, {"king", "bob"}, {"king", "catherine"}, {"andy", "matthew"}, {"bob", "alex"}, {"bob", "asha"}, {"null"}, {"bob"}, {"null"}},
		},
	}
	for _, tt := range tests {
		var sol ThroneInheritance
		for i, cmd := range tt.command {
			switch cmd {
			case "ThroneInheritance":
				sol = Constructor(tt.flag[i][0])
			case "birth":
				sol.Birth(tt.flag[i][0], tt.flag[i][1])
			case "death":
				sol.Death(tt.flag[i][0])
			case "getInheritanceOrder":
				t.Log(sol.GetInheritanceOrder())
			}
		}
	}
}
