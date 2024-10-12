package crawl

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
	maxPages           int
}

func GetConfig(rawBaseURL string, maxConcurrency, maxPages int) (*Config, error) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %s", err)
	}

	return &Config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}
