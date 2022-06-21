package defanginganipaddresseasy

import (
	"strings"
	"testing"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
func Test_defanging_an_ip_address(t *testing.T) {
	address := "1.1.1.1"
	t.Log(defangIPaddr(address))
	address = "255.100.50.0"
	t.Log(defangIPaddr(address))
}
