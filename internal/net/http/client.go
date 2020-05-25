package http

import (
	"fmt"
	"net/http"
)

// FetchURL function
func FetchURL(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()

		return nil, fmt.Errorf("%v", err)
	}

	return resp, nil
}
