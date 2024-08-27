package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	maxConcurrency := os.Args[2]
	maxPages := os.Args[3]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg, err := getConfig(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println(err)
	}

	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("crawl complete")
}
