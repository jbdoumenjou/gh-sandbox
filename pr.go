package main

import (
	"fmt"
	"time"
)

// PR represent PR data.
type PR struct {
	Number              int
	Title               string
	CreatedAt           time.Time
	FirstReviewAt       time.Time
	TimeDifference      time.Duration
	MergedAt            time.Time
	MergeTimeDifference time.Duration
	CommentCount        int
}

func (pr PR) String() string {
	firstReviewAt := "NA"
	if !pr.FirstReviewAt.IsZero() {
		firstReviewAt = pr.FirstReviewAt.Format(time.RFC3339)
	}

	mergedAt := "NA"
	if !pr.MergedAt.IsZero() {
		mergedAt = pr.MergedAt.Format(time.RFC3339)
	}

	timeDifference := "NA"
	if pr.TimeDifference != 0 {
		timeDifference = pr.TimeDifference.String()
	}

	prMergeDiff := "NA"
	if pr.MergeTimeDifference != 0 {
		prMergeDiff = pr.MergeTimeDifference.String()
	}

	return fmt.Sprintf(
		"| #%d | %s | %v | %v | %v | %v | %v | %d |",
		pr.Number, pr.Title, pr.CreatedAt.Format(time.RFC3339),
		firstReviewAt, timeDifference, mergedAt, prMergeDiff, pr.CommentCount,
	)
}
