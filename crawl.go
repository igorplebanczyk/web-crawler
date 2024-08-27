package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("failed to parse base URL: %s\n", err)
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse current URL: %s\n", err)
		return
	}

	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize current URL: %s\n", err)
		return
	}

	_, ok := pages[normalizedCurrentURL]
	if ok {
		pages[normalizedCurrentURL]++
		return
	}

	pages[normalizedCurrentURL] = 1

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to get HTML: %s\n", err)
		return
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("failed to get URLs from HTML: %s\n", err)
		return
	}

	for _, u := range urls {
		crawlPage(rawBaseURL, u, pages)
	}
}
