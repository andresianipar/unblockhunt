package html

import (
	"fmt"

	"github.com/andresianipar/unblockhunt/internal/html"
)

// GetHosts function
// Example input:  google.com
// Example output: [google.com a.google.com ...]
func GetHosts(host string) ([]string, error) {
	html.Hosts = make(map[string]bool)

	// FIXME: the url protocol should not be hardcoded
	if err := html.FindLinks(html.Crawl, []string{"https://" + host}); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var hosts []string

	for host := range html.Hosts {
		hosts = append(hosts, host)
	}

	return hosts, nil
}
