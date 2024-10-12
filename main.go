package main

import (
	"fmt"
	"web-crawler/cli"
)

func main() {
	err := cli.Execute()
	if err != nil {
		fmt.Printf("failed to execute CLI: %s", err)
	}
}
