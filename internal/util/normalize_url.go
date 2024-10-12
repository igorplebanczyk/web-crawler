package util

import "net/url"

func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	normalizedURL := parsedURL.Host + parsedURL.Path

	if parsedURL.RawQuery != "" {
		normalizedURL += "?" + parsedURL.RawQuery
	}

	if parsedURL.Fragment != "" {
		normalizedURL += "#" + parsedURL.Fragment
	}

	return normalizedURL, nil
}
