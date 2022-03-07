package satisfiabilityofequalityequations

import "testing"

func equationsPossible(equations []string) bool {
	n := len(equations)
	if n == 0 {
		return true
	}
	parent, size := make([]int, 26), make([]int, 26)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	trans := func(p, q byte) (int, int) {
		return int(p - 'a'), int(q - 'a')
	}
	find := func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}
	connected := func(p, q int) bool {
		rootP := find(p)
		rootQ := find(q)
		return rootP == rootQ
	}
	union := func(p, q int) {
		rootP := find(p)
		rootQ := find(q)
		if rootP == rootQ {
			return
		}
		if size[rootP] > size[rootQ] {
			parent[rootQ] = rootP
			size[rootP] += size[rootQ]
		} else {
			parent[rootP] = rootQ
			size[rootQ] += rootP
		}
	}
	for _, eq := range equations {
		if eq[1] == '=' {
			x, y := eq[0], eq[3]
			union(trans(x, y))
		}
	}
	for _, eq := range equations {
		if eq[1] == '!' {
			x, y := eq[0], eq[3]
			if connected(trans(x, y)) {
				return false
			}
		}
	}
	return true
}
func Test_satisfiability_of_equality_equations(t *testing.T) {
	strings := [][]string{
		{"a==b", "b!=c", "c==a"},
		{"c==c", "b==d", "x!=z"}}
	t.Log(equationsPossible(strings[0]))
	t.Log(equationsPossible(strings[1]))
}
