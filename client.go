package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// NewGHClient creates a new Github client.
func NewGHClient(ctx context.Context) *github.Client {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		log.Fatal("GITHUB_TOKEN environment variable is not set")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client
}
