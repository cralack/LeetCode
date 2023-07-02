package validateipaddress

import (
	"strconv"
	"strings"
	"testing"
)

func validIPAddress(queryIP string) string {
	// ipv4 judge
	if query := strings.Split(queryIP, "."); len(query) == 4 {
		for _, v := range query {
			if len(v) > 1 && v[0] == '0' {
				return "Neither"
			}
			if v, err := strconv.Atoi(v); err != nil || v > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	// ipv6 judge
	if query := strings.Split(queryIP, ":"); len(query) == 8 {
		for _, v := range query {
			if len(v) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(v, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
func Test_validate_ip_address(t *testing.T) {
	queryIP := "172.16.254.1"
	t.Log(validIPAddress(queryIP))
	queryIP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
	t.Log(validIPAddress(queryIP))
	queryIP = "2001:0db8:85a3::8A2E:037j:7334"
	t.Log(validIPAddress(queryIP))
	queryIP = "02001:0db8:85a3:0000:0000:8a2e:0370:7334"
	t.Log(validIPAddress(queryIP))
	queryIP = "256.256.256.256"
	t.Log(validIPAddress(queryIP))
}
