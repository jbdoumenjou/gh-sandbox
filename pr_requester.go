package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
)

// PRRequester allows to extract stats from a repo.
type PRRequester struct {
	client *github.Client
	owner  string
	repo   string
}

// NewPRRequester creates a new PR Requester.
func NewPRRequester(client *github.Client, owner, repo string) *PRRequester {
	return &PRRequester{
		client: client,
		owner:  owner,
		repo:   repo,
	}
}

// GetFirstReviewTime return the time between the pr creation and the first review.
func (p *PRRequester) GetFirstReviewTime(ctx context.Context, prNumber int) *time.Time {
	reviews, _, err := p.client.PullRequests.ListReviews(ctx, p.owner, p.repo, prNumber, &github.ListOptions{})
	if err != nil {
		log.Printf("Error fetching reviews for PR #%d: %v", prNumber, err)

		return nil
	}

	if len(reviews) > 0 {
		firstReview := reviews[0]

		return firstReview.SubmittedAt
	}

	return nil
}

// GetPRStats gets a list of PR containing stats.
func (p *PRRequester) GetPRStats(ctx context.Context, count int) []PR {
	// Fetch the last N pull requests
	prs, _, err := p.client.PullRequests.List(ctx, p.owner, p.repo, &github.PullRequestListOptions{
		State:       "all",
		ListOptions: github.ListOptions{PerPage: count},
	})
	if err != nil {
		log.Fatalf("Error fetching pull requests: %v", err)
	}

	prDataSet := []PR{}

	for _, pr := range prs {
		pData := PR{
			Title:     pr.GetTitle(),
			Number:    pr.GetNumber(),
			CreatedAt: pr.GetCreatedAt(),
		}

		createdAt := pr.GetCreatedAt()
		firstReviewTime := p.GetFirstReviewTime(ctx, pr.GetNumber())

		if firstReviewTime != nil {
			pData.FirstReviewAt = *firstReviewTime
			timeDifference := firstReviewTime.Sub(createdAt)
			pData.TimeDifference = timeDifference
		}

		if pr.MergedAt != nil {
			mergedAt := pr.GetMergedAt()
			mergeTimeDifference := mergedAt.Sub(createdAt)
			pData.MergedAt = mergedAt
			pData.MergeTimeDifference = mergeTimeDifference
		}

		// Fetch comments for the pull request
		comments, _, err := p.client.PullRequests.ListComments(
			ctx, p.owner, p.repo, pr.GetNumber(), &github.PullRequestListCommentsOptions{},
		)
		if err != nil {
			log.Fatalf("Error fetching pull request comments: %v", err)
		}

		pData.CommentCount = len(comments)

		prDataSet = append(prDataSet, pData)
	}

	return prDataSet
}

// LogMD displays a mardown table of the prs stats list on the standard output.
func LogMD(prs []PR) {
	// Create a logger that writes to the file
	logger := log.New(os.Stdout, "", 0)

	logger.Println(
		"| PR Number | Title | Created At | First Review At | Time Waiting First Review |" +
			" Merged At | Time Between Creation and Merge | Comment count |",
	)
	logger.Println("| --- | --- | --- | --- | --- | --- | --- | --- |")

	for _, pr := range prs {
		logger.Printf("%s\n", pr)
	}
}
