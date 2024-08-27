package main

import (
	"fmt"
	"net/url"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func getConfig(rawBaseURL, maxConcurrency, maxPages string) (*config, error) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %s", err)
	}

	maxConcurrencyInt, err := strconv.Atoi(maxConcurrency)
	if err != nil {
		return nil, fmt.Errorf("failed to parse max concurrency: %s", err)
	}

	maxPagesInt, err := strconv.Atoi(maxPages)
	if err != nil {
		return nil, fmt.Errorf("failed to parse max pages: %s", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrencyInt),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPagesInt,
	}, nil
}
