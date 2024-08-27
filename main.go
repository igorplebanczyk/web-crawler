package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	pages := make(map[string]int)
	crawlPage(rawBaseURL, rawBaseURL, pages)
	for k, v := range pages {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Println("crawl complete")
}
