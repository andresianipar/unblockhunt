package html

import (
	"fmt"
	"net/http"
	netURL "net/url"
	"strings"

	netHTTP "github.com/andresianipar/unblockhunt/internal/net/http"
	"golang.org/x/net/html"
)

// Hosts type
// TODO: should have a method to modify itself, i.e., avoid direct access
var Hosts map[string]bool

// FindLinks function
func FindLinks(c func(item string) ([]string, error), urls []string) error {
	seen := make(map[string]bool)

	// TODO: find other ways to set the maximum number of urls
	for len(urls) <= 50 {
		items := urls

		urls = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				links, err := c(item)

				if err != nil {
					return fmt.Errorf("%v", err)
				}
				urls = append(urls, links...)
			}
		}
	}

	return nil
}

// TODO: get a better approach
func pre(url string, links []string, n *html.Node, resp *http.Response) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link, err := resp.Request.URL.Parse(a.Val)

				if err != nil {
					continue
				}

				parsedURL, err := netURL.Parse(url)

				if err != nil {
					continue
				}

				var temp string
				ss := strings.Split(link.Host, ".")

				if len(ss) > 2 {
					temp = strings.Join(ss[1:3], ".")
					if temp == parsedURL.Host {
						if !Hosts[link.Host] {
							Hosts[link.Host] = true
						}
						links = append(links, link.String())
					}
				} else {
					temp = strings.Join(ss[:], ".")
					if temp == parsedURL.Host {
						if !Hosts[link.Host] {
							Hosts[link.Host] = true
						}
						links = append(links, link.String())
					}
				}
			}
		}
	}

	return links
}

func visit(url string, links []string, n *html.Node, resp *http.Response) []string {
	links = pre(url, links, n, resp)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(url, links, c, resp)
	}

	return links
}

// Crawl function
// Example input:  https://www.google.com
// Example output: [https://www.google.com/ https://a.google.com/ ...]
func Crawl(url string) ([]string, error) {
	resp, err := netHTTP.FetchURL(url)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()
	if ct := resp.Header.Get("Content-Type"); ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return nil, fmt.Errorf("title: %s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	links := visit(url, nil, doc, resp)

	return links, nil
}
