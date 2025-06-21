package minimum_deletions_to_make_string_k_special_mid

import (
	"slices"
	"testing"
)

func minimumDeletions(word string, k int) int {
	cnt := make([]int, 26)
	for _, b := range word {
		cnt[b-'a']++
	}

	slices.Sort(cnt)
	for i, c := range cnt {
		if c != 0 {
			cnt = cnt[i:]
			break
		}
	}

	maxSave := 0
	for i, base := range cnt {
		sum := 0
		for _, c := range cnt[i:] {
			sum += min(c, base+k) // 至多保留 base+k 个
		}
		maxSave = max(maxSave, sum)
	}

	return len(word) - maxSave
}

func minimumDeletionsSliceWindow(word string, k int) int {
	const sigma = 26
	cnt := make([]int, sigma)
	for _, b := range word {
		cnt[b-'a']++
	}
	slices.Sort(cnt)
	for i, c := range cnt {
		if c != 0 {
			cnt = cnt[i:]
			break
		}
	}

	var maxSave, sum, right int
	for _, base := range cnt {
		for right < len(cnt) && cnt[right] <= base+k {
			sum += cnt[right]
			right++
		}
		maxSave = max(maxSave, sum+(base+k)*(len(cnt)-right))
		sum -= base
	}
	return len(word) - maxSave
}
func TestMinDeletionKSpecial(t *testing.T) {
	tests := []struct {
		word string
		k    int
	}{
		{word: "aabcaba", k: 0},
		{word: "dabdcbdcdcd", k: 2},
		{word: "aaabaaa", k: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Log(minimumDeletions(tt.word, tt.k))
			t.Log(minimumDeletionsSliceWindow(tt.word, tt.k))
		})
	}
}
