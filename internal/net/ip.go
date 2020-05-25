package net

import (
	"fmt"
	"net"
)

const (
	// IPv4 constant
	IPv4 = 4
	// IPv6 constant
	IPv6 = 6
)

// CheckIPVersion function
// Example input:  172.16.254.1 in IPv4, or 2001:db8:0:1234:0:567:8:1 in IPv6
// Example output: one of the constants IPv4 or IPv6
func CheckIPVersion(ip string) (int, error) {
	if net.ParseIP(string(ip)) == nil {
		return 0, fmt.Errorf("invalid IP address: %s", ip)
	}
	for i := range ip {
		switch ip[i] {
		case '.':
			return IPv4, nil
		case ':':
			return IPv6, nil
		}
	}

	return 0, fmt.Errorf("invalid IP address: %s", ip)
}
