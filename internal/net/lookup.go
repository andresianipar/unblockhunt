package net

import (
	"fmt"
	"net"

	"github.com/andresianipar/unblockhunt/internal/builtin"
)

// HostsAIPs type
type HostsAIPs map[string][]string

// Filter function
func (hostsaips HostsAIPs) Filter(c func(ip string) (bool, error)) (HostsAIPs, error) {
	for host, ips := range hostsaips {
		for i, ip := range ips {
			ok, err := c(ip)

			if err != nil {
				return nil, fmt.Errorf("%v", err)
			}
			if !ok {
				ips = builtin.Remove(ips, i)
			}
		}
		hostsaips[host] = ips
	}

	return hostsaips, nil
}

// LookupIPs function
// Example input:  [google.com a.google.com ...]
// Example output: map[google.com:[172.16.254.1 2001:db8:0:1234:0:567:8:1 ...]
// 				   a.google.com:[172.16.254.2 2001:db8:0:1234:0:567:8:2 ...]
// 				   ...]
func LookupIPs(hosts []string) (HostsAIPs, error) {
	var hostsaips HostsAIPs = make(map[string][]string)

	for _, host := range hosts {
		ips, err := net.LookupIP(host)

		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}

		var temp []string

		for _, ip := range ips {
			temp = append(temp, fmt.Sprintf("%s", ip))
		}
		hostsaips[host] = temp
	}

	return hostsaips, nil
}
