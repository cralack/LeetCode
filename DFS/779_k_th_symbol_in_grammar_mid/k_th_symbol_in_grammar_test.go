package kthsymbolingrammar

import (
	"testing"
)

func kthGrammar(n int, k int) int {
	if k == 1 {
		return 0
	}
	if k > (1 << (n - 2)) {
		return 1 ^ kthGrammar(n-1, k-(1<<(n-2)))
	} else {
		return kthGrammar(n-1, k)
	}
}
func Test_k_th_symbol_in_grammar(t *testing.T) {
	n, k := 1, 1
	t.Log(kthGrammar(n, k))
	n, k = 2, 1
	t.Log(kthGrammar(n, k))
	n, k = 2, 2
	t.Log(kthGrammar(n, k))
}
