package main

import (
	"fmt"
	"sort"
)

func (cfg *config) printReport() {
	fmt.Println("===============================")
	fmt.Printf("REPORT for %s\n", cfg.baseURL)
	fmt.Println("===============================")

	sorted := sortMapByValue(cfg.pages)
	for _, v := range sorted {
		fmt.Printf("Found %d internal links to %s\n", v.count, v.url)
	}

	fmt.Println("===============================")
	fmt.Println("REPORT COMPLETE")
	fmt.Println("===============================")
}

func sortMapByValue(m map[string]int) []struct {
	url   string
	count int
} {
	var sorted []struct {
		url   string
		count int
	}
	for k, v := range m {
		sorted = append(sorted, struct {
			url   string
			count int
		}{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})

	return sorted
}
