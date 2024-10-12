package crawl

import (
	"fmt"
	"net/url"
	"web-crawler/internal/util"
)

func (cfg *Config) CrawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.Wg.Done()
	}()

	if cfg.getPagesLength() >= cfg.maxPages {
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse current URL: %s\n", err)
		return
	}

	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := util.NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize current URL: %s\n", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to get HTML: %s\n", err)
		return
	}

	urls, err := GetURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("failed to get URLs from HTML: %s\n", err)
		return
	}

	for _, u := range urls {
		cfg.Wg.Add(1)
		go cfg.CrawlPage(u)
	}
}

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.pages[normalizedURL]
	if ok {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *Config) getPagesLength() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages)
}
