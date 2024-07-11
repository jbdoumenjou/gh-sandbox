package main

import (
	"context"
	"flag"
	"log"
)

func main() {
	owner := flag.String("owner", "", "GitHub repository owner")
	repo := flag.String("repo", "", "GitHub repository name")
	count := flag.Int("count", 10, "Number of pull requests to fetch")

	flag.Parse()

	if *owner == "" || *repo == "" {
		log.Fatal("owner and repo flags are required")
	}

	ctx := context.Background()
	client := NewGHClient(ctx)

	requester := NewPRRequester(client, *owner, *repo)

	prStats := requester.GetPRStats(ctx, *count)

	LogMD(prStats)
}
