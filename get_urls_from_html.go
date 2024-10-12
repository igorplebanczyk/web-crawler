package main

import (
	"golang.org/x/net/html"
	"strings"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	node, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	return extractURLs(node, rawBaseURL), nil
}

func extractURLs(node *html.Node, rawBaseURL string) []string {
	extractedURLs := []string{}

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				url := attr.Val
				// Handle relative URLs by prepending rawBaseURL
				if !strings.HasPrefix(url, "http") {
					url = rawBaseURL + url
				}
				// Append only the absolute URL
				extractedURLs = append(extractedURLs, url)
			}
		}
	}

	// Recursively process the first child and then the next sibling
	if node.FirstChild != nil {
		extractedURLs = append(extractedURLs, extractURLs(node.FirstChild, rawBaseURL)...)
	}
	if node.NextSibling != nil {
		extractedURLs = append(extractedURLs, extractURLs(node.NextSibling, rawBaseURL)...)
	}

	return extractedURLs
}
