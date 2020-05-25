package ioutil

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/andresianipar/unblockhunt/internal/net"
)

// TODO: should be able to choose the right location of the hosts file
// based on the current OS
const (
	unixHostsPath    = "/etc/hosts"
	windowsHostsPath = "C:\\Windows\\System32\\drivers\\etc\\hosts"
)

// Example input: [172.16.254.1 google.com
// 				  172.16.254.2 a.google.com
// 				  ...]
func addHosts(hosts []string) error {
	temp, err := ioutil.ReadFile(unixHostsPath)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	lines := strings.Split(string(temp), "\n")

	for i, line := range lines {
		if line == "" {
			copy(lines[i+1:], lines[i:])
			lines[i] = fmt.Sprintf("%s", strings.Join(hosts, "\n"))
			break
		}
	}
	lines = append(lines, "")
	err = ioutil.WriteFile(unixHostsPath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

// AddHosts function
// Example input: map[google.com:[172.16.254.1 2001:db8:0:1234:0:567:8:1 ...]
// 				  a.google.com:[172.16.254.2 2001:db8:0:1234:0:567:8:2 ...]
// 				  ...]
func AddHosts(hostsaips net.HostsAIPs) error {
	var hosts []string

	for host, ips := range hostsaips {
		hosts = append(hosts, fmt.Sprintf("%s %s", ips[0], host))
	}

	err := addHosts(hosts)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
