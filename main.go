package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	// Check if at least 4 arguments are provided (os.Args includes the program name as the first argument)
	if len(args) < 4 {
		fmt.Println("Usage: <program> <website> <max_concurrency> <max_pages>")
		os.Exit(1)
	}

	rawBaseURL := args[1]
	maxConcurrency := args[2]
	maxPages := args[3]

	// Assuming getConfig is defined elsewhere and handles your settings
	cfg, err := getConfig(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Starting crawl of: %s...\n", rawBaseURL)

	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("Crawl complete")

	cfg.printReport()
}
