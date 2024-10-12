package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"web-crawler/internal/crawl"
)

var rootCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawler is a CLI tool to crawl a website and get internal links",
	Long:  `Crawler is a CLI tool to crawl a website and get internal links. It takes a website URL, a maximum concurrency level, and a maximum number of pages to crawl.`,
	RunE:  rootCmdRun,
}

func init() {
	rootCmd.Flags().StringP("website", "w", "", "Website URL to crawl")
	rootCmd.Flags().IntP("max-concurrency", "c", 10, "Maximum number of goroutines")
	rootCmd.Flags().IntP("max-pages", "p", 100, "Maximum number of pages to crawl")
}

func rootCmdRun(cmd *cobra.Command, _ []string) error {
	website, err := cmd.Flags().GetString("website")
	maxConcurrency, _ := cmd.Flags().GetInt("max-concurrency")
	maxPages, _ := cmd.Flags().GetInt("max-pages")

	if err != nil {
		return fmt.Errorf("failed to get website: %v", err)
	}

	cfg, err := crawl.GetConfig(website, maxConcurrency, maxPages)
	if err != nil {
		return fmt.Errorf("failed to get config: %v", err)
	}

	fmt.Printf("Starting crawl of: %s...\n", website)

	cfg.CrawlPage(website)
	cfg.Wg.Wait()

	fmt.Println("Crawl complete")

	cfg.PrintReport()

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}
