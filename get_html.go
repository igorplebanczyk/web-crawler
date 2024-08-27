package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return "", fmt.Errorf("non-200 status code: %d", res.StatusCode)
	}

	if res.Header["Content-Type"][0] != "text/html" {
		return "", fmt.Errorf("non-html content type")
	}

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(html), nil
}
