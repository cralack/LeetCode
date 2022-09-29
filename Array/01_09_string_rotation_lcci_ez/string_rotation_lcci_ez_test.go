package stringrotationlcci

import (
	"strings"
	"testing"
)

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
func Test_string_rotation_lcci_ez(t *testing.T) {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	t.Log(isFlipedString(s1, s2))
}
