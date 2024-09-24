package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	owner := flag.String("owner", "", "GitHub repository owner")
	repo := flag.String("repo", "", "GitHub repository name")
	count := flag.Int("count", 10, "Number of pull requests to fetch")

	flag.Parse()

	if *owner == "" || *repo == "" {
		log.Fatal("owner and repo flags are required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	client := NewGHClient(ctx)

	requester := NewPRRequester(client, *owner, *repo)

	prStats := requester.GetPRStats(ctx, *count)
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fileNamePrefix := fmt.Sprintf("pr_%s_%s", *repo, timestamp)
	writeJsonFile(prStats, fileNamePrefix)
	writeMDFile(prStats, fileNamePrefix)
}

func writeJsonFile(prs []PR, fileName string) {
	// Convert Go struct to JSON with indentation for clean formatting
	fileData, err := json.MarshalIndent(prs, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Create or open the file for writing
	file, err := os.Create(fileName + ".json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(fileData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
