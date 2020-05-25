package main

import (
	"fmt"
	"os"

	"github.com/andresianipar/unblockhunt/internal/net"
	"github.com/andresianipar/unblockhunt/pkg/html"
	"github.com/andresianipar/unblockhunt/pkg/io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	// TODO: 1. sanitize input
	// 		 2. process all hosts concurrently
	for _, host := range os.Args[1:] {
		hosts, err := html.GetHosts(host)

		checkErr(err)

		hostsaips, err := net.LookupIPs(hosts)

		checkErr(err)
		hostsaips, err = hostsaips.Filter(func(ip string) (bool, error) {
			iPVersion, err := net.CheckIPVersion(ip)

			if err != nil {
				return false, fmt.Errorf("%v", err)
			}
			if iPVersion != net.IPv4 {
				return false, nil
			}

			return true, nil
		})
		checkErr(err)
		err = ioutil.AddHosts(hostsaips)
		checkErr(err)
	}
}
